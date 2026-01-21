package services

import (
	"errors"
	"sync"

	"employee-maintenance/models"
)

var (
	ErrDepartmentNotFound = errors.New("department not found")
)

type DepartmentService struct {
	mu          sync.RWMutex
	departments map[int]models.Department
}

func NewDepartmentService() *DepartmentService {
	return &DepartmentService{
		departments: make(map[int]models.Department),
	}
}

func (s *DepartmentService) Create(dept models.Department) models.Department {
	s.mu.Lock()
	defer s.mu.Unlock()
	if dept.ID == 0 {
		dept.ID = s.nextID()
	}
	s.departments[dept.ID] = dept
	return dept
}

func (s *DepartmentService) nextID() int {
	maxID := 0
	for _, d := range s.departments {
		if d.ID > maxID {
			maxID = d.ID
		}
	}
	return maxID + 1
}

func (s *DepartmentService) Retrieve(id int) (models.Department, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	dept, exists := s.departments[id]
	if !exists {
		return models.Department{}, ErrDepartmentNotFound
	}
	return dept, nil
}

func (s *DepartmentService) RetrieveAll() []models.Department {
	s.mu.RLock()
	defer s.mu.RUnlock()
	result := make([]models.Department, 0, len(s.departments))
	for _, dept := range s.departments {
		result = append(result, dept)
	}
	return result
}

func (s *DepartmentService) Update(dept models.Department) (models.Department, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, exists := s.departments[dept.ID]; !exists {
		return models.Department{}, ErrDepartmentNotFound
	}
	s.departments[dept.ID] = dept
	return dept, nil
}

func (s *DepartmentService) Delete(id int) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, exists := s.departments[id]; !exists {
		return ErrDepartmentNotFound
	}
	delete(s.departments, id)
	return nil
}
