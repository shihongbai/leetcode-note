package tree

import (
	"fmt"
)

// 颜色常量
const (
	RED   = true
	BLACK = false
)

// 节点结构
type RBNode struct {
	value  int
	color  bool
	left   *RBNode
	right  *RBNode
	parent *RBNode
}

// 红黑树结构
type RBTree struct {
	root *RBNode
}

// 左旋
func (t *RBTree) leftRotate(x *RBNode) {
	y := x.right
	x.right = y.left
	if y.left != nil {
		y.left.parent = x
	}
	y.parent = x.parent
	if x.parent == nil {
		t.root = y
	} else if x == x.parent.left {
		x.parent.left = y
	} else {
		x.parent.right = y
	}
	y.left = x
	x.parent = y
}

// 右旋
func (t *RBTree) rightRotate(y *RBNode) {
	x := y.left
	y.left = x.right
	if x.right != nil {
		x.right.parent = y
	}
	x.parent = y.parent
	if y.parent == nil {
		t.root = x
	} else if y == y.parent.right {
		y.parent.right = x
	} else {
		y.parent.left = x
	}
	x.right = y
	y.parent = x
}

// 插入修复
func (t *RBTree) insertFix(z *RBNode) {
	for z.parent != nil && z.parent.color == RED {
		if z.parent == z.parent.parent.left {
			y := z.parent.parent.right
			if y != nil && y.color == RED {
				z.parent.color = BLACK
				y.color = BLACK
				z.parent.parent.color = RED
				z = z.parent.parent
			} else {
				if z == z.parent.right {
					z = z.parent
					t.leftRotate(z)
				}
				z.parent.color = BLACK
				z.parent.parent.color = RED
				t.rightRotate(z.parent.parent)
			}
		} else {
			y := z.parent.parent.left
			if y != nil && y.color == RED {
				z.parent.color = BLACK
				y.color = BLACK
				z.parent.parent.color = RED
				z = z.parent.parent
			} else {
				if z == z.parent.left {
					z = z.parent
					t.rightRotate(z)
				}
				z.parent.color = BLACK
				z.parent.parent.color = RED
				t.leftRotate(z.parent.parent)
			}
		}
	}
	t.root.color = BLACK
}

// 插入节点
func (t *RBTree) Insert(value int) {
	newRBNode := &RBNode{value: value, color: RED}
	if t.root == nil {
		t.root = newRBNode
	} else {
		current := t.root
		var parent *RBNode
		for current != nil {
			parent = current
			if value < current.value {
				current = current.left
			} else {
				current = current.right
			}
		}
		newRBNode.parent = parent
		if value < parent.value {
			parent.left = newRBNode
		} else {
			parent.right = newRBNode
		}
	}
	t.insertFix(newRBNode)
}

// 查找节点
func (t *RBTree) search(value int) *RBNode {
	current := t.root
	for current != nil && current.value != value {
		if value < current.value {
			current = current.left
		} else {
			current = current.right
		}
	}
	return current
}

// 删除修复
func (t *RBTree) deleteFix(x *RBNode) {
	for x != t.root && x.color == BLACK {
		if x == x.parent.left {
			sibling := x.parent.right
			if sibling.color == RED {
				sibling.color = BLACK
				x.parent.color = RED
				t.leftRotate(x.parent)
				sibling = x.parent.right
			}
			if sibling.left.color == BLACK && sibling.right.color == BLACK {
				sibling.color = RED
				x = x.parent
			} else {
				if sibling.right.color == BLACK {
					sibling.left.color = BLACK
					sibling.color = RED
					t.rightRotate(sibling)
					sibling = x.parent.right
				}
				sibling.color = x.parent.color
				x.parent.color = BLACK
				sibling.right.color = BLACK
				t.leftRotate(x.parent)
				x = t.root
			}
		} else {
			sibling := x.parent.left
			if sibling.color == RED {
				sibling.color = BLACK
				x.parent.color = RED
				t.rightRotate(x.parent)
				sibling = x.parent.left
			}
			if sibling.left.color == BLACK && sibling.right.color == BLACK {
				sibling.color = RED
				x = x.parent
			} else {
				if sibling.left.color == BLACK {
					sibling.right.color = BLACK
					sibling.color = RED
					t.leftRotate(sibling)
					sibling = x.parent.left
				}
				sibling.color = x.parent.color
				x.parent.color = BLACK
				sibling.left.color = BLACK
				t.rightRotate(x.parent)
				x = t.root
			}
		}
	}
	x.color = BLACK
}

// 删除节点
func (t *RBTree) Delete(value int) {
	z := t.search(value)
	if z == nil {
		return
	}

	var y, x *RBNode
	y = z
	originalColor := y.color
	if z.left == nil {
		x = z.right
		t.transplant(z, z.right)
	} else if z.right == nil {
		x = z.left
		t.transplant(z, z.left)
	} else {
		y = t.minimum(z.right)
		originalColor = y.color
		x = y.right
		if y.parent == z {
			if x != nil {
				x.parent = y
			}
		} else {
			t.transplant(y, y.right)
			y.right = z.right
			y.right.parent = y
		}
		t.transplant(z, y)
		y.left = z.left
		y.left.parent = y
		y.color = z.color
	}

	if originalColor == BLACK && x != nil {
		t.deleteFix(x)
	}
}

// 替换子树
func (t *RBTree) transplant(u, v *RBNode) {
	if u.parent == nil {
		t.root = v
	} else if u == u.parent.left {
		u.parent.left = v
	} else {
		u.parent.right = v
	}
	if v != nil {
		v.parent = u.parent
	}
}

// 找最小值
func (t *RBTree) minimum(n *RBNode) *RBNode {
	for n.left != nil {
		n = n.left
	}
	return n
}

// 中序遍历
func (t *RBTree) inorderTraversal(RBNode *RBNode) {
	if RBNode != nil {
		t.inorderTraversal(RBNode.left)
		fmt.Printf("%d ", RBNode.value)
		t.inorderTraversal(RBNode.right)
	}
}

// 打印树的中序遍历
func (t *RBTree) PrintInOrder() {
	t.inorderTraversal(t.root)
	fmt.Println()
}

func main() {
	tree := &RBTree{}
	values := []int{20, 15, 30, 10, 18, 25, 40}
	for _, v := range values {
		tree.Insert(v)
	}

	fmt.Println("红黑树的中序遍历:")
	tree.PrintInOrder()

	tree.Delete(15)
	fmt.Println("删除 15 后:")
	tree.PrintInOrder()
}
