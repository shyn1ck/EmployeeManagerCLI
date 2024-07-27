package db

import (
	_ "database/sql"
	"fmt"
)

func CreateEmployeesTable() error {
	query := `
	CREATE TABLE IF NOT EXISTS employees (
		id SERIAL PRIMARY KEY,
		first_name VARCHAR(100) NOT NULL,
		last_name VARCHAR(100) NOT NULL,
		email VARCHAR(100) UNIQUE NOT NULL,
		phone VARCHAR(20),
		position VARCHAR(100),
		department VARCHAR(100),
		salary NUMERIC(10, 2),
		hire_date DATE,
		is_active BOOLEAN DEFAULT TRUE
	);
	`
	_, err := dbConn.Exec(query)
	if err != nil {
		return err
	}
	fmt.Println("Таблица 'employees' успешно создана.")
	return nil
}
func InsertInitialData() error {
	query := `
	INSERT INTO employees (first_name, last_name, email, phone, position, department, salary, hire_date, is_active)
	VALUES
	('John', 'Doe', 'john.doe@example.com', '+1234567890', 'Software Engineer', 'Engineering', 90000.00, '2024-01-15', TRUE),
	('Jane', 'Smith', 'jane.smith@example.com', '+1234567891', 'Product Manager', 'Marketing', 95000.00, '2023-11-10', TRUE),
	('Alice', 'Johnson', 'alice.johnson@example.com', '+1234567892', 'UX Designer', 'Design', 85000.00, '2024-02-01', TRUE),
	('Bob', 'Brown', 'bob.brown@example.com', '+1234567893', 'Data Scientist', 'Analytics', 100000.00, '2024-03-20', TRUE),
	('Charlie', 'Williams', 'charlie.williams@example.com', '+1234567894', 'DevOps Engineer', 'Engineering', 92000.00, '2023-12-05', TRUE),
	('David', 'Jones', 'david.jones@example.com', '+1234567895', 'Marketing Specialist', 'Marketing', 80000.00, '2024-04-10', TRUE),
	('Emily', 'Miller', 'emily.miller@example.com', '+1234567896', 'HR Manager', 'Human Resources', 87000.00, '2024-05-15', TRUE),
	('Frank', 'Davis', 'frank.davis@example.com', '+1234567897', 'Front-End Developer', 'Engineering', 88000.00, '2024-06-30', TRUE),
	('Grace', 'Wilson', 'grace.wilson@example.com', '+1234567898', 'Back-End Developer', 'Engineering', 91000.00, '2023-09-25', TRUE),
	('Henry', 'Moore', 'henry.moore@example.com', '+1234567899', 'Content Strategist', 'Marketing', 82000.00, '2024-07-20', TRUE),
	('Isabella', 'Taylor', 'isabella.taylor@example.com', '+1234567800', 'Customer Support Specialist', 'Support', 78000.00, '2024-08-15', TRUE),
	('Jack', 'Anderson', 'jack.anderson@example.com', '+1234567801', 'IT Support Specialist', 'IT', 80000.00, '2023-10-10', TRUE),
	('Karen', 'Thomas', 'karen.thomas@example.com', '+1234567802', 'Business Analyst', 'Analytics', 89000.00, '2024-09-05', TRUE),
	('Leo', 'Jackson', 'leo.jackson@example.com', '+1234567803', 'SEO Specialist', 'Marketing', 83000.00, '2024-10-01', TRUE),
	('Mia', 'White', 'mia.white@example.com', '+1234567804', 'Sales Executive', 'Sales', 91000.00, '2024-11-10', TRUE),
	('Nathan', 'Harris', 'nathan.harris@example.com', '+1234567805', 'Software Tester', 'Engineering', 86000.00, '2023-12-20', TRUE),
	('Olivia', 'Martin', 'olivia.martin@example.com', '+1234567806', 'Graphic Designer', 'Design', 85000.00, '2024-01-10', TRUE),
	('Paul', 'Lee', 'paul.lee@example.com', '+1234567807', 'Project Manager', 'Engineering', 95000.00, '2024-02-15', TRUE),
	('Quinn', 'Walker', 'quinn.walker@example.com', '+1234567808', 'Legal Advisor', 'Legal', 93000.00, '2024-03-25', TRUE),
	('Rachel', 'Hall', 'rachel.hall@example.com', '+1234567809', 'Accountant', 'Finance', 87000.00, '2024-04-30', TRUE),
	('Steve', 'Allen', 'steve.allen@example.com', '+1234567810', 'Product Designer', 'Design', 88000.00, '2024-05-20', TRUE);
	`
	_, err := dbConn.Exec(query)
	if err != nil {
		return err
	}
	fmt.Println("The initial data has been successfully inserted.")
	return nil
}
