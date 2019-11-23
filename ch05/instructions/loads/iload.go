package loads

import "../base"
import "../../rtda"

type ILOAD struct{ base.Index8Instruction } //对大部分方法，局部变量表大小都不会超过256，因此用一个字节
type ILOAD_0 struct{ base.NoOperandsInstruction }
type ILOAD_1 struct{ base.NoOperandsInstruction }
type ILOAD_2 struct{ base.NoOperandsInstruction }
type ILOAD_3 struct{ base.NoOperandsInstruction }

/**
为了避免重复代码，定义一个函数供iload系列指令使用
*/
func _iload(frame *rtda.Frame, index uint) {
	val := frame.LocalVars().GetInt(index)
	frame.OperandStack().PushInt(val)
}

/**
iload指令的索引来自操作数
*/
func (self *ILOAD) Execute(frame *rtda.Frame) {
	_iload(frame, uint(self.Index))
}

/**
其余4条指令的索引隐含在操作码中
*/
func (self *ILOAD_0) Execute(frame *rtda.Frame) {
	_iload(frame, 0)
}
func (self *ILOAD_1) Execute(frame *rtda.Frame) {
	_iload(frame, 1)
}
func (self *ILOAD_2) Execute(frame *rtda.Frame) {
	_iload(frame, 2)
}
func (self *ILOAD_3) Execute(frame *rtda.Frame) {
	_iload(frame, 3)
}
