package main

import "testing"

func Test_sumAll(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{name: "Нулевые значения", nums: []int{0,0,0,0}, want: 0},
		{name: "Отрицательные значения", nums: []int{1,2,3,4}, want: 10},
		{name: "Положительные значения", nums: []int{-1,-2}, want: -3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := sumAll(tt.nums...)
			if got != tt.want {
				t.Errorf("sumAll() = %v, want %v", got, tt.want)
			}
		})
	}
}
