-- bc.lotteries definition

-- Drop table

-- DROP TABLE bc.lotteries;

CREATE TABLE bc.lotteries (
                              id uuid NOT NULL,
                              block_id int8 NOT NULL,
                              registration_start_date timestamp NOT NULL,
                              registration_end_date timestamp NULL,
                              lottery_start_date timestamp NULL,
                              lottery_end_date timestamp NULL,
                              process_end_date timestamp NULL,
                              process_status int4 NOT NULL,
                              created_at timestamp NOT NULL DEFAULT now(),
                              updated_at timestamp NOT NULL DEFAULT now(),
                              CONSTRAINT lotteries_pkey PRIMARY KEY (id)
);


-- bc.lotteries foreign keys

ALTER TABLE bc.lotteries ADD CONSTRAINT fk_lotteries_process FOREIGN KEY (process_status) REFERENCES cfg.dictionaries(id);
