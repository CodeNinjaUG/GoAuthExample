-- name: CreateUser :one
INSERT INTO users (
    first_name,
    last_name,
    phone,
    email,
    password
 ) VALUES (
  $1, $2, $3, $4, $5
 ) RETURNING *;

 -- name: GetUser :one
SELECT * FROM users
WHERE id = $1 LIMIT 1;

-- name: ListUsers :many
SELECT * FROM users
ORDER BY name LIMIT $1
OFFSET $2;

-- name: UpdateUser :one
UPDATE users
  set 
  first_name= $2,
  last_name = $3,
  phone= $4,
  email=$5
WHERE id = $1
RETURNING *;

-- name: DeleteUser :exec
DELETE FROM users WHERE id= $1;