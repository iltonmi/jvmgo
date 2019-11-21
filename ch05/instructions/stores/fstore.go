package loads

import "../base"
import "../../rtda"

type FSTORE struct{ base.Index8Instruction }
type FSTORE_0 struct{ base.NoOperandsInstruction }
type FSTORE_1 struct{ base.NoOperandsInstruction }
type FSTORE_2 struct{ base.NoOperandsInstruction }
type FSTORE_3 struct{ base.NoOperandsInstruction }

/**
为了避免重复代码，定义一个函数供FSTORE系列指令使用
*/
func _fstore(frame *rtda.Frame, index uint) {
	val := frame.OperandStack().PopFloat()
	frame.LocalVars().SetFloat(index, val)
}

/**
FSTORE指令的索引来自操作数
*/
func (self *FSTORE) Execute(frame *rtda.Frame) {
	_fstore(frame, uint(self.Index))
}

/**
其余3条指令的索引隐含在操作码中
*/
func (self *FSTORE_0) Execute(frame *rtda.Frame) {
	_fstore(frame, 0)
}
func (self *FSTORE_1) Execute(frame *rtda.Frame) {
	_fstore(frame, 1)
}
func (self *FSTORE_2) Execute(frame *rtda.Frame) {
	_fstore(frame, 2)
}
func (self *FSTORE_3) Execute(frame *rtda.Frame) {
	_fstore(frame, 3)
}
