package references

import (
	"../../rtda"
	"../../rtda/heap"
	"../base"
)

// Create new array
type ANEW_ARRAY struct {
	base.Index16Instruction
}

func (self *ANEW_ARRAY) Execute(frame *rtda.Frame) {
	cp := frame.Method().Class().ConstantPool()
	classRef := cp.GetConstant(self.Index).(*heap.ClassRef)
	//component of the new array
	componentClass := classRef.ResolvedClass()
	stack := frame.OperandStack()
	//弹出数组长度
	count := stack.PopInt()
	if count < 0 {
		panic("java.lang.NegativeArraySizeException")
	}
	//get the class of the array with elements of componentClass
	arrClass := componentClass.ArrayClass()
	arr := arrClass.NewArray(uint(count))
	stack.PushRef(arr)
}
