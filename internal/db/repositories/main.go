package repositories

import (
	"database/sql"
)

type MainRepository struct {
	db *sql.DB
}

func NewMainRepository(db *sql.DB) *MainRepository {
	return &MainRepository{
		db: db,
	}
}

func (r *MainRepository) Test() string {
	return "Test"
}
