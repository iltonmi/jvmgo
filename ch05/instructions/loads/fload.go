package loads

import "../base"
import "../../rtda"

type FLOAD struct{ base.Index8Instruction }
type FLOAD_0 struct{ base.NoOperandsInstruction }
type FLOAD_1 struct{ base.NoOperandsInstruction }
type FLOAD_2 struct{ base.NoOperandsInstruction }
type FLOAD_3 struct{ base.NoOperandsInstruction }

/**
为了避免重复代码，定义一个函数供fload系列指令使用
*/
func _fload(frame *rtda.Frame, index uint) {
	val := frame.LocalVars().GetFloat(index)
	frame.OperandStack().PushFloat(val)
}

/**
lload指令的索引来自操作数
*/
func (self *FLOAD) Execute(frame *rtda.Frame) {
	_fload(frame, uint(self.Index))
}

/**
其余3条指令的索引隐含在操作码中
*/
func (self *FLOAD_0) Execute(frame *rtda.Frame) {
	_fload(frame, 0)
}
func (self *FLOAD_1) Execute(frame *rtda.Frame) {
	_fload(frame, 1)
}
func (self *FLOAD_2) Execute(frame *rtda.Frame) {
	_fload(frame, 2)
}
func (self *FLOAD_3) Execute(frame *rtda.Frame) {
	_fload(frame, 3)
}
