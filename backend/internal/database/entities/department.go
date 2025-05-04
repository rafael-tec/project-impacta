package entities

type Department struct {
	Name        string
	Description string
	CratedAt    string
	Active      bool
}

type Employee struct {
	Name          string
	Age           int
	Salary        float64
	HiringDate    string
	DismissalDate string
	DepartmentID  string
	JobTitle      string
	Active        bool
}
