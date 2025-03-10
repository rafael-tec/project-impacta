package database

import (
	"admin-employee/internal/database/entities"
	"admin-employee/pkg/logs"
	"database/sql"
)

type repository struct {
	db *sql.DB
}

type Repository interface {
	InsertDepartment(entity entities.Department) error
}

func NewRepository(db *sql.DB) Repository {
	return repository{db: db}
}

var insertDepartment = `
	INSERT INTO departments (
		name,
		description,
		created_at,
		active
	) VALUES (
		?,
		?,
		NOW(),
		?
	);
`

func (r repository) InsertDepartment(entity entities.Department) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	stm, err := tx.Prepare(insertDepartment)
	if err != nil {
		return err
	}

	defer stm.Close()

	_, err = stm.Exec(
		entity.Name,
		entity.Description,
		entity.Active,
	)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	logs.Info.Println("1 Row inserted into table 'departments'")
	return nil
}
