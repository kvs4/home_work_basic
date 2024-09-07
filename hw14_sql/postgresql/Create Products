-- Table: shop.Products

-- DROP TABLE IF EXISTS shop."Products";

CREATE TABLE IF NOT EXISTS shop."Products"
(
    id uuid NOT NULL DEFAULT gen_random_uuid(),
    name character varying(200) COLLATE pg_catalog."default",
    price numeric(15,2),
    CONSTRAINT "Products_pkey" PRIMARY KEY (id)
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS shop."Products"
    OWNER to postgres;