package db

const (
	CreateEmployeesTable = `CREATE TABLE IF NOT EXISTS employees (
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
    );`

	DropEmployeesTable = `DROP TABLE IF EXISTS employees;`
)
