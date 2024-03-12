package merge_two_sorted_lists

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMergeTwoSortedLists(t *testing.T) {
	tests := []struct {
		input1 *Node
		input2 *Node
		want   *Node
	}{
		{
			input1: &Node{val: 1,
				next: &Node{val: 2,
					next: &Node{val: 4}}},
			input2: &Node{val: 1,
				next: &Node{val: 3,
					next: &Node{val: 4}}},
			want: &Node{val: 1,
				next: &Node{val: 1,
					next: &Node{val: 2,
						next: &Node{val: 3,
							next: &Node{val: 4,
								next: &Node{val: 4}}}}}},
		},
		{
			input1: nil,
			input2: nil,
			want:   nil,
		},
		{
			input1: nil,
			input2: &Node{val: 1},
			want:   &Node{val: 1},
		},
		{
			input1: &Node{val: 1},
			input2: &Node{val: 1},
			want:   &Node{val: 1, next: &Node{val: 1}},
		},
		{
			input1: &Node{val: 1,
				next: &Node{val: 2,
					next: &Node{val: 4}}},
			input2: &Node{val: 1,
				next: &Node{val: 3,
					next: &Node{val: 4,
						next: &Node{val: 5,
							next: &Node{val: 6}}}}},
			want: &Node{val: 1,
				next: &Node{val: 1,
					next: &Node{val: 2,
						next: &Node{val: 3,
							next: &Node{val: 4,
								next: &Node{val: 4,
									next: &Node{val: 5,
										next: &Node{val: 6}}}}}}}},
		},
	}

	for _, tc := range tests {
		got := mergeTwoSortedLists(tc.input1, tc.input2)
		assert.Equal(t, tc.want, got)
	}
}
