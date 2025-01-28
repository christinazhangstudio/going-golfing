package most_frequent_node

import "fmt"

// 0:3
// 0 --> 1:1   2:2   3:4
// 1 --> 4:5   5:6
// 2 --> 6:2   7:1
// 4 --> 8:1

// returns 1, 7, 8 if going bfs
// returns 1, 8, 7 if going dfs

type Node struct {
	id       int
	val      int // val might be a value that is returned as some metadata from some call
	children []*Node
}

func recurse(root *Node, valMap map[int][]int) {
	for _, child := range root.children {
		valMap[child.val] = append(valMap[child.val], child.id)
		if len(child.children) != 0 {
			recurse(child, valMap)
		}
	}
}

func mostFrequentNode(root *Node) []int {
	// val to their node ids
	// 1 --> 1, 7, 8
	valMap := make(map[int][]int)

	valMap[root.val] = []int{root.id}
	recurse(root, valMap)

	fmt.Println(valMap)

	var max int
	var idsToReturn []int
	for _, v := range valMap {
		if len(v) > max {
			// found a better max
			max = len(v)
			idsToReturn = v
		} else if len(v) == max {
			idsToReturn = append(idsToReturn, v...)
		}
	}

	return idsToReturn
}
