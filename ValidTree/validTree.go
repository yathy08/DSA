package main

import (
	"fmt"
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBSTHelper(node *TreeNode, min, max int) bool {
	if node == nil {
		return true
	}
	if node.Val <= min || node.Val >= max {
		return false
	}
	return isValidBSTHelper(node.Left, min, node.Val) && isValidBSTHelper(node.Right, node.Val, max)
}

func isValidBST(root *TreeNode) bool {
	return isValidBSTHelper(root, math.MinInt64, math.MaxInt64)
}


func main() {
	root := &TreeNode{Val: 10}
	root.Left = &TreeNode{Val: 5}
	root.Right = &TreeNode{Val: 15}
	root.Left.Left = &TreeNode{Val: 2}
	root.Left.Right = &TreeNode{Val: 6}
	root.Right.Left = &TreeNode{Val: 8}
	root.Right.Right = &TreeNode{Val: 8}

	fmt.Println("IS this Tree a valid BST:",isValidBST(root))
}
