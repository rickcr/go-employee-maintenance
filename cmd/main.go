package main

import (
	_ "embed"

	"employee-maintenance/server"
	"employee-maintenance/services"
)

//go:embed openapi.yaml
var openapiSpec []byte

func main() {
	server.SetOpenAPISpec(openapiSpec)

	employeeService := services.NewEmployeeService()
	departmentService := services.NewDepartmentService()
	srv := server.NewServer(employeeService, departmentService)
	srv.Start()
}
