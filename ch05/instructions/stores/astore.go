package loads

import "../base"
import "../../rtda"

type ASTORE struct{ base.Index8Instruction }
type ASTORE_0 struct{ base.NoOperandsInstruction }
type ASTORE_1 struct{ base.NoOperandsInstruction }
type ASTORE_2 struct{ base.NoOperandsInstruction }
type ASTORE_3 struct{ base.NoOperandsInstruction }

/**
  为了避免重复代码，定义一个函数供ASTORE系列指令使用
*/
func _astore(frame *rtda.Frame, index uint) {
	val := frame.OperandStack().PopDouble()
	frame.LocalVars().SetDouble(index, val)
}

/**
  ASTORE指令的索引来自操作数
*/
func (self *ASTORE) Execute(frame *rtda.Frame) {
	_astore(frame, uint(self.Index))
}

/**
  其余3条指令的索引隐含在操作码中
*/
func (self *ASTORE_0) Execute(frame *rtda.Frame) {
	_astore(frame, 0)
}
func (self *ASTORE_1) Execute(frame *rtda.Frame) {
	_astore(frame, 1)
}
func (self *ASTORE_2) Execute(frame *rtda.Frame) {
	_astore(frame, 2)
}
func (self *ASTORE_3) Execute(frame *rtda.Frame) {
	_astore(frame, 3)
}
