package main

import (
	"fmt"
	"math"
)

func calculate(num1 float64, action string, num2 float64) (float64, error) {
	switch action {
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
		return 0, fmt.Errorf("invalid operation")
	}
}

func main() {
	for {
		var num1, num2 float64
		var action string
		fmt.Println("Input left operand: ")
		_, err := fmt.Scanln(&num1)
		if err != nil {
			fmt.Printf("Error: invalid number input\n")
			continue
		}
		fmt.Println("Input operation: ")
		_, err = fmt.Scanln(&action)
		if err != nil {
			fmt.Printf("Error: invalid operation input\n")
			continue
		}
		fmt.Println("Input right operand: ")
		_, err = fmt.Scanln(&num2)
		if err != nil {
			fmt.Printf("Error: invalid number input\n")
			continue
		}
		result, err := calculate(num1, action, num2)
		if err != nil {
			fmt.Println("Error:", err)
		} else {
			if math.Mod(result, 1) == 0 {
				fmt.Printf("Result: %.0f\n", result)
			} else {
				fmt.Printf("Result: %.3f\n", result)
			}
		}
	}
}
