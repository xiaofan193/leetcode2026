package golang

import "testing"

func TestNthUglyNumber(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{"示例1", 10, 12},
		{"示例2", 1, 1},
		{"第2个", 2, 2},
		{"第7个", 7, 8},
		{"n=11 跳过 11", 11, 15},
		{"第15个", 15, 24},
		{"边界上限", 1690, 2123366400},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := nthUglyNumber(tt.n); got != tt.want {
				t.Errorf("nthUglyNumber(%d) = %d, want %d", tt.n, got, tt.want)
			}
		})
	}
}
