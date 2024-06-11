package main

import "fmt"

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func InOrderTraversal(root *TreeNode) {
	if root != nil {
		InOrderTraversal(root.Left)
		fmt.Println(root.Val)
		InOrderTraversal(root.Right)
	}
}

func PreOrderTraversal(root *TreeNode){
     if root != nil {
		fmt.Println(root.Val)
		PreOrderTraversal(root.Left)
		PreOrderTraversal(root.Right)
	 }
}

func PostOrdertraversal(root *TreeNode){
	if root != nil{
		PostOrdertraversal(root.Left)
		PostOrdertraversal(root.Right)
		fmt.Println(root.Val)
	}
}

func main() {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	root.Left.Left = &TreeNode{Val: 4}
	root.Left.Right = &TreeNode{Val: 5}

	fmt.Println("In-order Traversal:")
    InOrderTraversal(root)
    fmt.Println("Pre-order Traversal:")
    PreOrderTraversal(root)
    fmt.Println("Post-order Traversal:")
    PostOrdertraversal(root)
}