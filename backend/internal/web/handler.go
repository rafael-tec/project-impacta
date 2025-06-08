package web

import (
	"encoding/json"
	"net/http"
	"strconv"

	"admin-employee/internal/service"
	"admin-employee/pkg/logs"
)

type EmployeeHandler struct {
	service service.HRService
}

func NewEmployeeHandler(service service.HRService) EmployeeHandler {
	return EmployeeHandler{service: service}
}

func (h EmployeeHandler) CreateDepartment(w http.ResponseWriter, r *http.Request) {
	logs.Info.Println("Received request for POST '/department'")

	var request CreateDepartmentRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		logs.Error.Println("Unmarshal request body failed")
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	err = h.service.CreateDepartment(
		r.Context(),
		request.Name,
		request.Description,
		request.Active,
	)
	if err != nil {
		logs.Error.Println("Failure to create the department")
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(request)
}

func (h EmployeeHandler) CreateEmployee(w http.ResponseWriter, r *http.Request) {
	logs.Info.Println("Received request for POST '/employee'")

	var request CreateEmployeeRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		logs.Error.Println("Unmarshal request body failed")
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	err = h.service.CreateEmployee(
		r.Context(),
		request.Name,
		request.Age,
		request.Salary,
		request.HiringDate,
		request.DepartmentID,
		request.JobTitle,
		request.Active,
	)
	if err != nil {
		logs.Error.Println("Failure to create the employee")
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(request)
}

func (h EmployeeHandler) PatchEmployee(w http.ResponseWriter, r *http.Request) {
	logs.Info.Println("Received request for PATCH '/employee'")

	dismissalDate := r.URL.Query().Get("dismissal_date")
	if dismissalDate == "" {
		http.Error(w, "dismissal_date is required", http.StatusBadRequest)
	}

	employeeID := r.URL.Query().Get("employee_id")
	if dismissalDate == "" {
		http.Error(w, "employee_id is required", http.StatusBadRequest)
	}

	employeeIDInt, err := strconv.ParseInt(employeeID, 10, 64)
	if err != nil {
		http.Error(w, "employee_id is invalid number", http.StatusBadRequest)
	}

	err = h.service.DismissEmployee(r.Context(), employeeIDInt, dismissalDate)
	if err != nil {
		logs.Error.Println("Failure to create the employee")
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(nil)
}

func (h EmployeeHandler) GetEmployees(w http.ResponseWriter, r *http.Request) {
	logs.Info.Println("Received request for GET '/employees'")

	employees, err := h.service.GetEmployees(r.Context())
	if err != nil {
		logs.Error.Println("Failure to fetch employees")
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	var responseEmployees []GetEmployeeResponse
	for _, employee := range employees {
		response := GetEmployeeResponse{
			ID:            employee.ID,
			Name:          employee.Name,
			Age:           employee.Age,
			Salary:        employee.Salary,
			HiringDate:    employee.HiringDate,
			DismissalDate: employee.DismissalDate,
			DepartmentID:  employee.DepartmentID,
			JobTitle:      employee.JobTitle,
			Active:        employee.Active,
		}
		responseEmployees = append(responseEmployees, response)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(responseEmployees)
}
