package medium

import (
	"fmt"
	"testing"

	"leetcode/binary_tree/util"
	commonUtil "leetcode/util"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/maximum-width-of-binary-tree/description/

type NodeInfo[T any] struct {
	node  *util.TreeNode[T]
	index int
}

// method 1 iterative BFS
// 1) define a struct NodeInfo to store node and index
// 2) use a queue to store NodeInfo
// 3) for each level, get the leftMostIndex and rightMostIndex, then calculate the maxWidth
// 4) return the maxWidth
// TC = O(N), SC = O(N), N is the number of nodes in tree
func widthOfBinaryTree1[T int](root *util.TreeNode[T]) (maxWidth int) {
	if root == nil {
		return 0
	}

	maxWidth = 0
	queue := []NodeInfo[T]{{root, 0}}

	for len(queue) > 0 {
		leveSize := len(queue)
		leftMostIndex := queue[0].index

		for i := 0; i < leveSize; i++ {
			nodeInfo := queue[0]
			queue = queue[1:]

			if nodeInfo.node.Left != nil {
				queue = append(queue, NodeInfo[T]{nodeInfo.node.Left, nodeInfo.index * 2})
			}

			if nodeInfo.node.Right != nil {
				queue = append(queue, NodeInfo[T]{nodeInfo.node.Right, nodeInfo.index*2 + 1})
			}

			// calculate the maxWidth of each level by the leftMostIndex and rightMostIndex
			if i == leveSize-1 {
				maxWidth = commonUtil.Max(maxWidth, nodeInfo.index-leftMostIndex+1)
			}
		}
	}

	return maxWidth
}

func Test_widthOfBinaryTree1(t *testing.T) {
	type args struct {
		root *util.TreeNode[int]
	}
	type expected struct {
		result int
	}
	type testCase struct {
		name     string
		args     args
		expected expected
	}

	testCases := []testCase{
		{
			name: "1",
			args: args{
				root: util.ArrayToBinaryTree([]int{1, 3, 2, 5, 3, -1, 9}),
			},
			expected: expected{
				result: 4,
			},
		},
		{
			name: "2",
			args: args{
				root: util.ArrayToBinaryTree([]int{1, 3, -1, 5, 3}),
			},
			expected: expected{
				result: 2,
			},
		},
		{
			name: "3",
			args: args{
				root: util.ArrayToBinaryTree([]int{1, 3, 2, 5}),
			},
			expected: expected{
				result: 2,
			},
		},
	}

	for _, tc := range testCases {
		assert.Equal(
			t,
			tc.expected.result,
			widthOfBinaryTree1(tc.args.root),
			fmt.Sprintf("testCase name: %s", tc.name),
		)
	}
}
