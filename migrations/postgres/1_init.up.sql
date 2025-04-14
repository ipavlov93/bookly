-- +goose Up
create table users
(
    id            serial primary key,
    firebase_uid  varchar(100) unique     not null,
    business_name varchar(100)            null,
    first_name     varchar(50)              null,
    last_name      varchar(50)              null,
    email_address varchar(100) unique     not null,
    organization  varchar(100)            not null,
    description   text                    null,
    created_at    timestamp default now() not null,
    updated_at    timestamp default now() not null,
    deleted_at    timestamp               null
);

-- +goose Up
create table user_accounts
(
    id            serial primary key,
    user_id       int          not null,
    issuer_code   int          not null,
    -- field is null for EmailPassword issuer
    -- todo: create FK to Issuers table
    subject_uid   varchar(100) null,
    email_address varchar(100) not null,
    constraint user_accounts_users_user_id
        foreign key (user_id)
            references users (id)
            on update cascade on delete cascade
    -- commented due to subject_uid is null
--     constraint unique_user_subject unique (user_id, subject_uid)
);
create INDEX IF NOT EXISTS user_accounts_user_id ON user_accounts (user_id);
