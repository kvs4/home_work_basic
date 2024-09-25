-- name: CreateUser :one
INSERT INTO Shop."Users" (name, email, password)
VALUES ($1, $2, $3)
returning id, name;

-- name: GetUser :one
SELECT id, name, email, password
	FROM shop."Users"
WHERE name = $1;

-- name: GetUsers :many
SELECT id, name, email, password
	FROM shop."Users";

-- name: DeleteUser :one
DELETE FROM shop."Users" WHERE name = $1
returning id, name;

-- name: ChangeEmailUser :one
UPDATE shop."Users" SET email = $1
WHERE name = $2
returning id, name, email;


-- name: CreateProduct :one
INSERT INTO shop."Products" (name, price)
VALUES ($1, $2)
returning id, name;

-- name: GetProduct :one
SELECT id, name, price
	FROM shop."Products"
WHERE name = $1;

-- name: GetProducts :many
SELECT id, name, price
	FROM shop."Products";

-- name: DeleteProduct :one
DELETE FROM shop."Products" WHERE name = $1
returning id, name;

-- name: ChangePriceProduct :one
UPDATE shop."Products" SET price = $1
WHERE name = $2
returning id, name, price;


-- name: CreateOrder :one
INSERT INTO Shop."Orders" (user_id, order_date, total_amount)
VALUES ($1, $2, $3)
returning id, user_id, order_date, total_amount;

-- name: GetUsersOrders :many
SELECT 
	Users.name as user_name, 
	Users.email as user_email, 
	Orders.id as Order_id, 
	Orders.order_date, 
	Orders.total_amount, 
	Prod.name as product, 
	Prod.price
	FROM shop."Orders" Orders
	LEFT JOIN shop."Users" Users ON Orders.user_id = Users.id
	LEFT JOIN shop."OrderProducts" OrderProd ON Orders.id = OrderProd.order_id
		LEFT JOIN shop."Products" Prod ON OrderProd.product_id = Prod.id;

-- name: GetUserStatistics :one
WITH AvgPriceInOrder AS (
SELECT 
	Users.name as user_name, 
	Users.email as user_email,  
	Orders.total_amount, 
	AVG(Prod.price) as average_price
	FROM shop."Orders" Orders
	LEFT JOIN shop."Users" Users ON Orders.user_id = Users.id
	LEFT JOIN shop."OrderProducts" OrderProd ON Orders.id = OrderProd.order_id
		LEFT JOIN shop."Products" Prod ON OrderProd.product_id = Prod.id
	WHERE Users."name" = $1
	GROUP BY 
		user_name, 
		user_email, 
		Orders.total_amount 
		
	)
	
SELECT
	user_name,
	user_email, 
	SUM (total_amount) as sum_amount,
	AVG(average_price) as average_price
	
	FROM AvgPriceInOrder
GROUP BY 
		user_name, 
		user_email;

-- name: DeleteOrder :one
delete from shop."Orders" where id = $1
returning id, user_id, order_date, total_amount;


-- name: CreateOrderProducts :one
INSERT INTO shop."OrderProducts" (order_id, product_id)
VALUES ($1, $2)
returning id, order_id, product_id;

-- name: DeleteOrderProducts :one
delete from shop."OrderProducts" where order_id = $1
returning id, order_id, product_id;


