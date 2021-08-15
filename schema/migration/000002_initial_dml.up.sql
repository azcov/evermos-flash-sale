BEGIN;

INSERT INTO master_counters (id, name,  prefix, created_at, updated_at)
VALUES (1, 'order',  'INV', extract(epoch from now()), extract(epoch from now()) );

INSERT INTO users (name,  phone, email, created_at)
VALUES ('asep',  '628132000001', 'asep@gmail.com',extract(epoch from now()) ),
        ('dadang',  '628576000002', 'dadang@yahoo.com',extract(epoch from now()) );

INSERT INTO products (name,  price, qty, created_at)
VALUES ('handphone',  1999000, 2,extract(epoch from now()) ),
       ('tv',  2499000, 1,extract(epoch from now()) ),
       ('baju',  189000, 5,extract(epoch from now()) );

COMMIT;
