-- Table: shop.Users

-- DROP TABLE IF EXISTS shop."Users";

CREATE TABLE IF NOT EXISTS shop."Users"
(
    id uuid NOT NULL DEFAULT gen_random_uuid(),
    name character varying(200) COLLATE pg_catalog."default",
    email character varying(100) COLLATE pg_catalog."default",
    password character varying(50) COLLATE pg_catalog."default",
    CONSTRAINT "Users_pkey" PRIMARY KEY (id)
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS shop."Users"
    OWNER to postgres;