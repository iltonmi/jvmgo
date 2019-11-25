package loads

import "../base"
import "../../rtda"

type DLOAD struct{ base.Index8Instruction }
type DLOAD_0 struct{ base.NoOperandsInstruction }
type DLOAD_1 struct{ base.NoOperandsInstruction }
type DLOAD_2 struct{ base.NoOperandsInstruction }
type DLOAD_3 struct{ base.NoOperandsInstruction }

/**
为了避免重复代码，定义一个函数供fload系列指令使用
*/
func _dload(frame *rtda.Frame, index uint) {
	val := frame.LocalVars().GetDouble(index)
	frame.OperandStack().PushDouble(val)
}

/**
lload指令的索引来自操作数
*/
func (self *DLOAD) Execute(frame *rtda.Frame) {
	_dload(frame, uint(self.Index))
}

/**
其余3条指令的索引隐含在操作码中
*/
func (self *DLOAD_0) Execute(frame *rtda.Frame) {
	_dload(frame, 0)
}
func (self *DLOAD_1) Execute(frame *rtda.Frame) {
	_dload(frame, 1)
}
func (self *DLOAD_2) Execute(frame *rtda.Frame) {
	_dload(frame, 2)
}
func (self *DLOAD_3) Execute(frame *rtda.Frame) {
	_dload(frame, 3)
}
