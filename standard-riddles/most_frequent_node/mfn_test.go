package most_frequent_node

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMostFrequentNode(t *testing.T) {
	tests := []struct {
		input *Node
		want  []int
	}{
		{
			input: &Node{
				id:       0,
				val:      0,
				children: []*Node{},
			},
			want: []int{0},
		},
		{
			input: &Node{
				id:       99,
				val:      0,
				children: []*Node{},
			},
			want: []int{99},
		},
		{
			input: &Node{
				id:  0,
				val: 1,
				children: []*Node{
					{
						id:  2,
						val: 99,
						children: []*Node{
							{
								id:       3,
								val:      99,
								children: []*Node{}},
						},
					},
				},
			},
			want: []int{2, 3},
		},
		{
			input: &Node{
				id:  0,
				val: 3,
				children: []*Node{
					{
						id:  1,
						val: 1,
						children: []*Node{
							{
								id:  4,
								val: 5,
								children: []*Node{
									{
										id:       8,
										val:      1,
										children: []*Node{},
									},
								},
							},
							{
								id:       5,
								val:      6,
								children: []*Node{},
							},
						},
					},
					{
						id:  2,
						val: 2,
						children: []*Node{
							{
								id:       6,
								val:      2,
								children: []*Node{},
							},
							{
								id:       7,
								val:      1,
								children: []*Node{},
							},
						},
					},
					{
						id:       3,
						val:      4,
						children: []*Node{},
					},
				},
			},
			want: []int{1, 8, 7},
		},
	}

	for _, tc := range tests {
		got := mostFrequentNode(tc.input)
		assert.Equal(t, tc.want, got)
	}
}
