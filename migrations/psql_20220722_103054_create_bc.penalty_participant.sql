
-- +migrate Up
CREATE TABLE IF NOT EXISTS bc.penalty_participant(
    id uuid NOT NULL PRIMARY KEY,
    participant_id UUID  NOT NULL,
    amount float8  NOT NULL,
    penalty_percentage float8  NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    updated_at TIMESTAMP NOT NULL DEFAULT now(),
    constraint FK_participant_penalty foreign key(participant_id) references bc.participant(id)
);

-- +migrate Down
DROP TABLE IF EXISTS bc.penalty_participant;
