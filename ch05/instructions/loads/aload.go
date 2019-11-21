package loads

import "../base"
import "../../rtda"

type ALOAD struct{ base.Index8Instruction }
type ALOAD_0 struct{ base.NoOperandsInstruction }
type ALOAD_1 struct{ base.NoOperandsInstruction }
type ALOAD_2 struct{ base.NoOperandsInstruction }
type ALOAD_3 struct{ base.NoOperandsInstruction }

/**
为了避免重复代码，定义一个函数供fload系列指令使用
*/
func _aload(frame *rtda.Frame, index uint) {
	val := frame.LocalVars().GetRef(index)
	frame.OperandStack().PushRef(val)
}

/**
lload指令的索引来自操作数
*/
func (self *ALOAD) Execute(frame *rtda.Frame) {
	_aload(frame, uint(self.Index))
}

/**
其余3条指令的索引隐含在操作码中
*/
func (self *ALOAD_0) Execute(frame *rtda.Frame) {
	_aload(frame, 0)
}
func (self *ALOAD_1) Execute(frame *rtda.Frame) {
	_aload(frame, 1)
}
func (self *ALOAD_2) Execute(frame *rtda.Frame) {
	_aload(frame, 2)
}
func (self *ALOAD_3) Execute(frame *rtda.Frame) {
	_aload(frame, 3)
}
