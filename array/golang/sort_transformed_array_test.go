package golang

import (
	"reflect"
	"sort"
	"testing"
)

func TestSortTransformedArray(t *testing.T) {
	tests := []struct {
		name    string
		nums    []int
		a, b, c int
		want    []int
	}{
		{"示例1 a>0", []int{-4, -2, 2, 4}, 1, 3, 5, []int{3, 9, 15, 33}},
		{"示例2 a<0", []int{-4, -2, 2, 4}, -1, 3, 5, []int{-23, -5, 1, 7}},
		{"a=0 b>0 递增", []int{-4, -2, 2, 4}, 0, 1, 0, []int{-4, -2, 2, 4}},
		{"a=0 b<0 递减", []int{-4, -2, 2, 4}, 0, -1, 0, []int{-4, -2, 2, 4}},
		{"a=0 b=0 全等", []int{-4, -2, 2, 4}, 0, 0, 7, []int{7, 7, 7, 7}},
		{"a>0 奇数个", []int{-3, -1, 2}, 1, 0, 0, []int{1, 4, 9}},
		{"a<0 奇数个", []int{-3, -1, 2}, -1, 0, 0, []int{-9, -4, -1}},
		{"单个元素", []int{5}, 2, 3, 1, []int{66}},
		{"空数组", []int{}, 1, 1, 1, []int{}},
		{"重复元素", []int{2, 2, 2}, 1, -4, 4, []int{0, 0, 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := sortTransformedArray(append([]int{}, tt.nums...), tt.a, tt.b, tt.c)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortTransformedArray(%v, %d, %d, %d) = %v, want %v",
					tt.nums, tt.a, tt.b, tt.c, got, tt.want)
			}
		})
	}
}

// 对拍：与「逐个算 f(x) 再排序」的朴素解法比对
func TestSortTransformedArrayBruteForce(t *testing.T) {
	for n := 0; n <= 7; n++ {
		nums := make([]int, n)
		for i := range nums {
			nums[i] = -n + i
		}
		for _, a := range []int{-2, -1, 0, 1, 2} {
			for _, b := range []int{-3, 0, 3} {
				for _, c := range []int{-5, 0, 5} {
					want := make([]int, n)
					for i, x := range nums {
						want[i] = a*x*x + b*x + c
					}
					sort.Ints(want)

					got := sortTransformedArray(append([]int{}, nums...), a, b, c)
					if !reflect.DeepEqual(got, want) {
						t.Fatalf("nums=%v a=%d b=%d c=%d: got %v, want %v", nums, a, b, c, got, want)
					}
				}
			}
		}
	}
}
