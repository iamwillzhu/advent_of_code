package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"math"
	"os"
	"sort"
	"strconv"
)

func oneDayPartOne() {
	fmt.Println("Day 1 Part One")

	column1, column2, err := readColumnsFromCSV("1-day.csv")
	if err != nil {
		fmt.Println("Error reading CSV:", err)
		os.Exit(1)
	}

	// Sort both columns
	sort.Ints(column1)
	sort.Ints(column2)

	// Check if columns have equal length
	if len(column1) != len(column2) {
		fmt.Printf("Error: Columns have different lengths. Column 1: %d, Column 2: %d\n", len(column1), len(column2))
		os.Exit(1)
	}

	fmt.Printf("Sorted Column 1: %v\n", column1)
	fmt.Printf("Sorted Column 2: %v\n", column2)

	totalDistance, err := sumOfDifferences(column1, column2)
	if err != nil {
		fmt.Println("Error calculating total distance:", err)
		os.Exit(1)
	}

	fmt.Printf("Total distance: %d\n", totalDistance)
}

func oneDayPartTwo() {
	fmt.Println("Day 1 Part Two")

	column1, column2, err := readColumnsFromCSV("1-day.csv")
	if err != nil {
		fmt.Println("Error reading CSV:", err)
		os.Exit(1)
	}

	// Check if columns have equal length
	if len(column1) != len(column2) {
		fmt.Printf("Error: Columns have different lengths. Column 1: %d, Column 2: %d\n", len(column1), len(column2))
		os.Exit(1)
	}

	// Create a map to store the count of each number in column2
	column2Count := make(map[int]int)
	for _, num := range column2 {
		column2Count[num]++
	}

	// For each number in column1, print how many times it appears in column2
	similarityScore := 0
	for _, num := range column1 {
		similarityScore += column2Count[num] * num
	}

	fmt.Printf("Similarity score: %d\n", similarityScore)
}

func readColumnsFromCSV(filename string) ([]int, []int, error) {
	// Open the CSV file
	file, err := os.Open(filename)
	if err != nil {
		return nil, nil, fmt.Errorf("error opening file: %w", err)
	}
	defer file.Close()

	// Create a CSV reader
	reader := csv.NewReader(file)

	// Initialize two slices to store the columns
	var column1, column2 []int

	// Read and process each row
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, nil, fmt.Errorf("error reading record: %w", err)
		}

		// Convert string values to integers
		num1, err := strconv.Atoi(record[0])
		if err != nil {
			return nil, nil, fmt.Errorf("error converting first number: %w", err)
		}

		num2, err := strconv.Atoi(record[1])
		if err != nil {
			return nil, nil, fmt.Errorf("error converting second number: %w", err)
		}

		// Append numbers to respective slices
		column1 = append(column1, num1)
		column2 = append(column2, num2)
	}

	return column1, column2, nil
}

func sumOfDifferences(list1, list2 []int) (int, error) {
	// Check if lists have same length
	if len(list1) != len(list2) {
		return 0, fmt.Errorf("lists must have the same length")
	}

	// Calculate sum of differences
	total := 0
	for i := 0; i < len(list1); i++ {
		diff := math.Abs(float64(list1[i] - list2[i]))
		total += int(diff)
	}
	return total, nil
}
