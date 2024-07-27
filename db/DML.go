package db

const (
	InsertEmployee = `INSERT INTO employees 
		(first_name, last_name, email, phone, position, department, salary, hire_date, is_active) 
	VALUES 
		($1, $2, $3, $4, $5, $6, $7, $8, $9)
	RETURNING id;`

	SelectEmployeeByID = `SELECT 
		id, first_name, last_name, email, phone, position, department, salary, hire_date, is_active 
	FROM employees 
	WHERE id = $1;`

	SelectAllEmployees = `SELECT 
		id, first_name, last_name, email, phone, position, department, salary, hire_date, is_active 
	FROM employees;`

	UpdateEmployee = `UPDATE employees 
	SET first_name = $1, last_name = $2, email = $3, phone = $4, position = $5, department = $6, salary = $7, hire_date = $8, is_active = $9 
	WHERE id = $10;`

	DeleteEmployee = `DELETE FROM employees 
	WHERE id = $1;`

	DropEmployeesTable = `DROP TABLE IF EXISTS employees;`
)
