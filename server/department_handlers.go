package server

import (
	"encoding/json"
	"net/http"
	"strconv"

	"employee-maintenance/models"
	"employee-maintenance/services"
)

func (s *Server) RegisterDepartmentRoutes() {
	s.mux.HandleFunc("GET /departments", s.getDepartments)
	s.mux.HandleFunc("POST /departments", s.createDepartment)
	s.mux.HandleFunc("GET /departments/{id}", s.getDepartment)
	s.mux.HandleFunc("PUT /departments/{id}", s.updateDepartment)
	s.mux.HandleFunc("DELETE /departments/{id}", s.deleteDepartment)
}

func (s *Server) createDepartment(w http.ResponseWriter, r *http.Request) {
	var dept models.Department
	if err := json.NewDecoder(r.Body).Decode(&dept); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	newDept := s.departmentService.Create(dept)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newDept)
}

func (s *Server) getDepartments(w http.ResponseWriter, r *http.Request) {
	departments := s.departmentService.RetrieveAll()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(departments)
}

func (s *Server) getDepartment(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, "Invalid department ID", http.StatusBadRequest)
		return
	}

	dept, err := s.departmentService.Retrieve(id)
	if err != nil {
		if err == services.ErrDepartmentNotFound {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(dept)
}

func (s *Server) updateDepartment(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, "Invalid department ID", http.StatusBadRequest)
		return
	}

	var dept models.Department
	if err := json.NewDecoder(r.Body).Decode(&dept); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if dept.ID != id {
		http.Error(w, "ID in body does not match ID in URL", http.StatusBadRequest)
		return
	}
	updatedDept, err := s.departmentService.Update(dept)
	if err != nil {
		if err == services.ErrDepartmentNotFound {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedDept)
}

func (s *Server) deleteDepartment(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, "Invalid department ID", http.StatusBadRequest)
		return
	}

	err = s.departmentService.Delete(id)
	if err != nil {
		if err == services.ErrDepartmentNotFound {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
