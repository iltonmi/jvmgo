package math

import (
	"../../rtda"
	"../base"
	"math"
)

type DREM struct{ base.NoOperandsInstruction }
type FREM struct{ base.NoOperandsInstruction }
type IREM struct{ base.NoOperandsInstruction }
type LREM struct{ base.NoOperandsInstruction }

func (self *IREM) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	res := v1 % v2
	stack.PushInt(res)
}

func (self *LREM) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopLong()
	v1 := stack.PopLong()
	res := v1 % v2
	stack.PushLong(res)
}

/**
todo 浮点数类型有无限大值，即使除0也不会抛出异常
*/
func (self *FREM) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopFloat()
	v1 := stack.PopFloat()
	if v2 == 0.0 {
		panic("java.lang.ArithmeticException: / by zero")
	}
	res := float32(math.Mod(float64(v1), float64(v2)))
	stack.PushFloat(res)
}

func (self *DREM) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopDouble()
	v1 := stack.PopDouble()
	if v2 == 0.0 {
		panic("java.lang.ArithmeticException: / by zero")
	}
	res := math.Mod(v1, v2)
	stack.PushDouble(res)
}
