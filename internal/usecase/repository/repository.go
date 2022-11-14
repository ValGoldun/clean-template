package repository

import (
	"context"
	"database/sql"
	"github.com/ValGoldun/clean-template/internal/entity"
)

type Repository struct {
	db *sql.DB
}

func New(db *sql.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r Repository) GetUsers(ctx context.Context) ([]entity.User, error) {
	rows, err := r.db.QueryContext(ctx, "SELECT id, full_name, birth_date FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	entities := make([]entity.User, 0)

	for rows.Next() {
		e := entity.User{}

		err = rows.Scan(&e.ID, &e.FullName, &e.Birthdate)
		if err != nil {
			return nil, err
		}

		entities = append(entities, e)
	}

	return entities, nil
}
