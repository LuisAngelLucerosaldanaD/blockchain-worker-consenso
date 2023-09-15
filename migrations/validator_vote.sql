CREATE TABLE bc.validator_vote
(
    id              uuid         NOT NULL,
    participant_id uuid         NOT NULL,
    hash            varchar(255) NOT NULL,
    vote            bool         NOT NULL,
    created_at      timestamp    NOT NULL DEFAULT now(),
    updated_at      timestamp    NOT NULL DEFAULT now(),
    CONSTRAINT validator_votes_pkey PRIMARY KEY (id),
    CONSTRAINT fk_validator_votes_participant FOREIGN KEY (participant_id) REFERENCES bc.participant (id)
);
