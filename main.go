package main

import (
	"fmt"

	"github.com/zchaoyu1126/gostl/algorithm"
)

func main() {
	tree := algorithm.NewAVLTree()
	for i := 100; i >= 0; i-- {
		tree.Insert(i*i, i)
	}

	arr := []int{}
	tree.Traverse(&arr)
	fmt.Println(tree.IsBalanced())
	fmt.Println(arr)
}
