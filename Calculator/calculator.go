package calculator

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	// Step 1: Get Input
	if len(os.Args) != 4 {
		fmt.Println("Usage: calculator <number> <operation> <number>")
		return
	}

	num1, err := strconv.ParseFloat(os.Args[1], 64)
	if err != nil {
		fmt.Println("Error: First number is not valid.")
		return
	}

	num2, err := strconv.ParseFloat(os.Args[3], 64)
	if err != nil {
		fmt.Println("Error: Second number is not valid.")
		return
	}

	operator := os.Args[2]

	// Step 2: Calculate
	result, err := calculate(num1, num2, operator)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Step 3: Show Result
	fmt.Printf("Result: %.2f\n", result)
}

func calculate(num1, num2 float64, operator string) (float64, error) {
	switch operator {
	case "+":
		return num1 + num2, nil
	case "-":
		return num1 - num2, nil
	case "*":
		return num1 * num2, nil
	case "/":
		if num2 == 0 {
			return 0, fmt.Errorf("division by zero")
		}
		return num1 / num2, nil
	default:
		return 0, fmt.Errorf("unknown operator: %s", operator)
	}
}

// We import the necessary packages: fmt for printing messages, os for getting command-line arguments, and strconv for converting strings to numbers.
// In the main function, we check if the user provided three command-line arguments (two numbers and an operator). If not, we print usage instructions and exit.
// We parse the first and third arguments (numbers) as floating-point numbers using strconv.ParseFloat.
// We extract the operator from the second argument.
// We call the calculate function with the parsed numbers and operator. This function performs the actual calculation.
// The calculate function uses a switch statement to perform different calculations based on the operator.
// We handle division by zero as a special case to avoid runtime errors.
// Finally, we print the result with two decimal places using fmt.Printf.?
