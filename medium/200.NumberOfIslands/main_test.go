package numberofislands

import (
	"reflect"
	"testing"
)

func TestMaximumSafenessFactor(t *testing.T) {
	type test struct {
		grid [][]byte
		want int
	}

	tests := []test{
		{
			grid: [][]byte{
				{'1', '1', '1', '1', '0'},
				{'1', '1', '0', '1', '0'},
				{'1', '1', '0', '0', '0'},
				{'0', '0', '0', '0', '0'},
			},
			want: 2,
		},
		{
			grid: [][]byte{
				{'1', '1', '0', '0', '0'},
				{'1', '1', '0', '0', '0'},
				{'0', '0', '1', '0', '0'},
				{'0', '0', '0', '1', '1'},
			},
			want: 3,
		},
	}

	for _, tc := range tests {
		got := numIslands(tc.grid)
		if !reflect.DeepEqual(tc.want, got) {
			t.Fatalf("expected: %v, got: %v, case: %v", tc.want, got, tc.grid)
		}
	}

	return
}
