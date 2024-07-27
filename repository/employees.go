package repository

import (
	"EmployeeManagerCLI/db"
	"EmployeeManagerCLI/models"
	"fmt"
)

func CreateEmployee(e models.Employee) error {
	_, err := db.GetDBConn().Exec(db.InsertEmployee,
		e.FirstName, e.LastName, e.Email, e.Phone, e.Position, e.Department, e.Salary, e.HireDate, e.IsActive)
	if err != nil {
		fmt.Println("Error inserting data: ", err.Error())
		return err
	}
	return nil
}

func GetAllEmployees() ([]models.Employee, error) {
	rows, err := db.GetDBConn().Query(db.SelectAllEmployees)
	if err != nil {
		fmt.Println("Error executing query: ", err.Error())
		return nil, err
	}
	defer rows.Close()

	var employees []models.Employee
	for rows.Next() {
		var e models.Employee
		err = rows.Scan(&e.ID, &e.FirstName, &e.LastName, &e.Email, &e.Phone, &e.Position, &e.Department, &e.Salary, &e.HireDate, &e.IsActive)
		if err != nil {
			fmt.Println("Error reading data: ", err.Error())
			continue
		}
		employees = append(employees, e)
	}

	return employees, nil
}

func GetEmployeeByID(id int) (models.Employee, error) {
	var e models.Employee
	row := db.GetDBConn().QueryRow(db.SelectEmployeeByID, id)
	err := row.Scan(&e.ID, &e.FirstName, &e.LastName, &e.Email, &e.Phone, &e.Position, &e.Department, &e.Salary, &e.HireDate, &e.IsActive)
	if err != nil {
		fmt.Println("Error executing query: ", err.Error())
		return models.Employee{}, err
	}

	return e, nil
}

func UpdateEmployeeByID(e models.Employee) error {
	_, err := db.GetDBConn().Exec(db.UpdateEmployee,
		e.FirstName, e.LastName, e.Email, e.Phone, e.Position, e.Department, e.Salary, e.HireDate, e.IsActive, e.ID)
	if err != nil {
		fmt.Println("Error updating data: ", err.Error())
		return err
	}

	return nil
}

func DeleteEmployeeByID(id int) error {
	_, err := db.GetDBConn().Exec(db.DeleteEmployee, id)
	if err != nil {
		fmt.Println("Error deleting data: ", err.Error())
		return err
	}

	return nil
}

func DropEmployeesTable() error {
	_, err := db.GetDBConn().Exec(db.DropEmployeesTable)
	if err != nil {
		fmt.Println("Error dropping table: ", err.Error())
		return err
	}
	fmt.Println("Employees table dropped successfully.")
	return nil
}
