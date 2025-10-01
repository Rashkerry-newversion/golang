package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Go CLI Calculator")
	fmt.Println("------------------")

	for {
		// Read first number (or 'q' to quit)
		fmt.Print("Enter first number (or 'q' to quit): ")
		input1, _ := reader.ReadString('\n')
		input1 = strings.TrimSpace(input1)
		if strings.EqualFold(input1, "q") {
			fmt.Println("Goodbye!")
			return
		}
		num1, err := strconv.ParseFloat(input1, 64)
		if err != nil {
			fmt.Println("Invalid number. Please enter a valid number (e.g. 12.34).")
			continue
		}

		// Read second number
		fmt.Print("Enter second number: ")
		input2, _ := reader.ReadString('\n')
		input2 = strings.TrimSpace(input2)
		num2, err := strconv.ParseFloat(input2, 64)
		if err != nil {
			fmt.Println("Invalid number. Please enter a valid number.")
			continue
		}

		// Read operation
		fmt.Print("Choose operation (+, -, *, /): ")
		op, _ := reader.ReadString('\n')
		op = strings.TrimSpace(op)

		// Compute result
		var result float64
		valid := true

		switch op {
		case "+":
			result = num1 + num2
		case "-":
			result = num1 - num2
		case "*":
			result = num1 * num2
		case "/":
			if num2 == 0 {
				fmt.Println("Error: division by zero is not allowed.")
				valid = false
			} else {
				result = num1 / num2
			}
		default:
			fmt.Println("Unknown operation. Use +, -, * or /.")
			valid = false
		}

		// Print result if operation was valid
		if valid {
			// Remove trailing .0 for integers (optional cosmetic)
			if result == float64(int64(result)) {
				fmt.Printf("Result: %d\n\n", int64(result))
			} else {
				fmt.Printf("Result: %g\n\n", result)
			}
		}
	}
}