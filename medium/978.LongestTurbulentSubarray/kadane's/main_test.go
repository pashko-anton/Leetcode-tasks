package kadane_s

import (
	"reflect"
	"testing"
)

func TestMerge(t *testing.T) {
	type test struct {
		nums []int
		want int
	}

	tests := []test{
		{
			nums: []int{9, 4, 2, 10, 7, 8, 8, 1, 9},
			want: 5,
		},
		{
			nums: []int{4, 8, 12, 16},
			want: 2,
		},
		{
			nums: []int{0, 8, 45, 88, 48, 68, 28, 55, 17, 24},
			want: 8,
		},
	}

	for _, tc := range tests {
		got := maxTurbulenceSize(tc.nums)
		if !reflect.DeepEqual(tc.want, got) {
			t.Fatalf("expected: %v, got: %v, case: %v", tc.want, got, tc.nums)
		}
	}
}
