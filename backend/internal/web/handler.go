package web

import (
	"encoding/json"
	"net/http"

	"admin-employee/internal/service"
	"admin-employee/pkg/logs"
)

type CreateDepartmentRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Active      bool   `json:"active"`
}

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
