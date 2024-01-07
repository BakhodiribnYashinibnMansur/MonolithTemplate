-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE TABLE IF NOT EXISTS account(
    id uuid PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    full_name VARCHAR(64) NOT NULL,
    birthday_date TIMESTAMP NOT NULL,
    added_date TIMESTAMP NOT NULL,
    role VARCHAR(8) NOT NULL,
    phone_number VARCHAR(16) NOT NULL UNIQUE,
    password TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT (NOW()),
    updated_at TIMESTAMP NULL,
    deleted_at TIMESTAMP NULL
);
