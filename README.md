## 12.12 Flash Sale 

We are members of the engineering team of an online store. When we look at ratings for our online store application, we received the following
facts:
1. Customers were able to put items in their cart, check out, and then pay. After several days, many of our customers received calls from our Customer Service department stating that their orders have been canceled due to stock unavailability.
2. These bad reviews generally come within a week after our 12.12 event, in which we held a large flash sale and set up other major
discounts to promote our store.

After checking in with our Customer Service and Order Processing departments, we received the following additional facts:
1. Our inventory quantities are often misreported, and some items even go as far as having a negative inventory quantity.
2. The misreported items are those that performed very well on our 12.12 event.
3. Because of these misreported inventory quantities, the Order Processing department was unable to fulfill a lot of orders, and thus requested help from our Customer Service department to call our customers and notify them that we have had to cancel their orders.

## Analysis
there is no locking in the product when a product is purchased simultaneously.
for example:
Asep and Dadang checkout a product (Laptop) at same time. 
the app will check to db at the same time and will get product qty to make sure qty is enough to purchased.
```sql
SELECT qty 
FROM products
WHERE product_id = 1;
-- Laptop qty = 1
```
and the response of db will same.
it turns out that the quantity of the product is sufficient, the next step will be carried out.

## Solution
using database locking when get product to check the qty and update the product qty.
for example :
```sql
BEGIN;

SELECT qty 
FROM products
WHERE product_id = 1
FOR UPDATE;
-- qty = 1

UPDATE products
SET qty = 0
WHERE product_id = 1;

COMMIT;
```
 query FOR UPDATE is important in this case. make the selected products is locked.

---

## How To RUN
If you haven't install golang please [install](https://golang.org/doc/install) first.

 Run local:
```sh
# make sure you installed required package
$ make install
# copy app_config.example.json to app_config.json
$ cp app_config.example.json app_config.json
# make sure edit database (postgresql) credential config to yours
# then, migrating app mmigration to your postgresql
$ make migrate-up
# run the app locally
$ make local
```

Run container:
```sh
$ docker-compose up
# $ go run main.go
```