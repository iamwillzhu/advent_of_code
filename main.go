package main

import "fmt"

func main() {
	var choice int

	fmt.Println("Please select a program to run:")
	fmt.Println("1. Program One")
	fmt.Println("2. Program Two")
	fmt.Println("3. Program Three")
	fmt.Print("Enter your choice (1-3): ")

	fmt.Scan(&choice)

	switch choice {
	case 1:
		fmt.Println("Running Program One")
		fmt.Println("Please select which part to run:")
		fmt.Println("1. Part One")
		fmt.Println("2. Part Two")
		fmt.Print("Enter your choice (1-2): ")

		var partChoice int
		fmt.Scan(&partChoice)

		switch partChoice {
		case 1:
			fmt.Println("Running Part One")
			oneDayPartOne()
		case 2:
			fmt.Println("Running Part Two")
			oneDayPartTwo()
		default:
			fmt.Println("Invalid choice. Please select either 1 or 2.")
		}
	case 2:
		fmt.Println("Running Program Two")
		// Call your second program function here
	case 3:
		fmt.Println("Running Program Three")
		// Call your third program function here
	default:
		fmt.Println("Invalid choice. Please select a number between 1 and 3.")
	}
}
