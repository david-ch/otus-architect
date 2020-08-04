CREATE TABLE orders (
    id VARCHAR PRIMARY KEY,
    user_id INTEGER,
    product_id INTEGER,
    status VARCHAR,
    delivery_timeslot TIMESTAMP,
    created_by_request_id VARCHAR UNIQUE
);

commit;
