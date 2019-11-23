package stores

import "../base"
import "../../rtda"

type ISTORE struct{ base.Index8Instruction }
type ISTORE_0 struct{ base.NoOperandsInstruction }
type ISTORE_1 struct{ base.NoOperandsInstruction }
type ISTORE_2 struct{ base.NoOperandsInstruction }
type ISTORE_3 struct{ base.NoOperandsInstruction }

/**
为了避免重复代码，定义一个函数供ISTORE系列指令使用
*/
func _istore(frame *rtda.Frame, index uint) {
	val := frame.OperandStack().PopInt()
	frame.LocalVars().SetInt(index, val)
}

/**
ISTORE指令的索引来自操作数
*/
func (self *ISTORE) Execute(frame *rtda.Frame) {
	_istore(frame, uint(self.Index))
}

/**
其余3条指令的索引隐含在操作码中
*/
func (self *ISTORE_0) Execute(frame *rtda.Frame) {
	_istore(frame, 0)
}
func (self *ISTORE_1) Execute(frame *rtda.Frame) {
	_istore(frame, 1)
}
func (self *ISTORE_2) Execute(frame *rtda.Frame) {
	_istore(frame, 2)
}
func (self *ISTORE_3) Execute(frame *rtda.Frame) {
	_istore(frame, 3)
}
