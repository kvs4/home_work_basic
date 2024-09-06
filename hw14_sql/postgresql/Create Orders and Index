-- Table: shop.Orders

-- DROP TABLE IF EXISTS shop."Orders";

CREATE TABLE IF NOT EXISTS shop."Orders"
(
    id uuid NOT NULL DEFAULT gen_random_uuid(),
    user_id uuid,
    order_date timestamp with time zone NOT NULL,
    total_amount numeric(15,2),
    CONSTRAINT "Orders_pkey" PRIMARY KEY (id),
    CONSTRAINT user_id FOREIGN KEY (user_id)
        REFERENCES shop."Users" (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
        NOT VALID
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS shop."Orders"
    OWNER to postgres;
-- Index: idx_user_id

-- DROP INDEX IF EXISTS shop.idx_user_id;

CREATE INDEX IF NOT EXISTS idx_user_id
    ON shop."Orders" USING btree
    (user_id ASC NULLS LAST)
    WITH (deduplicate_items=False)
    TABLESPACE pg_default;