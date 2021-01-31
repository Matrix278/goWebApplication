# Wallester test
Web application having CRUD operations under customer object

### Technologies:

- Go language
- PostgreSQL DB
- HTML
- CSS
- JS
- Bootstrap 4

### Libraries:

- github.com/gorilla/mux
- github.com/lib/pq

### PSQL DataBase data:

- host: localhost
- port: 5432
- user: martin
- password: qweasdzxc
- DB name: application

### Table data:

```
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
```

1. (MODELS folder) The **models** package will store the database schema. Using **struct** type to represent or map the database schema in golang.

2. (MIDDLEWARE folder) The **middleware** package is the bridge between APIs and Database. This package will handle all the db operations like Insert, Select, Update, and Delete (CRUD).

3. (ROUTER folder) In the **router** package defines all the api endpoints.

4. (MAIN file) The **main.go** is server. It will start a server on **8080** port and serve all the **Router**.

### Routes:

1. localhost:8080/api/customer/{id}, GET CUSTOMER ( GET )
2. localhost:8080/api/customer, GET ALL CUSTOMERS ( GET )
3. localhost:8080/api/newcustomer, CREATE CUSTOMER ( POST )
4. localhost:8080/api/customer/{id}, UPDATE CUSTOMER BY ID ( PUT )
5. localhost:8080/api/deletecustomer/{id}, DELETE CUSTOMER BY ID ( DELETE )
