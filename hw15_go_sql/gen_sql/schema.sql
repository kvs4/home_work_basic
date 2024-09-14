CREATE SCHEMA IF NOT EXISTS shop
    AUTHORIZATION postgres;

CREATE TABLE IF NOT EXISTS shop."OrderProducts"
(
    id uuid NOT NULL DEFAULT gen_random_uuid(),
    order_id uuid,
    product_id uuid,
    CONSTRAINT "OrderProducts_pkey" PRIMARY KEY (id)
);

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
);

CREATE TABLE IF NOT EXISTS shop."Products"
(
    id uuid NOT NULL DEFAULT gen_random_uuid(),
    name character varying(200) COLLATE pg_catalog."default",
    price numeric(15,2),
    CONSTRAINT "Products_pkey" PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS shop."Users"
(
    id uuid NOT NULL DEFAULT gen_random_uuid(),
    name character varying(200) COLLATE pg_catalog."default",
    email character varying(100) COLLATE pg_catalog."default",
    password character varying(50) COLLATE pg_catalog."default",
    CONSTRAINT "Users_pkey" PRIMARY KEY (id)
);