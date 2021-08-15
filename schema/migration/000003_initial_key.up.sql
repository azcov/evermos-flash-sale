BEGIN;

CREATE UNIQUE INDEX IF NOT EXISTS uniq_trx_no ON orders(order_no);

ALTER TABLE orders ADD CONSTRAINT order_user_id_fkey FOREIGN KEY (user_id) REFERENCES users (id);

ALTER TABLE order_products ADD CONSTRAINT order_products_product_fkey FOREIGN KEY (product_id) REFERENCES products (id);

COMMIT;
