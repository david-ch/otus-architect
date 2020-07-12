CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username VARCHAR,
    firstname VARCHAR,
    lastname VARCHAR,
    passwordhash VARCHAR,
    personalstatus VARCHAR
);

commit;
