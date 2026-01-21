package services

import (
	"errors"
	"sync"

	"employee-maintenance/models"
)

var (
	ErrEmployeeNotFound = errors.New("employee not found")
)

type EmployeeService struct {
	mu        sync.RWMutex
	employees map[int]models.Employee
}

func NewEmployeeService() *EmployeeService {
	return &EmployeeService{
		employees: make(map[int]models.Employee),
	}
}

func (s *EmployeeService) Create(emp models.Employee) models.Employee {
	s.mu.Lock()
	defer s.mu.Unlock()
	if emp.ID == 0 {
		emp.ID = s.nextID()
	}
	s.employees[emp.ID] = emp
	return emp
}

func (s *EmployeeService) nextID() int {
	maxID := 0
	for _, e := range s.employees {
		if e.ID > maxID {
			maxID = e.ID
		}
	}
	return maxID + 1
}

func (s *EmployeeService) Retrieve(id int) (models.Employee, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	emp, exists := s.employees[id]
	if !exists {
		return models.Employee{}, ErrEmployeeNotFound
	}
	return emp, nil
}

func (s *EmployeeService) RetrieveAll() []models.Employee {
	s.mu.RLock()
	defer s.mu.RUnlock()
	result := make([]models.Employee, 0, len(s.employees))
	for _, emp := range s.employees {
		result = append(result, emp)
	}
	return result
}

func (s *EmployeeService) Update(emp models.Employee) (models.Employee, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, exists := s.employees[emp.ID]; !exists {
		return models.Employee{}, ErrEmployeeNotFound
	}
	s.employees[emp.ID] = emp
	return emp, nil
}

func (s *EmployeeService) Delete(id int) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, exists := s.employees[id]; !exists {
		return ErrEmployeeNotFound
	}
	delete(s.employees, id)
	return nil
}
