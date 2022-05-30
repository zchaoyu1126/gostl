## 初衷

最近在学习go1.18推出的泛型功能，正好自己又苦于每次写代码时都要调用这样的函数

```
func AbsInt(x int) int
func AbsFloat64(x float64) float64
```

虽然可以使用反射来实现泛型的功能，但反射相对而言还是比较复杂，故在此结合实践来学习go泛型的知识。在编写的过程中，应编写相应的单元测试文件，以此来稳固之前学习的单元测试知识。



## Useage

```shell
go get github.com/zchaoyu1126/gostl
import "github.com/zchaoyu1126/gostl/queue"
```



## 计划利用泛型实现的内容

- [x] 队列
- [ ] 栈
- [ ] 线段树：刷题高频+1
- [ ] 字典树
- [x] 绝对值
- [x] Max，Min函数
- [ ] 希望使用泛型简化sort，但应该不太可能优化

> 字典树和线段树的原因是二者是我在刷题过程中遇到的高频题，但是一直不太熟练，希望借此机会能够完善一下，写成一个模板，同时能支持泛型就最好不过了。
>
> sort纯属是突然有了一个脑洞，甚至不知道自己的目标是什么，还需考虑。



## 实现效果

队列的实现效果如下：

（关于api文档，还在学习中...）

```go
package main

import (
	"fmt"

	"github.com/zchaoyu1126/gostl/queue"
)

type Node struct {
	Val int
	Idx int
}

func main() {
	q := queue.NewQueue[int]()
	q.PushBack(1)
	q.PushBack(2)
	q.PushBack(5)
	for !q.Empty() {
		val, _ := q.PopFront()
		fmt.Println(val)
	}

	nq := queue.NewQueue(3.1415926, 5.412)
	for nq.Size() != 0 {
		fmt.Println(nq.PopFront())
	}

	nodes := []*Node{{9, 1}, {6, 2}, {3, 3}}
	nodeQ := queue.NewQueue(nodes...)
	for !nodeQ.Empty() {
		node, _ := nodeQ.PopFront()
		fmt.Println(node.Val, node.Idx)
	}
}
```

