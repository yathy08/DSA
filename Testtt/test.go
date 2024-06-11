package main

import "fmt"

type TreeNode struct{
	Val int
	Left *TreeNode
	Right *TreeNode
}

func (root *TreeNode) Insert(val int)*TreeNode{
	if root == nil{
		return &TreeNode{Val: val}
	}
	if val < root.Val {
		root.Left = root.Left.Insert(val)
	} else {
		root.Right= root.Right.Insert(val)
	}
	return root
}

func (root *TreeNode) Contains(val int)bool{
	if root == nil {
		return false
	}
	if root.Val == val{
		return true
	}  else if val < root.Val{
		return root.Left.Contains(val)
	} else {
		return root.Right.Contains(val)
	} 
}

func main() {
	root := &TreeNode{}

	root = root.Insert(1)
	root.Left = root.Left.Insert(2)
	root.Right = root.Right.Insert(3)
	root.Left.Left = root.Left.Left.Insert(4)
	root.Left.Right = root.Left.Right.Insert(5)
	fmt.Println("Does BST contain 9",root.Contains(9))
}