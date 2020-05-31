CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username VARCHAR,
    firstname VARCHAR,
    lastname VARCHAR,
    email VARCHAR,
    phone VARCHAR
);

INSERT INTO users (username, firstname, lastname, email, phone) 
    VALUES ('ivan42', 'ivan', 'ivanov', 'superivan@some.com', '111-22-33');

commit;
