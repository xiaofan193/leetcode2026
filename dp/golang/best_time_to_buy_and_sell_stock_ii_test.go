package golang

import "testing"

func TestMaxProfit_ii(t *testing.T) {
	tests := []struct {
		name   string
		prices []int
		want   int
	}{
		{"示例1", []int{7, 1, 5, 3, 6, 4}, 7},
		{"示例2", []int{1, 2, 3, 4, 5}, 4},
		{"示例3", []int{7, 6, 4, 3, 1}, 0},
		{"多笔交易", []int{3, 3, 5, 0, 0, 3, 1, 4}, 8},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProfit_ii(tt.prices); got != tt.want {
				t.Errorf("maxProfit_ii(%v) = %v, want %v", tt.prices, got, tt.want)
			}
		})
	}
}
