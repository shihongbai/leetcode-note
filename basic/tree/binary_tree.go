package tree

import "fmt"

type Node struct {
	value int
	left  *Node
	right *Node
}

type BinaryTree struct {
	root *Node
}

func (t *BinaryTree) postOrder(node *Node) {
	if node == nil {
		return
	}

	t.postOrder(node.left)
	t.postOrder(node.right)
	fmt.Print(node.value, " ")
}

func (t *BinaryTree) inOrder(node *Node) {
	if node == nil {
		return
	}

	t.inOrder(node.left)
	fmt.Print(node.value, " ")
	t.inOrder(node.right)
}

func (t *BinaryTree) preOrder(node *Node) {
	if node == nil {
		return
	}

	fmt.Print(node.value, " ")
	t.preOrder(node.left)
	t.preOrder(node.right)
}

func NewBinaryTree(root *Node) *BinaryTree {
	return &BinaryTree{root: root}
}
