package database

import (
	"admin-employee/internal/database/entities"
	"admin-employee/pkg/logs"
	"context"
	"database/sql"
)

type repository struct {
	db *sql.DB
}

type Repository interface {
	CreateDepartment(ctx context.Context, entity entities.Department) error
	CreateEmployee(ctx context.Context, entity entities.Employee) error
	DismissEmployee(ctx context.Context, id int64, dismissalDate string) error
	FetchEmployees(ctx context.Context) ([]entities.Employee, error)
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

var insertEmployee = `
	INSERT INTO employees (
		name,
		age,
		salary,
		created_at,
		hiring_date,
		department_id,
		job_title,
		active
	) VALUES (
	 	?,
		?,
		?,
		NOW(),
		?,
		?,
		?,
		?
	);
`

var selectEmployees = `
	SELECT
		id,
		name,
		age,
		salary,
		hiring_date,
		dismissal_date,
		department_id,
		job_title,
		active
	FROM employees;
`

var updateDismissalEmployee = `
	UPDATE employees
	SET dismissal_date = ?, active = false
	WHERE id = ?;
`

func (r repository) CreateDepartment(ctx context.Context, entity entities.Department) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	defer tx.Rollback()

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

func (r repository) CreateEmployee(ctx context.Context, entity entities.Employee) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	defer tx.Rollback()

	stm, err := tx.Prepare(insertEmployee)
	if err != nil {
		return err
	}

	defer stm.Close()

	_, err = stm.Exec(
		entity.Name,
		entity.Age,
		entity.Salary,
		entity.HiringDate,
		entity.DepartmentID,
		entity.JobTitle,
		entity.Active,
	)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	logs.Info.Println("1 Row inserted into table 'employees'")
	return nil
}

func (r repository) DismissEmployee(ctx context.Context, id int64, dismissalDate string) error {
	tx, err := r.db.Begin()
	if err != nil {
		return nil
	}

	stm, err := tx.Prepare(updateDismissalEmployee)
	if err != nil {
		return nil
	}

	defer stm.Close()

	_, err = stm.Exec(dismissalDate, id)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	logs.Info.Println("1 Row updated into table 'employees'")
	return nil
}

func (r repository) FetchEmployees(ctx context.Context) ([]entities.Employee, error) {
	results, err := r.db.QueryContext(ctx, selectEmployees)
	if err != nil {
		return []entities.Employee{}, nil
	}

	var employees []entities.Employee
	for results.Next() {
		var e entities.Employee
		err = results.Scan(
			&e.ID,
			&e.Name,
			&e.Age,
			&e.Salary,
			&e.HiringDate,
			&e.DismissalDate,
			&e.DepartmentID,
			&e.JobTitle,
			&e.Active,
		)
		if err != nil {
			return []entities.Employee{}, err
		}

		employees = append(employees, e)
	}
	return employees, nil
}
