package golang

import (
	"reflect"
	"testing"
)

func TestPlusOne(t *testing.T) {
	tests := []struct {
		name   string
		digits []int
		want   []int
	}{
		{"示例1", []int{1, 2, 3}, []int{1, 2, 4}},
		{"示例2", []int{4, 3, 2, 1}, []int{4, 3, 2, 2}},
		{"示例3 只有一位9", []int{9}, []int{1, 0}},
		{"进位后停在非9位", []int{8, 9, 9, 9}, []int{9, 0, 0, 0}},
		{"进位后停在非9位(短)", []int{1, 9, 9}, []int{2, 0, 0}},
		{"全9", []int{9, 9, 9}, []int{1, 0, 0, 0}},
		{"末位非9", []int{1, 2, 9, 8}, []int{1, 2, 9, 9}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := plusOne(tt.digits)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("plusOne(%v) = %v, want %v", tt.digits, got, tt.want)
			}
		})
	}
}
