-- bc.participant definition

-- Drop table

-- DROP TABLE bc.participant;

CREATE TABLE bc.participant
(
    id          uuid      NOT NULL,
    lottery_id  uuid      NOT NULL,
    wallet_id   uuid      NOT NULL,
    amount      int8      NOT NULL,
    accepted    bool      NOT NULL,
    type_charge int4      NOT NULL,
    returned    bool      NOT NULL,
    created_at  timestamp NOT NULL DEFAULT now(),
    updated_at  timestamp NOT NULL DEFAULT now(),
    CONSTRAINT participant_pkey PRIMARY KEY (id),
    CONSTRAINT fk_participants_lottery FOREIGN KEY (lottery_id) REFERENCES bc.lottery (id),
    CONSTRAINT fk_participants_type_charge FOREIGN KEY (type_charge) REFERENCES cfg.dictionaries (id),
    CONSTRAINT fk_participants_wallets FOREIGN KEY (wallet_id) REFERENCES auth.wallet (id)
);
