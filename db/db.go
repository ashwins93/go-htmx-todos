package db

import "github.com/jmoiron/sqlx"

type DB struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) *DB {
	return &DB{db}
}
