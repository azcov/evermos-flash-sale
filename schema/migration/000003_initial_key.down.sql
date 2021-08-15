BEGIN;

DROP UNIQUE INDEX IF EXISTS uniq_trx_no ON orders(trx_no);

ALTER TABLE orders DROP CONSTRAINT IF EXISTS order_user_id_fkey;

ALTER TABLE order_products DROP CONSTRAINT IF EXISTS order_products_product_fkey;

COMMIT;
