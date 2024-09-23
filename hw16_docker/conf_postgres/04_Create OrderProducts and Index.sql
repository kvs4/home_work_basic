-- Table: shop.OrderProducts

-- DROP TABLE IF EXISTS shop."OrderProducts";

CREATE TABLE IF NOT EXISTS shop."OrderProducts"
(
    id uuid NOT NULL DEFAULT gen_random_uuid(),
    order_id uuid,
    product_id uuid,
    CONSTRAINT "OrderProducts_pkey" PRIMARY KEY (id)
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS shop."OrderProducts"
    OWNER to postgres;
-- Index: idx_order_id

-- DROP INDEX IF EXISTS shop.idx_order_id;

CREATE INDEX IF NOT EXISTS idx_order_id
    ON shop."OrderProducts" USING btree
    (order_id ASC NULLS LAST)
    WITH (deduplicate_items=False)
    TABLESPACE pg_default;
-- Index: idx_product_id

-- DROP INDEX IF EXISTS shop.idx_product_id;

CREATE INDEX IF NOT EXISTS idx_product_id
    ON shop."OrderProducts" USING btree
    (product_id ASC NULLS LAST)
    WITH (deduplicate_items=False)
    TABLESPACE pg_default;