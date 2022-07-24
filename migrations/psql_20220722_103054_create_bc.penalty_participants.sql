
-- +migrate Up
CREATE TABLE IF NOT EXISTS bc.penalty_participants(
    id uuid NOT NULL PRIMARY KEY,
    lottery_id UUID  NOT NULL,
    participants_id UUID  NOT NULL,
    amount float8  NOT NULL,
    penalty_percentage float8  NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    updated_at TIMESTAMP NOT NULL DEFAULT now()
);

-- +migrate Down
DROP TABLE IF EXISTS bc.penalty_participants;
