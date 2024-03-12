package merge_two_sorted_lists

// 1 --> 2 --> 4
// 1 --> 3 --> 4
// out: 1 --> 1 --> 2 --> 3 --> 4 --> 4

// []
// []
// out: []

// []
// 0
// out: []

type Node struct {
	val  int
	next *Node
}

func mergeTwoSortedLists(input1 *Node, input2 *Node) *Node {
	result := &Node{}
	tail := result
	for input1 != nil && input2 != nil {
		if input1.val <= input2.val {
			tail.next = input1
			input1 = input1.next
		} else {
			tail.next = input2
			input2 = input2.next
		}
		tail = tail.next
	}

	// if multiple items remain in any one list, while the
	// other was processed all the way
	for input1 != nil && input1.next != nil {
		tail.next = input1
		input1 = input1.next
		tail = tail.next
	}
	for input2 != nil && input2.next != nil {
		tail.next = input2
		input2 = input2.next
		tail = tail.next
	}

	//last of lists that cannot be progressed (no next)
	if input1 != nil {
		tail.next = input1
	}
	if input2 != nil {
		tail.next = input2
	}

	return result.next
}
