package web

type CreateDepartmentRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Active      bool   `json:"active"`
}

type CreateEmployeeRequest struct {
	Name          string  `json:"name"`
	Age           int     `json:"age"`
	Salary        float64 `json:"salary"`
	HiringDate    string  `json:"hiring_date"`
	DismissalDate string  `json:"dismissal_date"`
	DepartmentID  string  `json:"department_id"`
	JobTitle      string  `json:"job_title"`
	Active        bool    `json:"active"`
}
