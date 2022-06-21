package algorithm

import (
	"github.com/zchaoyu1126/gostl/cmath"
)

type AVLTreeNode struct {
	h          int          // h 高度信息
	bf         int          // bf 平衡因子
	val, index int          // val 存储在该节点的值，index索引信息
	lson, rson *AVLTreeNode // lson,rson 左右儿子
}

type AVLTree struct {
	root *AVLTreeNode
}

func NewAVLTree() *AVLTree {
	return new(AVLTree)
}

func (t *AVLTree) Root() *AVLTreeNode {
	return t.root
}

func (t *AVLTree) Insert(val, index int) {
	t.root = t.root.insert(val, index)
}

func (t *AVLTree) Delete(index int) {
	t.root.delete(index)
}

func (t *AVLTree) Find(index int) {
	
}
func (t *AVLTree) IsBalanced() bool {
	return t.root.isBalanced()
}

func (t *AVLTree) Traverse(arr *[]int) {
	var traverse func(*AVLTreeNode)
	traverse = func(node *AVLTreeNode) {
		if node == nil {
			return
		}
		traverse(node.lson)
		*arr = append(*arr, node.val)
		traverse(node.rson)
	}
	traverse(t.root)
}

func (node *AVLTreeNode) updateHeight() {
	if node.lson == nil && node.rson == nil {
		node.h = 1
		return
	}
	if node.lson == nil {
		node.h = node.rson.h + 1
		return
	}
	if node.rson == nil {
		node.h = node.lson.h + 1
		return
	}
	node.h = cmath.Max(node.lson.h, node.rson.h) + 1
}

func (node *AVLTreeNode) updateBalanceFactor() {
	if node.lson == nil && node.rson == nil {
		node.bf = 0
		return
	}
	if node.lson == nil {
		node.bf = -node.rson.h
		return
	}
	if node.rson == nil {
		node.bf = node.lson.h
		return
	}
	node.bf = node.lson.h - node.rson.h
}

func (node *AVLTreeNode) rightRotate() *AVLTreeNode {
	newRoot := node.lson
	node.lson = newRoot.rson
	newRoot.rson = node
	// update newRoot and node's height, balance factor
	node.updateHeight()
	newRoot.updateHeight()
	node.updateBalanceFactor()
	newRoot.updateBalanceFactor()
	return newRoot
}

func (node *AVLTreeNode) leftRotate() *AVLTreeNode {
	newRoot := node.rson
	node.rson = newRoot.lson
	newRoot.lson = node
	// update newRoot and node's height, balance factor
	node.updateHeight()
	newRoot.updateHeight()
	node.updateBalanceFactor()
	newRoot.updateBalanceFactor()
	return newRoot
}

func (node *AVLTreeNode) fix() *AVLTreeNode {
	if node.bf > 1 || node.bf < -1 {
		if node.bf > 1 && node.lson.bf > 0 { // LL情况
			node = node.rightRotate()
		}
		if node.bf < -1 && node.rson.bf < 0 { // RR情况
			node = node.leftRotate()
		}
		if node.bf > 1 && node.lson.bf < 0 { // LR情况
			node.lson = node.lson.leftRotate()
			node = node.rightRotate()
		}
		if node.bf < -1 && node.rson.bf > 0 { // RL情况
			node.rson = node.rson.rightRotate()
			node = node.leftRotate()
		}
	}
	return node
}

func (node *AVLTreeNode) insert(val, index int) *AVLTreeNode {
	if node == nil {
		node = &AVLTreeNode{h: 1, bf: 0, val: val, index: index}
		return node
	}
	if index < node.index {
		node.lson = node.lson.insert(val, index)
	} else if index > node.index {
		node.rson = node.rson.insert(val, index)
	}
	node.updateBalanceFactor()
	node.updateHeight()
	node = node.fix()
	return node
}

func (node *AVLTreeNode) delete(index int) *AVLTreeNode {
	if node == nil {
		return node
	}
	if index < node.index {
		node.lson = node.lson.delete(index)
	} else if index > node.index {
		node.rson = node.rson.delete(index)
	} else {
		if node.lson != nil && node.rson != nil {
			// 左右都还有结点时，获取左子树的右下角即最大值
			cur := node.lson
			for cur.rson != nil {
				cur = cur.rson
			}
			node.val = cur.val
			node.index = cur.index
			node.lson = node.lson.delete(cur.index)
		} else {
			if node.lson != nil {
				node = node.lson
			} else if node.rson != nil {
				node = node.rson
			} else {
				return nil
			}
		}
	}
	node = node.fix()
	return node
}

func (node *AVLTreeNode) isBalanced() bool {
	if node == nil {
		return true
	}
	if node.bf > 1 || node.bf < -1 {
		return false
	}
	return node.lson.isBalanced() && node.rson.isBalanced()
}
