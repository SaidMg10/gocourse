package basics

import "fmt"

func main() {
	// Switch statement in go is (switch case default) (fallthrough)
	// switch expression {
	// case value1:
	// code to be executed if expression equals value1
	// fallthrough
	// case value2:
	// code to be executed if expression equals value2
	// case value3:
	// code to be executed if expression equals value3
	// default:
	// code to be executed if expression does not match any value
	// }

	// Switch statement in other languages (switch case break default)
	// switch expression {
	// case value1:
	// code to be executed if expression equals value1
	// break;
	// case value2:
	// code to be executed if expression equals value2
	// break;
	// case value3:
	// code to be executed if expression equals value3
	// break;
	// default:
	// code to be executed if expression does not match any value
	// break;
	// }

	// fruit := "apple"

	// switch fruit {
	// case "apple":
	//	fmt.Println("It's an apple.")
	// case "banana":
	//	fmt.Println("It's a banana.")
	// default:
	//	fmt.Println("Unknown Fruit!")
	// }

	// Multiple Conditions
	// day := "Monday"

	// switch day {
	// case "Monday", "Tuesday", "Wednesday", "Thursday", "Friday":
	//	fmt.Println("It's a weekday.")
	// case "Sunday":
	//	fmt.Println("It's a weekend.")
	// default:
	//	fmt.Println("Invalid day.")
	// }

	// number := 15

	// switch {
	// case number < 10:
	//	fmt.Println("Number is less than 10.")
	// case number >= 10 && number < 20:
	//	fmt.Println("The number is between 10 and 19.")
	// default:
	//	fmt.Println("Number is 20 or more.")
	// }

	// num := 2

	// switch {
	// case num > 1:
	//	fmt.Println("Greater than 1.")
	// case num == 2:
	//	fmt.Println("Number is 2.")
	// default:
	//	fmt.Println("Not Two.")
	// }

	checkType(10)
	checkType(3.14)
	checkType("hello")
	checkType(true)

	average(10)
}

func checkType(x interface{}) {
	switch x.(type) {
	case int:
		fmt.Println("It's an integer")
	case int32:
		fmt.Println("It's an integer 32")
	case float64:
		fmt.Println("It's float")
	case string:
		fmt.Println("It's a string")
	default:
		fmt.Println("Unknown Type")
	}
}

func average(x int) {
	switch {
	case x >= 90:
		fmt.Println("Grade A")
	case x >= 80:
		fmt.Println("Grade B")
	case x >= 70:
		fmt.Println("Grade C")
	default:
		fmt.Println("Grade D")
	}
}
