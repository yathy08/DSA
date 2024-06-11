package main
import "fmt"

type TreeNode struct {
	val int
	left *TreeNode
	right *TreeNode
}

type BST struct{
	root *TreeNode
}

func (bst *BST) Insert(val int) {
	bst.root = insert(bst.root,val)
}

func insert(node *TreeNode,val int) *TreeNode {
	if node == nil {
		return &TreeNode{val : val}
	}
	if val < node.val {
		node.left = insert(node.left, val)
	} else {
		node.right = insert(node.right, val)
	}
	return node
}
func (bst *BST) Contains(val int) bool {
	return contains(bst.root, val)
}

func contains(node *TreeNode, val int) bool {
	if node == nil {
		return false
	}
	if val == node.val {
		return true
	}
	if val < node.val {
		return contains(node.left, val)
	}
	return contains(node.right, val)
}

func (bst *BST) Delete(val int) {
	bst.root = delete(bst.root, val)
}

func delete(node *TreeNode, val int) *TreeNode {
	if node == nil {
		return nil
	}
	if val < node.val {
		node.left = delete(node.left, val)
	} else if val > node.val {
		node.right = delete(node.right, val)
	} else {
		if node.left == nil {
			return node.right
		} else if node.right == nil {
			return node.left
		}
		minNode := findMin(node.right)
		node.val = minNode.val
		node.right = delete(node.right, minNode.val)
	}
	return node
}

func findMin(node *TreeNode) *TreeNode {
	for node.left != nil {
		node = node.left
	}
	return node
}

func (bst *BST) InOrder() {
	inOrder(bst.root)
	fmt.Println()
}

func inOrder(node *TreeNode) {
	if node == nil {
		return
	}
	inOrder(node.left)
	fmt.Print(node.val, " ")
	inOrder(node.right)
}

func (bst *BST) PreOrder() {
	preOrder(bst.root)
	fmt.Println()
}

func preOrder(node *TreeNode) {
	if node == nil {
		return
	}
	fmt.Print(node.val, " ")
	preOrder(node.left)
	preOrder(node.right)
}

func (bst *BST) PostOrder() {
	postOrder(bst.root)
	fmt.Println()
}

func postOrder(node *TreeNode) {
	if node == nil {
		return
	}
	postOrder(node.left)
	postOrder(node.right)
	fmt.Print(node.val, " ")
}

func main() {
	bst := &BST{}
	bst.Insert(5)
	bst.Insert(3)
	bst.Insert(7)
	bst.Insert(2)
	bst.Insert(4)
	bst.Insert(6)
	bst.Insert(8)

	fmt.Println("InOrder traversal:")
	bst.InOrder()

	fmt.Println("PreOrder traversal:")
	bst.PreOrder()

	fmt.Println("PostOrder traversal:")
	bst.PostOrder()

	fmt.Println("Contains 4:", bst.Contains(4))
	fmt.Println("Contains 9:", bst.Contains(9))

	bst.Delete(5)

	fmt.Println("InOrder traversal after deleting 5:")
	bst.InOrder()
}