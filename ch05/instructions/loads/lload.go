package loads

import "../base"
import "../../rtda"

type LLOAD struct{ base.Index8Instruction }
type LLOAD_0 struct{ base.NoOperandsInstruction }
type LLOAD_1 struct{ base.NoOperandsInstruction }
type LLOAD_2 struct{ base.NoOperandsInstruction }
type LLOAD_3 struct{ base.NoOperandsInstruction }

/**
为了避免重复代码，定义一个函数供lload系列指令使用
*/
func _lload(frame *rtda.Frame, index uint) {
	val := frame.LocalVars().GetLong(index)
	frame.OperandStack().PushLong(val)
}

/**
lload指令的索引来自操作数
*/
func (self *LLOAD) Execute(frame *rtda.Frame) {
	_lload(frame, uint(self.Index))
}

/**
其余3条指令的索引隐含在操作码中
*/
func (self *LLOAD_0) Execute(frame *rtda.Frame) {
	_lload(frame, 0)
}
func (self *LLOAD_1) Execute(frame *rtda.Frame) {
	_lload(frame, 1)
}
func (self *LLOAD_2) Execute(frame *rtda.Frame) {
	_lload(frame, 2)
}
func (self *LLOAD_3) Execute(frame *rtda.Frame) {
	_lload(frame, 3)
}
