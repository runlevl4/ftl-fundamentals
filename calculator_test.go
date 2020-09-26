package calculator_test

import (
	"calculator"
	"testing"
)

type testCase struct {
	a, b float64
	want float64
	name string
}

func TestAdd(t *testing.T) {
	t.Parallel()

	testCases := []testCase{
		{a: 2, b: 2, want: 4, name: "Adding equal numbers."},
		{a: 10, b: 2, want: 12, name: "Adding smaller number to larger."},
		{a: 2, b: 8, want: 10, name: "Adding larger number to smaller."},
		{a: -4, b: 8, want: 4, name: "Adding negative and positive numbers."},
	}
	
	for _, tc := range testCases {
		t.Log(tc.name)
		got := calculator.Add(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("want %f, got %f", tc.want, got)
		}
	}

}

func TestSubtract(t *testing.T) {
	t.Parallel()

	testCases := []testCase{
		{a: 2, b: 2, want: 0, name: "Subtract equal numbers."},
		{a: 10, b: 2, want: 8, name: "Subtract smaller from larger."},
		{a: 2, b: 8, want: -6, name: "Subtract larger from smaller."},
		{a: -4, b: 8, want: -12, name: "Subtract positive from negative."},
	}

	for _, tc := range testCases {
		t.Log(tc.name)
		got := calculator.Subtract(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("want %f, got %f", tc.want, got)
		}
	}	
	
}

func TestMultiply(t *testing.T) {
	t.Parallel()

	testCases := []testCase{
		{a: 2, b: 2, want: 4, name: "Multiply equal numbers."},
		{a: 10, b: 2, want: 20, name: "Multiply larger by smaller."},
		{a: 2, b: 8, want: 16, name: "Multiply smaller by larger."},
		{a: -4, b: 8, want: -32, name: "Multiple positive and negative."},
	}

	for _, tc := range testCases {
		t.Log(tc.name)
		got := calculator.Multiply(tc.a, tc.b)	
		if tc.want != got {
			t.Errorf("want %f, got %f", tc.want, got)
		}
	}	
}
