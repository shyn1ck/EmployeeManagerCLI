package main

import (
	"EmployeeManagerCLI/db"
	"EmployeeManagerCLI/models"
	"EmployeeManagerCLI/repository"
	"fmt"
	"os"
	"time"
)

func main() {
	err := db.ConnectToDB()
	if err != nil {
		fmt.Println("Error connecting to database. Error:", err)
		os.Exit(1)
	}

	checkIfTableExists()

	for {
		choice := displayMenuAndGetChoice()
		handleChoice(choice)
	}
}

func displayMenuAndGetChoice() int {
	fmt.Println("Employee Manager CLI")
	fmt.Println("1. Create Employee")
	fmt.Println("2. Get All Employees")
	fmt.Println("3. Get Employee by ID")
	fmt.Println("4. Update Employee")
	fmt.Println("5. Delete Employee")
	fmt.Println("6. Exit")

	var choice int
	fmt.Print("Enter your choice: ")
	_, err := fmt.Scan(&choice)
	if err != nil {
		fmt.Println("Invalid choice. Please enter a number between 1 and 6.")
		return 0
	}
	return choice
}

func handleChoice(choice int) {
	switch choice {
	case 1:
		createEmployee()
	case 2:
		getAllEmployees()
	case 3:
		getEmployeeByID()
	case 4:
		updateEmployee()
	case 5:
		deleteEmployee()
	case 6:
		fmt.Println("Exiting...")
		os.Exit(0)
	default:
		fmt.Println("Invalid choice. Please try again.")
	}
}

func checkIfTableExists() {
	rows, err := db.GetDBConn().Query("SELECT to_regclass('public.employees');")
	if err != nil {
		fmt.Println("Error during checking table existence. Error:", err)
		return
	}
	defer rows.Close()

	var tableName *string
	if rows.Next() {
		err := rows.Scan(&tableName)
		if err != nil {
			fmt.Println("Error during scanning. Error:", err)
			return
		}
	}

	if tableName == nil {
		fmt.Println("Table 'employees' does not exist.")
	} else {
		fmt.Println("Table 'employees' exists.")
	}
}

func createEmployee() {
	e := models.Employee{}

	fmt.Print("Enter first name: ")
	_, err := fmt.Scan(&e.FirstName)
	if err != nil {
		fmt.Println("Error reading input. Error:", err)
		return
	}

	fmt.Print("Enter last name: ")
	_, err = fmt.Scan(&e.LastName)
	if err != nil {
		fmt.Println("Error reading input. Error:", err)
		return
	}

	fmt.Print("Enter email: ")
	_, err = fmt.Scan(&e.Email)
	if err != nil {
		fmt.Println("Error reading input. Error:", err)
		return
	}

	fmt.Print("Enter phone: ")
	_, err = fmt.Scan(&e.Phone)
	if err != nil {
		fmt.Println("Error reading input. Error:", err)
		return
	}

	fmt.Print("Enter position: ")
	_, err = fmt.Scan(&e.Position)
	if err != nil {
		fmt.Println("Error reading input. Error:", err)
		return
	}

	fmt.Print("Enter department: ")
	_, err = fmt.Scan(&e.Department)
	if err != nil {
		fmt.Println("Error reading input. Error:", err)
		return
	}

	fmt.Print("Enter salary: ")
	_, err = fmt.Scan(&e.Salary)
	if err != nil {
		fmt.Println("Invalid salary input. Error:", err)
		return
	}

	e.HireDate = time.Now()
	e.IsActive = true

	err = repository.CreateEmployee(e)
	if err != nil {
		fmt.Println("Error during insertion. Error:", err)
		return
	}
	fmt.Println("Employee created successfully.")
}

func getAllEmployees() {
	employees, err := repository.GetAllEmployees()
	if err != nil {
		fmt.Println("Error during query. Error:", err)
		return
	}
	if len(employees) == 0 {
		fmt.Println("No employees found.")
		return
	}
	for _, e := range employees {
		fmt.Printf("%+v\n", e)
	}
}

func getEmployeeByID() {
	var id int
	fmt.Print("Enter employee ID: ")
	_, err := fmt.Scan(&id)
	if err != nil {
		fmt.Println("Invalid ID input. Error:", err)
		return
	}

	e, err := repository.GetEmployeeByID(id)
	if err != nil {
		fmt.Println("Error during query. Error:", err)
		return
	}
	fmt.Printf("Employee: %+v\n", e)
}

func updateEmployee() {
	e := models.Employee{}

	fmt.Print("Enter employee ID to update: ")
	_, err := fmt.Scan(&e.ID)
	if err != nil {
		fmt.Println("Invalid ID input. Error:", err)
		return
	}

	fmt.Print("Enter new first name: ")
	_, err = fmt.Scan(&e.FirstName)
	if err != nil {
		fmt.Println("Error reading input. Error:", err)
		return
	}

	fmt.Print("Enter new last name: ")
	_, err = fmt.Scan(&e.LastName)
	if err != nil {
		fmt.Println("Error reading input. Error:", err)
		return
	}

	fmt.Print("Enter new email: ")
	_, err = fmt.Scan(&e.Email)
	if err != nil {
		fmt.Println("Error reading input. Error:", err)
		return
	}

	fmt.Print("Enter new phone: ")
	_, err = fmt.Scan(&e.Phone)
	if err != nil {
		fmt.Println("Error reading input. Error:", err)
		return
	}

	fmt.Print("Enter new position: ")
	_, err = fmt.Scan(&e.Position)
	if err != nil {
		fmt.Println("Error reading input. Error:", err)
		return
	}

	fmt.Print("Enter new department: ")
	_, err = fmt.Scan(&e.Department)
	if err != nil {
		fmt.Println("Error reading input. Error:", err)
		return
	}

	fmt.Print("Enter new salary: ")
	_, err = fmt.Scan(&e.Salary)
	if err != nil {
		fmt.Println("Invalid salary input. Error:", err)
		return
	}

	e.HireDate = time.Now()
	e.IsActive = true

	err = repository.UpdateEmployeeByID(e)
	if err != nil {
		fmt.Println("Error during update. Error:", err)
		return
	}
	fmt.Println("Employee updated successfully.")
}

func deleteEmployee() {
	var id int
	fmt.Print("Enter employee ID to delete: ")
	_, err := fmt.Scan(&id)
	if err != nil {
		fmt.Println("Invalid ID input. Error:", err)
		return
	}

	err = repository.DeleteEmployeeByID(id)
	if err != nil {
		fmt.Println("Error during delete. Error:", err)
		return
	}
	fmt.Println("Employee deleted successfully.")
}
