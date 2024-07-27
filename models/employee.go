package models

import "time"

type Employee struct {
	ID         int
	FirstName  string
	LastName   string
	Email      string
	Phone      string
	Position   string
	Department string
	Salary     float64
	HireDate   time.Time
	IsActive   bool
}
