CREATE TABLE accounts (
    id BIGINT AUTO_INCREMENT ,
    owner varchar(255) NOT NULL,
    currency varchar(255) NOT NULL,
    balance bigint NOT NULL,
    created_at datetime ,
    PRIMARY KEY(id)
);

CREATE TABLE entries (
    id BIGINT AUTO_INCREMENT,
    account_id bigint NOT NULL,
    amount bigint NOT NULL,
    created_at datetime default CURRENT_TIMESTAMP,
    PRIMARY KEY(id),
    FOREIGN KEY (account_id) REFERENCES accounts(id)
);

CREATE TABLE transfers (
    id BIGINT AUTO_INCREMENT,
    from_accountid bigint NOT NULL,
    to_accountid bigint NOT NULL,
    amount bigint NOT NULL,
    created_at datetime default CURRENT_TIMESTAMP,
    PRIMARY KEY(id),
    FOREIGN KEY (from_accountid) REFERENCES accounts(id),
    FOREIGN KEY (to_accountid) REFERENCES accounts(id)
);

ALTER TABLE accounts AUTO_INCREMENT=220211995700;
CREATE INDEX account_index_0
ON accounts(owner);

CREATE INDEX account_index_1
ON entries(account_id);

CREATE INDEX account_index_2
ON transfers(from_accountid);

CREATE INDEX account_index_3
ON transfers(to_accountid);

CREATE INDEX account_index_4
ON transfers(from_accountid, to_accountid);