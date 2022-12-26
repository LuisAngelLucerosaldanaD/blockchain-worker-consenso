-- bc.rewards definition

-- Drop table

-- DROP TABLE bc.rewards;

CREATE TABLE bc.rewards (
                            id uuid NOT NULL,
                            lottery_id uuid NOT NULL,
                            id_wallet uuid NOT NULL,
                            amount int8 NOT NULL,
                            created_at timestamp NOT NULL DEFAULT now(),
                            updated_at timestamp NOT NULL DEFAULT now(),
                            CONSTRAINT rewards_pkey PRIMARY KEY (id)
);


-- bc.rewards foreign keys

ALTER TABLE bc.rewards ADD CONSTRAINT fk_rewards_lottery FOREIGN KEY (lottery_id) REFERENCES bc.lotteries(id);
ALTER TABLE bc.rewards ADD CONSTRAINT fk_rewards_wallets FOREIGN KEY (id_wallet) REFERENCES auth.wallets(id);
