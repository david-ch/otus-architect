CREATE TABLE stock
(
    product_id INTEGER,
    amount     INTEGER
);

CREATE TABLE stock_reservation
(
    product_id INTEGER,
    order_id   VARCHAR,
    amount     INTEGER
);

INSERT INTO stock (product_id, amount)
VALUES (42, 50);

commit;
