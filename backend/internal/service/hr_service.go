package service

import (
	"admin-employee/internal/database"
	"admin-employee/internal/database/entities"
)

type HRService interface {
	CreateDepartment(
		name string,
		description string,
		active bool,
	) error

	CreateEmployee(
		name string,
		age int,
		salary float64,
		hiringDate string,
		dismissalDate string,
		departmentID string,
		jobTitle string,
		active bool,
	) error
}

type hrService struct {
	repo database.Repository
}

func NewHRService(repo database.Repository) hrService {
	return hrService{repo: repo}
}

func (s hrService) CreateDepartment(
	name string,
	description string,
	active bool,
) error {
	entity := entities.Department{
		Name:        name,
		Description: description,
		Active:      active,
	}

	if err := s.repo.InsertDepartment(entity); err != nil {
		return err
	}
	return nil
}

func (s hrService) CreateEmployee(
	name string,
	age int,
	salary float64,
	hiringDate string,
	dismissalDate string,
	departmentID string,
	jobTitle string,
	active bool,
) error {
	entity := entities.Employee{
		Name:          name,
		Age:           age,
		Salary:        salary,
		HiringDate:    hiringDate,
		DismissalDate: dismissalDate,
		DepartmentID:  departmentID,
		JobTitle:      jobTitle,
		Active:        active,
	}

	if err := s.repo.InsertEmployee(entity); err != nil {
		return err
	}
	return nil
}
