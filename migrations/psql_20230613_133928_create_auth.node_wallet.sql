
-- +migrate Up
CREATE TABLE IF NOT EXISTS auth.node_wallet(
    id uuid NOT NULL PRIMARY KEY,
    wallet_id VARCHAR (100) NOT NULL,
    name VARCHAR (100) NOT NULL,
    ip VARCHAR (20) NOT NULL,
    deleted_at TIMESTAMP  NULL,
    penalty_at TIMESTAMP  NULL,
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    updated_at TIMESTAMP NOT NULL DEFAULT now()
);

-- +migrate Down
DROP TABLE IF EXISTS auth.node_wallet;
