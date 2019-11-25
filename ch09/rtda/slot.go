package rtda

import (
	"./heap"
)

type Slot struct {
	//存放整数
	num int32
	//存放引用
	ref *heap.Object
}
