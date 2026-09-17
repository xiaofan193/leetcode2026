package golang

import (
	"math/rand"
	"sort"
	"testing"
)

// kthSmallestBruteForce 枚举所有组合再排序，只用来对拍 kthSmallest
func kthSmallestBruteForce(mat [][]int, k int) int {
	// sums 保存"前几行"能凑出的全部和
	sums := []int{0}
	for _, row := range mat {
		next := make([]int, 0, len(sums)*len(row))
		for _, s := range sums {
			for _, v := range row {
				next = append(next, s+v)
			}
		}
		sums = next
	}
	sort.Ints(sums)
	return sums[k-1]
}

func TestKthSmallest(t *testing.T) {
	tests := []struct {
		name string
		mat  [][]int
		k    int
		want int
	}{
		{"示例1", [][]int{{1, 3, 11}, {2, 4, 6}}, 5, 7},
		{"示例2", [][]int{{1, 3, 11}, {2, 4, 6}}, 9, 17},
		{"示例3", [][]int{{1, 10, 10}, {1, 4, 5}, {2, 3, 6}}, 7, 9},
		{"示例4", [][]int{{1, 1, 10}, {2, 2, 9}}, 7, 12},
		{"单元素矩阵", [][]int{{5}}, 1, 5},
		{"只有一行", [][]int{{1, 2, 3}}, 2, 2},
		{"全是重复值", [][]int{{2, 2}, {2, 2}}, 4, 4},
		{"k 超过一行的长度", [][]int{{1, 1, 2}, {1, 2, 3}}, 9, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := kthSmallest(tt.mat, tt.k); got != tt.want {
				t.Errorf("kthSmallest(%v, %d) = %d, want %d", tt.mat, tt.k, got, tt.want)
			}
		})
	}
}

func TestKthSmallestAgainstBruteForce(t *testing.T) {
	rng := rand.New(rand.NewSource(1439))
	for range 500 {
		m := rng.Intn(4) + 1
		n := rng.Intn(4) + 1

		// 每行随机取值后排序，保证矩阵每行非递减
		mat := make([][]int, m)
		for i := range mat {
			row := make([]int, n)
			for j := range row {
				row[j] = rng.Intn(20) + 1
			}
			sort.Ints(row)
			mat[i] = row
		}

		// k 的合法范围是 1..n^m（每行选 1 个，共 n^m 种组合）
		total := 1
		for range mat {
			total *= n
		}
		k := rng.Intn(total) + 1

		want := kthSmallestBruteForce(mat, k)
		if got := kthSmallest(mat, k); got != want {
			t.Fatalf("mat=%v k=%d: got %d, want %d", mat, k, got, want)
		}
	}
}
