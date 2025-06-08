package service

import (
	"admin-employee/internal/database"
	"admin-employee/internal/database/entities"
	"context"
	"strconv"
)

type HRService interface {
	CreateDepartment(
		ctx context.Context,
		name string,
		description string,
		active bool,
	) error

	CreateEmployee(
		ctx context.Context,
		name string,
		age string,
		salary string,
		hiringDate string,
		departmentID string,
		jobTitle string,
		active bool,
	) error

	DismissEmployee(
		ctx context.Context,
		id int64,
		dismissalDate string,
	) error

	GetEmployees(
		ctx context.Context,
	) ([]entities.Employee, error)
}

type hrService struct {
	repo database.Repository
}

func NewHRService(repo database.Repository) hrService {
	return hrService{repo: repo}
}

func (s hrService) CreateDepartment(
	ctx context.Context,
	name string,
	description string,
	active bool,
) error {
	entity := entities.Department{
		Name:        name,
		Description: description,
		Active:      active,
	}

	if err := s.repo.CreateDepartment(ctx, entity); err != nil {
		return err
	}
	return nil
}

func (s hrService) CreateEmployee(
	ctx context.Context,
	name string,
	age string,
	salary string,
	hiringDate string,
	departmentID string,
	jobTitle string,
	active bool,
) error {
	ageInt, err := strconv.ParseInt(age, 10, 64)
	if err != nil {
		return err
	}

	salaryFloat, err := strconv.ParseFloat(salary, 64)
	if err != nil {
		return err
	}

	entity := entities.Employee{
		Name:         name,
		Age:          ageInt,
		Salary:       salaryFloat,
		HiringDate:   hiringDate,
		DepartmentID: departmentID,
		JobTitle:     jobTitle,
		Active:       active,
	}

	if err := s.repo.CreateEmployee(ctx, entity); err != nil {
		return err
	}
	return nil
}

func (s hrService) DismissEmployee(ctx context.Context, id int64, dismissalDate string) error {
	if err := s.repo.DismissEmployee(ctx, id, dismissalDate); err != nil {
		return err
	}
	return nil
}

func (s hrService) GetEmployees(ctx context.Context) ([]entities.Employee, error) {
	employees, err := s.repo.FetchEmployees(ctx)
	if err != nil {
		return []entities.Employee{}, err
	}
	return employees, nil
}
