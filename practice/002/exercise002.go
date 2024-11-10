package main

import (
	"fmt"
	"log"
)

func main() {
	fmt.Println("Exercise 001")

	var input int
	fmt.Println("Please enter number: ")
	_, err := fmt.Scanln(&input)

	if err != nil {
		log.Fatal("Please enter a number")
	}

	result, err := Ex002(input)

	if err != nil {
		log.Fatalf("Error for input %v: %v", input, err)
	}

	fmt.Printf("Factorial of %d: %d", input, result)
}

func Ex002(input int) (uint64, error) {
	var factorial uint64 = 1

	if input < 0 {
		return 0, fmt.Errorf("Factorials for negative values are not allowed")
	}

	if input == 1 {
		return 1, nil
	}

	for i := 1; i <= input; i++ {
		factorial *= uint64(i)
	}
	return factorial, nil
}
