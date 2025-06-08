package web

type CreateDepartmentRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Active      bool   `json:"active"`
}

type CreateEmployeeRequest struct {
	Name          string `json:"name"`
	Age           string `json:"age"`
	Salary        string `json:"salary"`
	HiringDate    string `json:"hiring_date"`
	DismissalDate string `json:"dismissal_date"`
	DepartmentID  string `json:"department_id"`
	JobTitle      string `json:"job_title"`
	Active        bool   `json:"active"`
}

type PatchEmployeeRequest struct {
	ID            int64  `json:"id"`
	DismissalDate string `json:"dismissal_date"`
}

type GetEmployeeResponse struct {
	ID            int64   `json:"id"`
	Name          string  `json:"name"`
	Age           int64   `json:"age"`
	Salary        float64 `json:"salary"`
	HiringDate    string  `json:"hiring_date"`
	DismissalDate *string `json:"dismissal_date"`
	DepartmentID  string  `json:"department_id"`
	JobTitle      string  `json:"job_title"`
	Active        bool    `json:"active"`
}
