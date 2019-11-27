package references

import (
	"../../rtda"
	"../../rtda/heap"
	"../base"
)

type MULTI_ANEW_ARRAY struct {
	index      uint16
	dimensions uint8
}

func (self *MULTI_ANEW_ARRAY) FetchOperands(reader *base.ByteCodeReader) {
	self.index = reader.ReadUInt16()
	self.dimensions = reader.ReadUInt8()
}

func (self *MULTI_ANEW_ARRAY) Execute(frame *rtda.Frame) {
	cp := frame.Method().Class().ConstantPool()
	classRef := cp.GetConstant(uint(self.index)).(heap.ClassRef)
	//与anewarray指令不同，解析类符号引用得到的是数组类而不是数组元素的类
	arrClass := classRef.ResolvedClass()
	stack := frame.OperandStack()
	counts := popAndCheckCounts(stack, int(self.dimensions))
	array := newMultiDimensionalArray(counts, arrClass)
	stack.PushRef(array)
}

func popAndCheckCounts(stack *rtda.OperandStack, dimensions int) []int32 {
	counts := make([]int32, dimensions)
	for i := dimensions - 1; i >= 0; i-- {
		counts[i] = stack.PopInt()
		//确保每一个维度的数组长度都大于0
		if counts[i] < 0 {
			panic("java.lang.NegativeArraySizeException")
		}
		return counts
	}
	return counts
}

func newMultiDimensionalArray(counts []int32, arrClass *heap.Class) *heap.Object {
	count := uint(counts[0])
	arr := arrClass.NewArray(count)
	if len(counts) > 1 {
		refs := arr.Refs()
		for i := range refs {
			refs[i] = newMultiDimensionalArray(counts[1:], arrClass.ComponentClass())
		}
	}
	return arr
}
