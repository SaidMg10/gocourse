package basics

import "fmt"

type Employee struct {
	FirstName string
	LastName  string
	Age       int
}

func functionma() {
	// PascalCase
	// Structs, interfaces, enums

	// snake_case
	// Commonly use for naming variables, constants and filenames
	// Eg. user_id, first_name, http_request

	// UPPERCASE
	// Is used exclusively for naming constants in go
	// Use case is constants

	// Preferly use this for variables than use snake_case
	// Eg. javaScript, htmlDocument, isValid

	// Package names should be in lowercase

	const MAXRETRIES = 5

	employeeID := 1001
	fmt.Println("EmployeeID: ", employeeID)
}
