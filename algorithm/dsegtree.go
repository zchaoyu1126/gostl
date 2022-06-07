package algorithm

import "github.com/zchaoyu1126/gostl/cmath"

type RealNumber interface {
	uint8 | uint16 | uint32 | uint64 | int8 | int16 | int32 | int64 | float32 | float64
}

type DSegTreeNode[T RealNumber] struct {
	lson, rson *DSegTreeNode[T]
	l, r, m    int
	add        T
	sum        T
	maxNum     T
}

type DSegTree[T RealNumber] struct {
	root *DSegTreeNode[T]
}

func NewDSegTreeNode[T RealNumber](l, r int) *DSegTreeNode[T] {
	return &DSegTreeNode[T]{
		l: l,
		r: r,
		m: (l + r) >> 1,
	}
}

func NewDSegTree[T RealNumber](l, r int) *DSegTree[T] {
	return &DSegTree[T]{root: NewDSegTreeNode[T](l, r)}
}

func (t *DSegTree[T]) QuerySum(l, r int, node *DSegTreeNode[T]) T {
	if l > r {
		return 0
	}
	if l <= node.l && node.r <= r {
		return node.sum
	}
	t.pushDown(node)
	var sum T
	if l <= node.m {
		sum += t.QuerySum(l, r, node.lson)
	}
	if r >= node.m {
		sum += t.QuerySum(l, r, node.rson)
	}
	return sum
}

func (t *DSegTree[T]) QueryMax(l, r int, node *DSegTreeNode[T]) T {
	if l > r {
		return 0
	}
	if l <= node.l && node.r <= r {
		return node.maxNum
	}
	t.pushDown(node)
	var maxN T
	if l <= node.m {
		maxN = cmath.Max(maxN, t.QueryMax(l, r, node.lson))
	}
	if r >= node.m {
		maxN = cmath.Max(maxN, t.QueryMax(l, r, node.rson))
	}
	return maxN
}

func (t *DSegTree[T]) pushDown(node *DSegTreeNode[T]) {
	if node.lson == nil {
		node.lson = NewDSegTreeNode[T](node.l, node.m)
	}
	if node.rson == nil {
		node.rson = NewDSegTreeNode[T](node.m+1, node.r)
	}
	if node.add != 0 {
		node.lson.add += node.add
		node.rson.add += node.add
		node.lson.sum += T(float64(node.lson.r-node.lson.l+1) * float64(node.add))
		node.rson.sum += T(float64(node.rson.r-node.rson.l+1) * float64(node.add))
		node.lson.maxNum += node.add
		node.rson.maxNum += node.add
		node.add = 0
	}
}

func (t *DSegTree[T]) pushUp(node *DSegTreeNode[T]) {
	node.maxNum = cmath.Max(node.lson.maxNum, node.rson.maxNum)
	node.sum = node.lson.sum + node.rson.sum
}

func (t *DSegTree[T]) Modify(l, r int, v T, node *DSegTreeNode[T]) {
	if l > r {
		return
	}
	if l <= node.l && node.r <= r {
		node.sum += T(float64(node.r-node.l+1) * float64(v))
		node.add += v
		node.maxNum += v
	}
	t.pushDown(node)
	if l <= node.m {
		t.Modify(l, r, v, node.lson)
	}
	if r > node.m {
		t.Modify(l, r, v, node.rson)
	}
	t.pushUp(node)
}
