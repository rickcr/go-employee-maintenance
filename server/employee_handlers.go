package server

import (
	"encoding/json"
	"net/http"
	"strconv"

	"employee-maintenance/models"
	"employee-maintenance/services"
)

func (s *Server) RegisterEmployeeRoutes() {
	s.mux.HandleFunc("GET /employees", s.getEmployees)
	s.mux.HandleFunc("POST /employees", s.createEmployee)
	s.mux.HandleFunc("GET /employees/{id}", s.getEmployee)
	s.mux.HandleFunc("PUT /employees/{id}", s.updateEmployee)
	s.mux.HandleFunc("DELETE /employees/{id}", s.deleteEmployee)
}

func (s *Server) createEmployee(w http.ResponseWriter, r *http.Request) {
	var emp models.Employee
	if err := json.NewDecoder(r.Body).Decode(&emp); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	newEmp := s.employeeService.Create(emp)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newEmp)
}

func (s *Server) getEmployees(w http.ResponseWriter, r *http.Request) {
	employees := s.employeeService.RetrieveAll()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(employees)
}

func (s *Server) getEmployee(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, "Invalid employee ID", http.StatusBadRequest)
		return
	}

	emp, err := s.employeeService.Retrieve(id)
	if err != nil {
		if err == services.ErrEmployeeNotFound {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(emp)
}

func (s *Server) updateEmployee(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, "Invalid employee ID", http.StatusBadRequest)
		return
	}

	var emp models.Employee
	if err := json.NewDecoder(r.Body).Decode(&emp); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if emp.ID != id {
		http.Error(w, "ID in body does not match ID in URL", http.StatusBadRequest)
		return
	}
	updatedEmp, err := s.employeeService.Update(emp)
	if err != nil {
		if err == services.ErrEmployeeNotFound {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedEmp)
}

func (s *Server) deleteEmployee(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, "Invalid employee ID", http.StatusBadRequest)
		return
	}

	err = s.employeeService.Delete(id)
	if err != nil {
		if err == services.ErrEmployeeNotFound {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
