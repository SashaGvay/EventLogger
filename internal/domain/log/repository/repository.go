package repository

import pgx "github.com/jackc/pgx/v5"

type Repository struct {
	db *pgx.Conn
}

func New(db *pgx.Conn) *Repository {
	return &Repository{
		db: db,
	}
}
