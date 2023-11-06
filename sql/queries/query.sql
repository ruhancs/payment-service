-- name: CreateOrder :exec
INSERT INTO orders (id,amount,plan,customer_id,first_name,last_name,email,status,transaction_id,created_at) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10);

-- name: GetOrders :one
SELECT * FROM orders WHERE id = $1 LIMIT 1;

-- name: UpdateOrderStatus :one
UPDATE orders SET status = $2, transaction_id = $3  WHERE id = $1 RETURNING *;

-- name: CreateTransaction :one
INSERT INTO transactions (
    id,
    amount,
    currency,
    payment_intent,
    payment_method,
    expire_month,
    expire_year,
    transaction_status,
    created_at,
    updated_at
    ) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10) RETURNING *;

-- name: CreateCustomer :exec
INSERT INTO customers (id,first_name,last_name,email,is_active) VALUES ($1,$2,$3,$4,$5);

-- name: GetCustomerByEmail :one
SELECT * FROM customers WHERE email = $1 LIMIT 1;

-- name: GetCustomerById :one
SELECT * FROM customers WHERE id = $1 LIMIT 1;