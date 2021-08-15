BEGIN;

CREATE TABLE IF NOT EXISTS master_counters (
  id serial NOT NULL PRIMARY KEY,
  name text NOT NULL,
  counter int NOT NULL DEFAULT 1,
  prefix text NOT NULL,
  created_at bigint NOT NULL,
  updated_at bigint NOT NULL
);

CREATE TABLE IF NOT EXISTS users (
  id serial NOT NULL PRIMARY KEY,
  name text NOT NULL,
  phone text,
  email text NOT NULL,
  created_at bigint NOT NULL
);

CREATE TABLE IF NOT EXISTS products (
  id serial NOT NULL PRIMARY KEY,
  name text NOT NULL,
  price bigint NOT NULL,
  qty int NOT NULL DEFAULT 0,
  created_at bigint NOT NULL
);

CREATE TABLE IF NOT EXISTS orders (
   id serial NOT NULL PRIMARY KEY,
   order_no text NOT NULL,
   user_id int NOT NULL,
   total_price bigint NOT NULL,
   created_at bigint NOT NULL
);

CREATE TABLE IF NOT EXISTS order_products (
   order_id int NOT NULL,
   product_id int NOT NULL,
   price bigint NOT NULL,
   qty int NOT NULL DEFAULT 0,
   created_at bigint NOT NULL
);

COMMIT;
