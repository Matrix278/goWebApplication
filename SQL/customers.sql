DROP TABLE IF EXISTS customers;
CREATE TABLE customers (
    id SERIAL PRIMARY KEY,
    first_name VARCHAR(100) NOT NULL,
    last_name VARCHAR(100) NOT NULL,
    birth_date TIMESTAMP NOT NULL,
    gender VARCHAR(10) NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    address VARCHAR(200)
);
INSERT INTO customers(
        first_name,
        last_name,
        birth_date,
        gender,
        email,
        address
    )
VALUES (
        'Test',
        'Tester',
        '2016-06-22 19:10:25',
        'Male',
        'test@tester.com',
        'Test City address'
    );