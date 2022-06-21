package algorithm

type BSTreeNode struct {
	lson, rson *BSTreeNode
	val        int
}

type BSTree struct {
	root *BSTreeNode
}

func NewBSTreeNode(val int) *BSTreeNode {
	return &BSTreeNode{val: val}
}

func NewBSTree() *BSTree {
	return &BSTree{root: nil}
}

func (t *BSTree) Insert(val int) bool {
	node := NewBSTreeNode(val)
	if t.root == nil {
		t.root = node
		return true
	}
	var prev *BSTreeNode
	cur := t.root
	for cur != nil {
		prev = cur
		if cur.val < val {
			cur = cur.rson
		} else if cur.val > val {
			cur = cur.lson
		} else {
			return false
		}
	}
	if prev.val < val {
		prev.rson = node
	} else {
		prev.lson = node
	}
	return true
}

func (t *BSTree) Delete() {

}

func (t *BSTree) Search(val int) *BSTreeNode {
	cur := t.root
	for cur != nil {
		if cur.val < val {
			cur = cur.rson
		} else if cur.val > val {
			cur = cur.lson
		} else if cur.val == val {
			return cur
		}
	}
	return nil
}

func (t *BSTree) Traverse() {

}
