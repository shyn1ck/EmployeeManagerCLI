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

	for {
		choice := DisplayMenuAndGetChoice()
		HandleChoice(choice)
	}
}

func DisplayMenuAndGetChoice() int {
	fmt.Println("\n\033[31;1m╔══════════════════════════════════╗\033[0m")
	fmt.Println("\033[31;1m║*** Employee Management System ***║\033[0m")
	fmt.Println("\033[31;1m╚══════════════════════════════════╝\033[0m")

	fmt.Println("\033[32;1m1. Add Employee\033[0m")
	fmt.Println("\033[33;1m2. Get Employee by ID\033[0m")
	fmt.Println("\033[34;1m3. Update Employee\033[0m")
	fmt.Println("\033[35;1m4. Delete Employee\033[0m")
	fmt.Println("\033[36;1m5. Show All Employees\033[0m")
	fmt.Println("\033[31;1m6. Create Employees Table\033[0m")
	fmt.Println("\033[32;1m7. Insert Initial Data\033[0m")
	fmt.Println("\033[33;1m8. Drop Employees Table\033[0m")
	fmt.Println("\033[34;1m9. Exit\033[0m")

	var choice int
	fmt.Print("\n\033[32;1mEnter your choice: \033[0m")
	_, err := fmt.Scan(&choice)
	if err != nil {
		fmt.Println("Invalid input. Please try again.")
		return DisplayMenuAndGetChoice()
	}

	return choice
}

func HandleChoice(choice int) {
	switch choice {
	case 1:
		AddEmployee()
	case 2:
		GetEmployeeByID()
	case 3:
		UpdateEmployee()
	case 4:
		DeleteEmployee()
	case 5:
		listAllEmployees()
	case 6:
		CreateTable()
	case 7:
		InsertInitialData()
	case 8:
		DropTable()
	case 9:
		fmt.Println("Exiting program.")
		os.Exit(0)
	default:
		fmt.Println("Invalid choice. Try again.")
	}
}

func CreateTable() {
	err := db.CreateEmployeesTable()
	if err != nil {
		fmt.Println("Error creating table. Error:", err)
	} else {
		fmt.Println("Table 'employees' created successfully.")
	}
}

func InsertInitialData() {
	err := db.InsertInitialData()
	if err != nil {
		fmt.Println("Error inserting initial data. Error:", err)
	} else {
		fmt.Println("Initial data inserted successfully.")
	}
}

func AddEmployee() {
	var emp models.Employee
	fmt.Println("Enter employee details:")

	fmt.Print("First Name: ")
	fmt.Scan(&emp.FirstName)
	fmt.Print("Last Name: ")
	fmt.Scan(&emp.LastName)
	fmt.Print("Email: ")
	fmt.Scan(&emp.Email)
	fmt.Print("Phone: ")
	fmt.Scan(&emp.Phone)
	fmt.Print("Position: ")
	fmt.Scan(&emp.Position)
	fmt.Print("Department: ")
	fmt.Scan(&emp.Department)
	fmt.Print("Salary: ")
	fmt.Scan(&emp.Salary)
	fmt.Print("Hire Date (YYYY-MM-DD): ")
	var hireDate string
	fmt.Scan(&hireDate)
	emp.HireDate, _ = time.Parse("2006-01-02", hireDate)
	fmt.Print("Active (true/false): ")
	fmt.Scan(&emp.IsActive)

	err := repository.CreateEmployee(emp)
	if err != nil {
		fmt.Println("Error adding employee. Error:", err)
	} else {
		fmt.Println("Employee added successfully.")
	}
}

func GetEmployeeByID() {
	fmt.Print("Enter employee ID: ")
	var id int
	fmt.Scan(&id)

	employee, err := repository.GetEmployeeByID(id)
	if err != nil {
		fmt.Println("Error getting employee. Error:", err)
	} else if (employee == models.Employee{}) {
		fmt.Println("Employee with this ID not found.")
	} else {
		fmt.Printf("ID: %d\nFirst Name: %s\nLast Name: %s\nEmail: %s\nPhone: %s\nPosition: %s\nDepartment: %s\nSalary: %.2f\nHire Date: %s\nActive: %t\n",
			employee.ID, employee.FirstName, employee.LastName, employee.Email, employee.Phone, employee.Position, employee.Department, employee.Salary, employee.HireDate.Format("2006-01-02"), employee.IsActive)
	}
}

func UpdateEmployee() {
	var emp models.Employee
	fmt.Print("Enter employee ID to update: ")
	fmt.Scan(&emp.ID)

	fmt.Println("Enter new employee details:")

	fmt.Print("First Name: ")
	fmt.Scan(&emp.FirstName)
	fmt.Print("Last Name: ")
	fmt.Scan(&emp.LastName)
	fmt.Print("Email: ")
	fmt.Scan(&emp.Email)
	fmt.Print("Phone: ")
	fmt.Scan(&emp.Phone)
	fmt.Print("Position: ")
	fmt.Scan(&emp.Position)
	fmt.Print("Department: ")
	fmt.Scan(&emp.Department)
	fmt.Print("Salary: ")
	fmt.Scan(&emp.Salary)
	fmt.Print("Hire Date (YYYY-MM-DD): ")
	var hireDate string
	fmt.Scan(&hireDate)
	emp.HireDate, _ = time.Parse("2006-01-02", hireDate)
	fmt.Print("Active (true/false): ")
	fmt.Scan(&emp.IsActive)

	err := repository.UpdateEmployeeByID(emp)
	if err != nil {
		fmt.Println("Error updating employee. Error:", err)
	} else {
		fmt.Println("Employee updated successfully.")
	}
}

func DeleteEmployee() {
	fmt.Print("Enter employee ID to delete: ")
	var id int
	fmt.Scan(&id)

	err := repository.DeleteEmployeeByID(id)
	if err != nil {
		fmt.Println("Error deleting employee. Error:", err)
	} else {
		fmt.Println("Employee deleted successfully.")
	}
}

func DropTable() {
	err := repository.DropEmployeesTable()
	if err != nil {
		fmt.Println("Error dropping table. Error:", err)
	} else {
		fmt.Println("Table 'employees' has been dropped.")
	}
}
func listAllEmployees() {
	employees, err := repository.GetAllEmployees()
	if err != nil {
		fmt.Println("Error getting list of employees. Error:", err)
	} else if len(employees) == 0 {
		fmt.Println("Employee list is empty.")
	} else {
		for _, emp := range employees {
			fmt.Printf("ID: %d\nFirst Name: %s\nLast Name: %s\nEmail: %s\nPhone: %s\nPosition: %s\nDepartment: %s\nSalary: %.2f\nHire Date: %s\nActive: %t\n\n",
				emp.ID, emp.FirstName, emp.LastName, emp.Email, emp.Phone, emp.Position, emp.Department, emp.Salary, emp.HireDate.Format("2006-01-02"), emp.IsActive)
		}
	}
}
