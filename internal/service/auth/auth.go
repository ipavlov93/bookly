package auth

import (
	"context"
	"google.golang.org/api/option"
	"time"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
)

type AuthService struct {
	firebaseAuthService *auth.Client
}

func NewAuthService(firebaseCredentialFilePath string) AuthService {
	opt := option.WithCredentialsFile(firebaseCredentialFilePath)

	ctx := context.Background()
	firebaseApp, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		panic(err)
	}

	firebaseAuthService, err := firebaseApp.Auth(ctx)
	if err != nil {
		panic(err)
	}

	if firebaseAuthService == nil {
		panic("firebase auth service init failed")
	}

	return AuthService{
		firebaseAuthService: firebaseAuthService,
	}
}

// VerifyIDToken uses Firebase SDK (to implement OIDC flow) to verify ID token.
// Returns an err if token is invalid, expired, disabled or revoked.
// TODO: Add key to cache
func (s AuthService) VerifyIDToken(idToken string) (token *auth.Token, err error) {
	//if idToken == "" {
	//	return token, ErrInvalidToken
	//}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	token, err = s.firebaseAuthService.VerifyIDTokenAndCheckRevoked(ctx, idToken)
	if err != nil {
		return nil, err
	}
	//if err != nil {
	//	switch {
	//	case auth.IsIDTokenExpired(err):
	//		return token, ErrExpiredToken
	//	case auth.IsUserDisabled(err):
	//		return token, ErrDisabledToken
	//	case auth.IsIDTokenRevoked(err):
	//		return token, ErrRevokedToken
	//	}
	//
	//	return token, err
	//}

	// TODO: Add key to cache

	return token, nil
}

// TODO: remove token from cache
func (s AuthService) RevokeRefreshTokens(ctx context.Context, idToken string) error {
	token, err := s.firebaseAuthService.VerifyIDToken(ctx, idToken)
	if err != nil {
		return err
	}

	err = s.firebaseAuthService.RevokeRefreshTokens(context.Background(), token.UID)
	if err != nil {
		return err
	}

	// TODO: remove auth tokens from cache

	return nil
}
