package golang

import (
	"reflect"
	"testing"
)

func TestMerge(t *testing.T) {
	tests := []struct {
		name  string
		nums1 []int
		m     int
		nums2 []int
		n     int
		want  []int
	}{
		{"示例1", []int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3, []int{1, 2, 2, 3, 5, 6}},
		{"示例2", []int{1}, 1, []int{}, 0, []int{1}},
		{"示例3", []int{0}, 0, []int{1}, 1, []int{1}},
		{"nums2 全更小", []int{4, 5, 6, 0, 0, 0}, 3, []int{1, 2, 3}, 3, []int{1, 2, 3, 4, 5, 6}},
		{"有相等元素", []int{1, 1, 1, 0, 0}, 3, []int{1, 2}, 2, []int{1, 1, 1, 1, 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			merge(tt.nums1, tt.m, tt.nums2, tt.n)
			if !reflect.DeepEqual(tt.nums1, tt.want) {
				t.Errorf("merge() nums1 = %v, want %v", tt.nums1, tt.want)
			}
		})
	}
}
