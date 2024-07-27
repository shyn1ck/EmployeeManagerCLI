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
		fmt.Println("Error during inserting value. Error: ", err.Error())
		return err
	}

	return nil
}

func GetAllEmployees() ([]models.Employee, error) {
	rows, err := db.GetDBConn().Query(db.SelectAllEmployees)
	if err != nil {
		fmt.Println("Error during query. Error: ", err.Error())
		return nil, err
	}
	defer rows.Close()

	var employees []models.Employee
	for rows.Next() {
		var e models.Employee
		err = rows.Scan(&e.ID, &e.FirstName, &e.LastName, &e.Email, &e.Phone, &e.Position, &e.Department, &e.Salary, &e.HireDate, &e.IsActive)
		if err != nil {
			fmt.Println("Error during query. Error: ", err.Error())
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
		fmt.Println("Error during query. Error: ", err.Error())
		return models.Employee{}, err
	}

	return e, nil
}

func UpdateEmployeeByID(e models.Employee) error {
	_, err := db.GetDBConn().Exec(db.UpdateEmployee,
		e.FirstName, e.LastName, e.Email, e.Phone, e.Position, e.Department, e.Salary, e.HireDate, e.IsActive, e.ID)
	if err != nil {
		fmt.Println("Error during query. Error: ", err.Error())
		return err
	}

	return nil
}

func DeleteEmployeeByID(id int) error {
	_, err := db.GetDBConn().Exec(db.DeleteEmployee, id)
	if err != nil {
		fmt.Println("Error during query. Error: ", err.Error())
		return err
	}

	return nil
}
