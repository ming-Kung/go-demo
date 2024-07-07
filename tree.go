package main

import "fmt"

type TreeNode struct {
	data  int
	left  *TreeNode
	right *TreeNode
}

//构建二叉树
func (t *TreeNode) BuildTree(arr []int) {
	if len(arr) > 0 {
		for _, val := range arr {
			if t.data == 0 {
				t.data = val
			} else {
				node := &TreeNode{
					data: val,
				}
				treeNode := t
				for true {
					if val <= treeNode.data {
						if treeNode.left != nil {
							treeNode = treeNode.left
						} else {
							treeNode.left = node
							break
						}
					} else if val > treeNode.data {
						if treeNode.right != nil {
							treeNode = treeNode.right
						} else {
							treeNode.right = node
							break
						}
					}
				}
			}
		}
	}
}

//插入节点
func (t *TreeNode) InsertNode(data int) {
	node := t
	for true {
		if data <= node.data {
			if node.left != nil {
				node = node.left
			} else {
				node.left = &TreeNode{
					data: data,
				}
				break
			}
		} else {
			if node.right != nil {
				node = node.right
			} else {
				node.right = &TreeNode{
					data: data,
				}
				break
			}
		}
	}
}

//删除节点
func (t *TreeNode) DelNode(data int) *TreeNode {
	node := t
	var (
		father    *TreeNode = nil
		direction string    = ""
	)
	for true {
		if data == node.data {
			//删除节点有左右子节点
			if node.left != nil && node.right != nil {
				//右边子树的最小节点（或左边子树的最大节点）提升为当前节点
				f := node
				n := node.right
				for true {
					if n.left == nil {
						break
					}
					f = n
					n = n.left
				}
				if node != f {
					f.left = nil
				}
				n.left = node.left
				n.right = node.right.right
				if father == nil {
					return n
				} else {
					if direction == "left" {
						father.left = n
					} else if direction == "right" {
						father.right = n
					}
				}
				break
			} else if node.left != nil || node.right != nil {
				//删除节点只有一个子节点
				if node.left != nil {
					if father == nil {
						return node.left
					}
					if direction == "left" {
						father.left = node.left
					} else if direction == "right" {
						father.right = node.left
					}
				} else {
					if father == nil {
						return node.right
					}
					if direction == "left" {
						father.left = node.right
					} else if direction == "right" {
						father.right = node.right
					}
				}
				break
			} else {
				//删除节点没有子节点
				if father == nil {
					return nil
				}
				if direction == "left" {
					father.left = nil
				} else if direction == "right" {
					father.right = nil
				}
				break
			}
		} else if data < node.data {
			if node.left != nil {
				father = node
				direction = "left"
				node = node.left
			} else {
				break
			}
		} else if data > node.data {
			if node.right != nil {
				father = node
				direction = "right"
				node = node.right
			} else {
				break
			}
		}
	}
	return t
}

//前序遍历
func (t *TreeNode) PreOrder(root *TreeNode) {
	if root == nil {
		return
	}
	fmt.Println(root.data)
	if root.left != nil {
		t.PreOrder(root.left)
	}
	if root.right != nil {
		t.PreOrder(root.right)
	}
}

//中序遍历
func (t *TreeNode) CurrentOrder(root *TreeNode) {
	if root == nil {
		return
	}
	if root.left != nil {
		t.CurrentOrder(root.left)
	}
	fmt.Println(root.data)
	if root.right != nil {
		t.CurrentOrder(root.right)
	}
}

//后续遍历
func (t *TreeNode) LastOrder(root *TreeNode) {
	if root == nil {
		return
	}
	if root.left != nil {
		t.LastOrder(root.left)
	}
	if root.right != nil {
		t.LastOrder(root.right)
	}
	fmt.Println(root.data)
}
func main() {
	tree := &TreeNode{}
	tree.BuildTree([]int{51, 32, 43, 12, 66})
	tree.InsertNode(62)
	//tree.PreOrder(tree)
	//tree.CurrentOrder(tree)
	//tree.LastOrder(tree)
	tree.PreOrder(tree.DelNode(32))
}
