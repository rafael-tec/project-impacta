package main

import (
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	"admin-employee/internal/database"
	"admin-employee/internal/service"
	"admin-employee/internal/web"
	"admin-employee/pkg/logs"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	dbConfig, err := database.NewDBConfig()
	if err != nil {
		panic(err)
	}

	sqlDB, err := database.ConnectDB(*dbConfig)
	if err != nil {
		panic(err)
	}

	repository := database.NewRepository(sqlDB)
	service := service.NewHRService(repository)
	handler := web.NewEmployeeHandler(service)

	router := mux.NewRouter()
	router.HandleFunc("/department", handler.CreateDepartment).Methods("POST")
	router.HandleFunc("/employee", handler.CreateEmployee).Methods("POST")
	router.HandleFunc("/employee", handler.PatchEmployee).Methods("PATCH")
	router.HandleFunc("/employees", handler.GetEmployees).Methods("GET")

	corsOptions := handlers.AllowedOrigins([]string{"http://localhost:3000"})
	corsMethods := handlers.AllowedMethods([]string{"POST", "GET", "PUT", "PATCH"})
	corsHeaders := handlers.AllowedHeaders([]string{"Content-Type"})

	logs.Init()
	logs.Info.Println("Running server on http://localhost:8080/")
	http.ListenAndServe(":8080", handlers.CORS(corsOptions, corsMethods, corsHeaders)(router))
}
