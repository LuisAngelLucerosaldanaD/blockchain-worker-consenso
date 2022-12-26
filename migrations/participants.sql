-- bc.participants definition

-- Drop table

-- DROP TABLE bc.participants;

CREATE TABLE bc.participants (
                                 id uuid NOT NULL,
                                 lottery_id uuid NOT NULL,
                                 wallet_id uuid NOT NULL,
                                 amount int8 NOT NULL,
                                 accepted bool NOT NULL,
                                 type_charge int4 NOT NULL,
                                 returned bool NOT NULL,
                                 created_at timestamp NOT NULL DEFAULT now(),
                                 updated_at timestamp NOT NULL DEFAULT now(),
                                 CONSTRAINT participants_pkey PRIMARY KEY (id)
);


-- bc.participants foreign keys

ALTER TABLE bc.participants ADD CONSTRAINT fk_participants_lottery FOREIGN KEY (lottery_id) REFERENCES bc.lotteries(id);
ALTER TABLE bc.participants ADD CONSTRAINT fk_participants_type_charge FOREIGN KEY (type_charge) REFERENCES cfg.dictionaries(id);
ALTER TABLE bc.participants ADD CONSTRAINT fk_participants_wallets FOREIGN KEY (wallet_id) REFERENCES auth.wallets(id);
