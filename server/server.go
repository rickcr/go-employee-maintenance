package server

import (
	"log"
	"net/http"

	"employee-maintenance/services"
)

type Server struct {
	employeeService   *services.EmployeeService
	departmentService *services.DepartmentService
	mux               *http.ServeMux
}

func NewServer(empService *services.EmployeeService, deptService *services.DepartmentService) *Server {
	s := &Server{
		employeeService:   empService,
		departmentService: deptService,
		mux:               http.NewServeMux(),
	}
	s.registerRoutes()
	return s
}

func (s *Server) registerRoutes() {
	s.RegisterEmployeeRoutes()
	s.RegisterDepartmentRoutes()
	s.RegisterSwaggerRoutes()
}

func (s *Server) Start() {
	log.Println("Server starting on http://localhost:8080")
	log.Println("Swagger UI available at http://localhost:8080/swagger")
	if err := http.ListenAndServe(":8080", s.mux); err != nil {
		log.Fatal("Server failed to start:", err)
	}
}
