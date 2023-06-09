// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.13.0
// source: user.sql

package db

import (
	"context"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users (
    first_name,
    last_name,
    phone,
    email,
    password
 ) VALUES (
  $1, $2, $3, $4, $5
 ) RETURNING id, first_name, last_name, phone, email, password, created_at, updated_at
`

type CreateUserParams struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Phone     string `json:"phone"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.queryRow(ctx, q.createUserStmt, createUser,
		arg.FirstName,
		arg.LastName,
		arg.Phone,
		arg.Email,
		arg.Password,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.FirstName,
		&i.LastName,
		&i.Phone,
		&i.Email,
		&i.Password,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteUser = `-- name: DeleteUser :exec
DELETE FROM users WHERE id= $1
`

func (q *Queries) DeleteUser(ctx context.Context, id int32) error {
	_, err := q.exec(ctx, q.deleteUserStmt, deleteUser, id)
	return err
}

const getUser = `-- name: GetUser :one
SELECT id, first_name, last_name, phone, email, password, created_at, updated_at FROM users
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetUser(ctx context.Context, id int32) (User, error) {
	row := q.queryRow(ctx, q.getUserStmt, getUser, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.FirstName,
		&i.LastName,
		&i.Phone,
		&i.Email,
		&i.Password,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const listUsers = `-- name: ListUsers :many
SELECT id, first_name, last_name, phone, email, password, created_at, updated_at FROM users
ORDER BY name LIMIT $1
OFFSET $2
`

type ListUsersParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListUsers(ctx context.Context, arg ListUsersParams) ([]User, error) {
	rows, err := q.query(ctx, q.listUsersStmt, listUsers, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []User{}
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.ID,
			&i.FirstName,
			&i.LastName,
			&i.Phone,
			&i.Email,
			&i.Password,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateUser = `-- name: UpdateUser :one
UPDATE users
  set 
  first_name= $2,
  last_name = $3,
  phone= $4,
  email=$5
WHERE id = $1
RETURNING id, first_name, last_name, phone, email, password, created_at, updated_at
`

type UpdateUserParams struct {
	ID        int32  `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Phone     string `json:"phone"`
	Email     string `json:"email"`
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) (User, error) {
	row := q.queryRow(ctx, q.updateUserStmt, updateUser,
		arg.ID,
		arg.FirstName,
		arg.LastName,
		arg.Phone,
		arg.Email,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.FirstName,
		&i.LastName,
		&i.Phone,
		&i.Email,
		&i.Password,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
