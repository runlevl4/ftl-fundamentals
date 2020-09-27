// Package calculator provides a library for simple calculations in Go.
package calculator

import (
	"errors"
	"math"
)

// Add takes two numbers and returns the result of adding them together.
func Add(a, b float64) float64 {
	return a + b
}

// Subtract takes two numbers and returns the result of subtracting the second
// from the first.
func Subtract(a, b float64) float64 {
	return a - b
}

// Multiply takes two numbers and returns the result of multiplying them
// together.
func Multiply(a, b float64) float64 {
	return a * b
}

// Divide takes two numbers and returns the result of dividing them.
func Divide(a, b float64) (float64, error) {
	x := a / b

	if x == math.NaN() {
		return 0, errors.New("invalid input")
	}
	return x, nil
}

// Sqrt takes a number and returns its square root.
func Sqrt(a float64) (float64, error) {
	if a < 0 {
		return 0, errors.New("invalid input")
	}

	return math.Ceil(math.Sqrt(a)*100) / 100, nil
}
