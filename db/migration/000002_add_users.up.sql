CREATE TABLE users (
    username varchar(255) NOT NULL,
    hashed_password varchar(255) NOT NULL,
    email varchar(255) NOT NULL,
    full_name varchar(255) NOT NULL,
    password_changed_at datetime DEFAULT '0001-01-01 00:00:00',
    created_at datetime default CURRENT_TIMESTAMP,
    UNIQUE(email),
    PRIMARY KEY(username)
);



ALTER TABLE accounts
ADD FOREIGN KEY(owner) REFERENCES users(username);

CREATE UNIQUE INDEX user_index
ON accounts (owner , currency);