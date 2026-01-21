package client

// Not using right now, but could be useful in tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"employee-maintenance/models"
)

type EmployeeClient struct {
	baseURL    string
	httpClient *http.Client
}

func NewEmployeeClient(baseURL string) *EmployeeClient {
	return &EmployeeClient{
		baseURL:    baseURL,
		httpClient: &http.Client{},
	}
}

func NewEmployeeClientWithHTTPClient(baseURL string, httpClient *http.Client) *EmployeeClient {
	return &EmployeeClient{
		baseURL:    baseURL,
		httpClient: httpClient,
	}
}

func (c *EmployeeClient) Create(emp models.Employee) (models.Employee, error) {
	body, err := json.Marshal(emp)
	if err != nil {
		return models.Employee{}, fmt.Errorf("failed to marshal employee: %w", err)
	}

	resp, err := c.httpClient.Post(
		c.baseURL+"/employees",
		"application/json",
		bytes.NewReader(body),
	)
	if err != nil {
		return models.Employee{}, fmt.Errorf("failed to create employee: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated && resp.StatusCode != http.StatusOK {
		return models.Employee{}, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var created models.Employee
	if err := json.NewDecoder(resp.Body).Decode(&created); err != nil {
		return models.Employee{}, fmt.Errorf("failed to decode response: %w", err)
	}

	return created, nil
}

func (c *EmployeeClient) Retrieve(id int) (models.Employee, error) {
	resp, err := c.httpClient.Get(fmt.Sprintf("%s/employees/%d", c.baseURL, id))
	if err != nil {
		return models.Employee{}, fmt.Errorf("failed to retrieve employee: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return models.Employee{}, fmt.Errorf("employee not found: %d", id)
	}

	if resp.StatusCode != http.StatusOK {
		return models.Employee{}, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var emp models.Employee
	if err := json.NewDecoder(resp.Body).Decode(&emp); err != nil {
		return models.Employee{}, fmt.Errorf("failed to decode response: %w", err)
	}

	return emp, nil
}

func (c *EmployeeClient) RetrieveAll() ([]models.Employee, error) {
	resp, err := c.httpClient.Get(c.baseURL + "/employees")
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve employees: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var employees []models.Employee
	if err := json.NewDecoder(resp.Body).Decode(&employees); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return employees, nil
}

func (c *EmployeeClient) Update(emp models.Employee) (models.Employee, error) {
	body, err := json.Marshal(emp)
	if err != nil {
		return models.Employee{}, fmt.Errorf("failed to marshal employee: %w", err)
	}

	req, err := http.NewRequest(
		http.MethodPut,
		fmt.Sprintf("%s/employees/%d", c.baseURL, emp.ID),
		bytes.NewReader(body),
	)
	if err != nil {
		return models.Employee{}, fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return models.Employee{}, fmt.Errorf("failed to update employee: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return models.Employee{}, fmt.Errorf("employee not found: %d", emp.ID)
	}

	if resp.StatusCode != http.StatusOK {
		return models.Employee{}, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var updated models.Employee
	if err := json.NewDecoder(resp.Body).Decode(&updated); err != nil {
		return models.Employee{}, fmt.Errorf("failed to decode response: %w", err)
	}

	return updated, nil
}

func (c *EmployeeClient) Delete(id int) error {
	req, err := http.NewRequest(
		http.MethodDelete,
		fmt.Sprintf("%s/employees/%d", c.baseURL, id),
		nil,
	)
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("failed to delete employee: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return fmt.Errorf("employee not found: %d", id)
	}

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusNoContent {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode, string(body))
	}

	return nil
}
