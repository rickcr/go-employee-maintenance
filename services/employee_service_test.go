package services

import (
	"testing"

	"employee-maintenance/models"
)

func TestEmployeeService_Create(t *testing.T) {
	service := NewEmployeeService()
	emp := models.Employee{
		ID:        1,
		FirstName: "John",
		LastName:  "Doe",
		Email:     "john.doe@example.com",
		Department: models.Department{
			ID:   1,
			Name: "Engineering",
		},
	}

	created := service.Create(emp)

	if created.ID != emp.ID || created.FirstName != emp.FirstName {
		t.Errorf("Create() returned %v, want %v", created, emp)
	}
}

func TestEmployeeService_Create_AutoGenerateID(t *testing.T) {
	service := NewEmployeeService()
	dept := models.Department{ID: 1, Name: "Engineering"}

	emp1 := service.Create(models.Employee{FirstName: "John", LastName: "Doe", Email: "john@example.com", Department: dept})
	if emp1.ID != 1 {
		t.Errorf("First auto-generated ID = %v, want 1", emp1.ID)
	}

	emp2 := service.Create(models.Employee{FirstName: "Jane", LastName: "Smith", Email: "jane@example.com", Department: dept})
	if emp2.ID != 2 {
		t.Errorf("Second auto-generated ID = %v, want 2", emp2.ID)
	}

	emp3 := service.Create(models.Employee{FirstName: "Bob", LastName: "Wilson", Email: "bob@example.com", Department: dept})
	if emp3.ID != 3 {
		t.Errorf("Third auto-generated ID = %v, want 3", emp3.ID)
	}
}

func TestEmployeeService_Create_AutoGenerateID_AfterDelete(t *testing.T) {
	service := NewEmployeeService()
	dept := models.Department{ID: 1, Name: "Engineering"}

	service.Create(models.Employee{FirstName: "John", LastName: "Doe", Email: "john@example.com", Department: dept})
	service.Create(models.Employee{FirstName: "Jane", LastName: "Smith", Email: "jane@example.com", Department: dept})
	emp3 := service.Create(models.Employee{FirstName: "Bob", LastName: "Wilson", Email: "bob@example.com", Department: dept})
	if emp3.ID != 3 {
		t.Errorf("Third ID = %v, want 3", emp3.ID)
	}

	service.Delete(2)

	emp4 := service.Create(models.Employee{FirstName: "Alice", LastName: "Brown", Email: "alice@example.com", Department: dept})
	if emp4.ID != 4 {
		t.Errorf("After delete, new ID = %v, want 4 (should use max+1, not fill gaps)", emp4.ID)
	}
}

func TestEmployeeService_Retrieve(t *testing.T) {
	service := NewEmployeeService()
	emp := models.Employee{
		ID:        1,
		FirstName: "John",
		LastName:  "Doe",
		Email:     "john.doe@example.com",
		Department: models.Department{
			ID:   1,
			Name: "Engineering",
		},
	}
	service.Create(emp)

	retrieved, err := service.Retrieve(1)
	if err != nil {
		t.Errorf("Retrieve() error = %v, want nil", err)
	}
	if retrieved.ID != emp.ID || retrieved.Email != emp.Email {
		t.Errorf("Retrieve() = %v, want %v", retrieved, emp)
	}
}

func TestEmployeeService_Retrieve_NotFound(t *testing.T) {
	service := NewEmployeeService()

	_, err := service.Retrieve(999)
	if err != ErrEmployeeNotFound {
		t.Errorf("Retrieve() error = %v, want %v", err, ErrEmployeeNotFound)
	}
}

func TestEmployeeService_RetrieveAll(t *testing.T) {
	service := NewEmployeeService()
	dept := models.Department{ID: 1, Name: "Engineering"}
	emp1 := models.Employee{ID: 1, FirstName: "John", LastName: "Doe", Email: "john@example.com", Department: dept}
	emp2 := models.Employee{ID: 2, FirstName: "Jane", LastName: "Smith", Email: "jane@example.com", Department: dept}
	service.Create(emp1)
	service.Create(emp2)

	all := service.RetrieveAll()
	if len(all) != 2 {
		t.Errorf("RetrieveAll() returned %d items, want 2", len(all))
	}
}

func TestEmployeeService_Update(t *testing.T) {
	service := NewEmployeeService()
	emp := models.Employee{
		ID:        1,
		FirstName: "John",
		LastName:  "Doe",
		Email:     "john.doe@example.com",
		Department: models.Department{
			ID:   1,
			Name: "Engineering",
		},
	}
	service.Create(emp)

	updated := models.Employee{
		ID:        1,
		FirstName: "John",
		LastName:  "Doe",
		Email:     "john.updated@example.com",
		Department: models.Department{
			ID:   2,
			Name: "Marketing",
		},
	}
	result, err := service.Update(updated)
	if err != nil {
		t.Errorf("Update() error = %v, want nil", err)
	}
	if result.Email != "john.updated@example.com" {
		t.Errorf("Update() Email = %v, want john.updated@example.com", result.Email)
	}
	if result.Department.Name != "Marketing" {
		t.Errorf("Update() Department.Name = %v, want Marketing", result.Department.Name)
	}
}

func TestEmployeeService_Update_NotFound(t *testing.T) {
	service := NewEmployeeService()
	emp := models.Employee{ID: 999, FirstName: "Test", LastName: "User", Email: "test@example.com"}

	_, err := service.Update(emp)
	if err != ErrEmployeeNotFound {
		t.Errorf("Update() error = %v, want %v", err, ErrEmployeeNotFound)
	}
}

func TestEmployeeService_Delete(t *testing.T) {
	service := NewEmployeeService()
	emp := models.Employee{
		ID:        1,
		FirstName: "John",
		LastName:  "Doe",
		Email:     "john.doe@example.com",
		Department: models.Department{
			ID:   1,
			Name: "Engineering",
		},
	}
	service.Create(emp)

	err := service.Delete(1)
	if err != nil {
		t.Errorf("Delete() error = %v, want nil", err)
	}

	_, err = service.Retrieve(1)
	if err != ErrEmployeeNotFound {
		t.Errorf("After Delete(), Retrieve() error = %v, want %v", err, ErrEmployeeNotFound)
	}
}

func TestEmployeeService_Delete_NotFound(t *testing.T) {
	service := NewEmployeeService()

	err := service.Delete(999)
	if err != ErrEmployeeNotFound {
		t.Errorf("Delete() error = %v, want %v", err, ErrEmployeeNotFound)
	}
}
