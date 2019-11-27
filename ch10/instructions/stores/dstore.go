package stores

import "../base"
import "../../rtda"

type DSTORE struct{ base.Index8Instruction }
type DSTORE_0 struct{ base.NoOperandsInstruction }
type DSTORE_1 struct{ base.NoOperandsInstruction }
type DSTORE_2 struct{ base.NoOperandsInstruction }
type DSTORE_3 struct{ base.NoOperandsInstruction }

/**
为了避免重复代码，定义一个函数供DSTORE系列指令使用
*/
func _dstore(frame *rtda.Frame, index uint) {
	val := frame.OperandStack().PopDouble()
	frame.LocalVars().SetDouble(index, val)
}

/**
DSTORE指令的索引来自操作数
*/
func (self *DSTORE) Execute(frame *rtda.Frame) {
	_dstore(frame, uint(self.Index))
}

/**
其余3条指令的索引隐含在操作码中
*/
func (self *DSTORE_0) Execute(frame *rtda.Frame) {
	_dstore(frame, 0)
}
func (self *DSTORE_1) Execute(frame *rtda.Frame) {
	_dstore(frame, 1)
}
func (self *DSTORE_2) Execute(frame *rtda.Frame) {
	_dstore(frame, 2)
}
func (self *DSTORE_3) Execute(frame *rtda.Frame) {
	_dstore(frame, 3)
}
