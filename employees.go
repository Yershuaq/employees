package employees

import "fmt"

type Employee interface {
	GetDetails() string
}

type FullTimeEmployee struct {
	ID     uint64
	Name   string
	Salary uint32
}

func (fte FullTimeEmployee) GetDetails() string {
	return fmt.Sprintf("ID: %d, Name: %s, Salary: %d", fte.ID, fte.Name, fte.Salary)
}

type Company struct {
	Employees map[string]Employee
}

func NewCompany() *Company {
	return &Company{Employees: make(map[string]Employee)}
}

func (c *Company) AddEmployee(emp Employee) {
	key := fmt.Sprintf("%d", len(c.Employees)+1)
	c.Employees[key] = emp
}

func (c *Company) ListEmployees() {
	for _, emp := range c.Employees {
		fmt.Println(emp.GetDetails())
	}
}
