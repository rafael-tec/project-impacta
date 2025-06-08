package entities

type Employee struct {
	ID            int64
	Name          string
	Age           int64
	Salary        float64
	HiringDate    string
	DismissalDate *string
	DepartmentID  string
	JobTitle      string
	Active        bool
}
