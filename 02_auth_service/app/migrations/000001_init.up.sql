CREATE TABLE users (
    UUID uuid not null,
    email         varchar(255) not null unique,
    password_hash varchar(255) not null,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);