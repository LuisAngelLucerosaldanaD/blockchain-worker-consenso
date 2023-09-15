-- bc.miner_response definition

-- Drop table

-- DROP TABLE bc.miner_response;

CREATE TABLE bc.miner_response
(
    id              uuid         NOT NULL,
    participant_id uuid         NOT NULL,
    hash            varchar(255) NOT NULL,
    status          int4         NOT NULL,
    nonce           int8         NOT NULL,
    difficulty      int4         NOT NULL,
    created_at      timestamp    NOT NULL DEFAULT now(),
    updated_at      timestamp    NOT NULL DEFAULT now(),
    CONSTRAINT miner_response_pkey PRIMARY KEY (id),
    CONSTRAINT fk_miner_response_participant FOREIGN KEY (participant_id) REFERENCES bc.participant (id),
    CONSTRAINT fk_miner_response_process FOREIGN KEY (status) REFERENCES cfg.dictionaries (id)
);

drop table bc.miner_response;
