package service

import (
	"admin-employee/internal/database"
	"admin-employee/internal/database/entities"
	"strconv"
)

type HRService interface {
	CreateDepartment(
		name string,
		description string,
		active bool,
	) error

	CreateEmployee(
		name string,
		age string,
		salary string,
		hiringDate string,
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

	if err := s.repo.InsertEmployee(entity); err != nil {
		return err
	}
	return nil
}
