package calculator_test

import (
	"calculator"
	"testing"
)

type testCase struct {
	a, b float64
	want float64
}

func TestAdd(t *testing.T) {
	t.Parallel()

	testCases := []testCase{
		{a: 2, b: 2, want: 4},
		{a: 10, b: 2, want: 12},
		{a: 2, b: 8, want: 10},
		{a: -4, b: 8, want: 4},
	}
	
	for _, tc := range testCases {
		got := calculator.Add(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("want %f, got %f", tc.want, got)
		}
	}

}

func TestSubtract(t *testing.T) {
	t.Parallel()

	testCases := []testCase{
		{a: 2, b: 2, want: 0},
		{a: 10, b: 2, want: 8},
		{a: 2, b: 8, want: -6},
		{a: -4, b: 8, want: -12},
	}

	for _, tc := range testCases {
		got := calculator.Subtract(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("want %f, got %f", tc.want, got)
		}
	}	
	
}

func TestMultiply(t *testing.T) {
	t.Parallel()

	testCases := []testCase{
		{a: 2, b: 2, want: 4},
		{a: 10, b: 2, want: 20},
		{a: 2, b: 8, want: 16},
		{a: -4, b: 8, want: -32},
	}

	for _, tc := range testCases {
		got := calculator.Multiply(tc.a, tc.b)	
		if tc.want != got {
			t.Errorf("want %f, got %f", tc.want, got)
		}
	}	
}
