-- bc.validator_votes definition

-- Drop table

-- DROP TABLE bc.validator_votes;

CREATE TABLE bc.validator_votes (
                                    id uuid NOT NULL,
                                    lottery_id uuid NOT NULL,
                                    participants_id uuid NOT NULL,
                                    hash varchar(255) NOT NULL,
                                    vote bool NOT NULL,
                                    created_at timestamp NOT NULL DEFAULT now(),
                                    updated_at timestamp NOT NULL DEFAULT now(),
                                    CONSTRAINT validator_votes_pkey PRIMARY KEY (id)
);


-- bc.validator_votes foreign keys

ALTER TABLE bc.validator_votes ADD CONSTRAINT fk_validator_votes_lottery FOREIGN KEY (lottery_id) REFERENCES bc.lotteries(id);
ALTER TABLE bc.validator_votes ADD CONSTRAINT fk_validator_votes_participants FOREIGN KEY (participants_id) REFERENCES bc.participants(id);
