// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: users.sql

package database

import (
	"context"
	"database/sql"
)

const createAuthor = `-- name: CreateAuthor :execresult
CREATE TABLE Users (
    id INT AUTO_INCREMENT PRIMARY KEY,
    Firstname VARCHAR(255) NOT NULL,
    Lastname VARCHAR(255) NOT NULL,
    Gender VARCHAR(10) NOT NULL,
    StudentId VARCHAR(50) NOT NULL,
    Course VARCHAR(255) NOT NULL,
    Level VARCHAR(50) NOT NULL,
    Email VARCHAR(255) NOT NULL
)
`

func (q *Queries) CreateAuthor(ctx context.Context) (sql.Result, error) {
	return q.db.ExecContext(ctx, createAuthor)
}