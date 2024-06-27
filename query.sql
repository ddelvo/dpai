-- Soal 1
SELECT users.name, SUM(orders.amount) AS total FROM users
INNER JOIN orders ON users.id = orders.user_id
WHERE orders.created_at >= '2022-01-01'
GROUP BY users.name
HAVING SUM(orders.amount) >= 1000;

-- Soal 2
CREATE EXTENSION IF NOT EXISTS postgres_fdw;

CREATE SERVER first_db
    FOREIGN DATA WRAPPER postgres_fdw
    OPTIONS (host 'localhost', dbname 'first', port '49153');

ALTER SERVER first_db
    OPTIONS (SET port '5432');

CREATE USER MAPPING FOR current_user
    SERVER first_db
    OPTIONS (user 'postgres', password 'postgrespw');

CREATE FOREIGN TABLE users (
    id INT,
    name VARCHAR(255),
    email VARCHAR(255)
    )
    SERVER first_db
    OPTIONS (schema_name 'public', table_name 'users');

SELECT
    users.name AS user_name,
    SUM(orders.amount) AS total_amount,
    MAX(orders.created_at) AS last_order_created_at
FROM users JOIN orders ON users.id = orders.user_id
GROUP BY users.name ORDER BY users.name;