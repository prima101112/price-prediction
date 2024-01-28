package algorithm

import (
	"math"
	"sort"
)

// Function to calculate the average of a slice of float64
func Average(numbers []float64) float64 {
	total := 0.0
	for _, num := range numbers {
		total += num
	}
	return total / float64(len(numbers))
}

// Function to calculate the median of a slice of float64
func Median(numbers []float64) float64 {
	// Sort the numbers
	sort.Float64s(numbers)

	// Calculate the middle index
	mid := len(numbers) / 2

	// If the length is odd, return the middle number
	if len(numbers)%2 == 1 {
		return numbers[mid]
	}

	// If the length is even, return the average of the two middle numbers
	return (numbers[mid-1] + numbers[mid]) / 2.0
}

// Function to find the lowest value in a slice of float64
func Lowest(numbers []float64) float64 {
	// Initialize lowest with positive infinity
	lowest := math.Inf(1)

	// Iterate through the slice to find the lowest value
	for _, num := range numbers {
		if num < lowest {
			lowest = num
		}
	}

	return lowest
}

// Function to find the highest value in a slice of float64
func Highest(numbers []float64) float64 {
	// Initialize highest with negative infinity
	highest := math.Inf(-1)

	// Iterate through the slice to find the highest value
	for _, num := range numbers {
		if num > highest {
			highest = num
		}
	}

	return highest
}
