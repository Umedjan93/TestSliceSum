package main

import "testing"

func TestSliceTotal(t *testing.T) {

	tests := []struct {
		name    string
		Slice   []int
		SumWant int
	}{
		{name: "All of numbers are positive", Slice: []int{5, 4, 3, 2, 1}, SumWant: 15},
		{name: "All of numbers are negative", Slice: []int{-1, -2, -3, -4, -5}, SumWant: -15},
		{name: "All numbers are mixed", Slice: []int{-1, 1, -2, 2, 0}, SumWant: 0},
	}
	for _, tt := range tests {
		SliceSum := SliceTotal(tt.Slice)
		{
			if SliceSum != tt.SumWant {
				t.Errorf("Index of SliceSum test %s, SliceSum %d and SumWant %d",
					tt.name, SliceSum, tt.SumWant)
			}
		}
	}
}