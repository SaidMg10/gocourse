package basics

import "fmt"

func main() {
	// func <name>(parameters list) returnType {
	// code block
	// return value
	// }

	// For Public functions  use UperCase
	// For Private functions use lowercase

	// calc := addConsole()
	// sum := add(10, 20)

	// fmt.Println(calc)
	// fmt.Println(sum)
	// fmt.Println(add(2, 3))

	//greet := func() {
	//	fmt.Println("Hello Anonymous Function")
	//}

	// greet()

	// operation := add

	// result := operation(3, 5)

	// fmt.Println(result)

	// Pasing a function as an argument
	result := applyOpeeration(5, 3, add)
	fmt.Println("5 + 3 =", result)

	// Returning and using a function
	multiplyBy2 := createMultiplier(2)
	fmt.Println("6 * 2 =", multiplyBy2(6))
}

func add(a, b int) int {
	return a + b
}

// Function that takes a function as an argument
func applyOpeeration(x int, y int, operation func(int, int) int) int {
	return operation(x, y)
}

// Function that returns a function
func createMultiplier(factor int) func(int) int {
	return func(x int) int {
		return x * factor
	}
}

// func addConsole() int {
//	fmt.Println("Coloque un numero")
//	var a, b int
//	fmt.Println("Ingrese el valor de a")
//	fmt.Scanln(&a)
//	fmt.Println("Ingrese el valor de b")
//	fmt.Scanln(&b)
//
//	return a + b
//}
