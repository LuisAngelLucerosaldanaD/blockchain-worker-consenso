-- bc.miner_response definition

-- Drop table

-- DROP TABLE bc.miner_response;

CREATE TABLE bc.miner_response (
                                   id uuid NOT NULL,
                                   lottery_id uuid NOT NULL,
                                   participants_id uuid NOT NULL,
                                   hash varchar(255) NOT NULL,
                                   status int4 NOT NULL,
                                   created_at timestamp NOT NULL DEFAULT now(),
                                   updated_at timestamp NOT NULL DEFAULT now(),
                                   nonce int8 NOT NULL,
                                   difficulty int4 NOT NULL,
                                   CONSTRAINT miner_response_pkey PRIMARY KEY (id)
);


-- bc.miner_response foreign keys

ALTER TABLE bc.miner_response ADD CONSTRAINT fk_miner_response_lottery FOREIGN KEY (lottery_id) REFERENCES bc.lotteries(id);
ALTER TABLE bc.miner_response ADD CONSTRAINT fk_miner_response_participants FOREIGN KEY (participants_id) REFERENCES bc.participants(id);
ALTER TABLE bc.miner_response ADD CONSTRAINT fk_miner_response_process FOREIGN KEY (status) REFERENCES cfg.dictionaries(id);
