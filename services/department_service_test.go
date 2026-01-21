package services

import (
	"testing"

	"employee-maintenance/models"
)

func TestDepartmentService_Create(t *testing.T) {
	service := NewDepartmentService()
	dept := models.Department{ID: 1, Name: "Engineering"}

	created := service.Create(dept)

	if created.ID != dept.ID || created.Name != dept.Name {
		t.Errorf("Create() returned %v, want %v", created, dept)
	}
}

func TestDepartmentService_Create_AutoGenerateID(t *testing.T) {
	service := NewDepartmentService()

	dept1 := service.Create(models.Department{Name: "Engineering"})
	if dept1.ID != 1 {
		t.Errorf("First auto-generated ID = %v, want 1", dept1.ID)
	}

	dept2 := service.Create(models.Department{Name: "Marketing"})
	if dept2.ID != 2 {
		t.Errorf("Second auto-generated ID = %v, want 2", dept2.ID)
	}

	dept3 := service.Create(models.Department{Name: "Sales"})
	if dept3.ID != 3 {
		t.Errorf("Third auto-generated ID = %v, want 3", dept3.ID)
	}
}

func TestDepartmentService_Create_AutoGenerateID_AfterDelete(t *testing.T) {
	service := NewDepartmentService()

	service.Create(models.Department{Name: "Engineering"})
	service.Create(models.Department{Name: "Marketing"})
	dept3 := service.Create(models.Department{Name: "Sales"})
	if dept3.ID != 3 {
		t.Errorf("Third ID = %v, want 3", dept3.ID)
	}

	service.Delete(2)

	dept4 := service.Create(models.Department{Name: "HR"})
	if dept4.ID != 4 {
		t.Errorf("After delete, new ID = %v, want 4 (should use max+1, not fill gaps)", dept4.ID)
	}
}

func TestDepartmentService_Retrieve(t *testing.T) {
	service := NewDepartmentService()
	dept := models.Department{ID: 1, Name: "Engineering"}
	service.Create(dept)

	retrieved, err := service.Retrieve(1)
	if err != nil {
		t.Errorf("Retrieve() error = %v, want nil", err)
	}
	if retrieved.ID != dept.ID || retrieved.Name != dept.Name {
		t.Errorf("Retrieve() = %v, want %v", retrieved, dept)
	}
}

func TestDepartmentService_Retrieve_NotFound(t *testing.T) {
	service := NewDepartmentService()

	_, err := service.Retrieve(999)
	if err != ErrDepartmentNotFound {
		t.Errorf("Retrieve() error = %v, want %v", err, ErrDepartmentNotFound)
	}
}

func TestDepartmentService_RetrieveAll(t *testing.T) {
	service := NewDepartmentService()
	dept1 := models.Department{ID: 1, Name: "Engineering"}
	dept2 := models.Department{ID: 2, Name: "Marketing"}
	service.Create(dept1)
	service.Create(dept2)

	all := service.RetrieveAll()
	if len(all) != 2 {
		t.Errorf("RetrieveAll() returned %d items, want 2", len(all))
	}
}

func TestDepartmentService_Update(t *testing.T) {
	service := NewDepartmentService()
	dept := models.Department{ID: 1, Name: "Engineering"}
	service.Create(dept)

	updated := models.Department{ID: 1, Name: "Software Engineering"}
	result, err := service.Update(updated)
	if err != nil {
		t.Errorf("Update() error = %v, want nil", err)
	}
	if result.Name != "Software Engineering" {
		t.Errorf("Update() Name = %v, want Software Engineering", result.Name)
	}
}

func TestDepartmentService_Update_NotFound(t *testing.T) {
	service := NewDepartmentService()
	dept := models.Department{ID: 999, Name: "Test"}

	_, err := service.Update(dept)
	if err != ErrDepartmentNotFound {
		t.Errorf("Update() error = %v, want %v", err, ErrDepartmentNotFound)
	}
}

func TestDepartmentService_Delete(t *testing.T) {
	service := NewDepartmentService()
	dept := models.Department{ID: 1, Name: "Engineering"}
	service.Create(dept)

	err := service.Delete(1)
	if err != nil {
		t.Errorf("Delete() error = %v, want nil", err)
	}

	_, err = service.Retrieve(1)
	if err != ErrDepartmentNotFound {
		t.Errorf("After Delete(), Retrieve() error = %v, want %v", err, ErrDepartmentNotFound)
	}
}

func TestDepartmentService_Delete_NotFound(t *testing.T) {
	service := NewDepartmentService()

	err := service.Delete(999)
	if err != ErrDepartmentNotFound {
		t.Errorf("Delete() error = %v, want %v", err, ErrDepartmentNotFound)
	}
}
