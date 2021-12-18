package main

import "time"

type Employees struct {
	EmpNo     int8      `json:"emp_no,omitempty"`
	BirthDate time.Time `json:"birth_date"`
	FirstName string    `json:"first_name,omitempty"`
	LastName  string    `json:"last_name,omitempty"`
	Gender    string    `json:"gender,omitempty"`
	HireDate  time.Time `json:"hire_date"`
}
