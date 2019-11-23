package stores

import "../base"
import "../../rtda"

type LSTORE struct{ base.Index8Instruction }
type LSTORE_0 struct{ base.NoOperandsInstruction }
type LSTORE_1 struct{ base.NoOperandsInstruction }
type LSTORE_2 struct{ base.NoOperandsInstruction }
type LSTORE_3 struct{ base.NoOperandsInstruction }

/**
为了避免重复代码，定义一个函数供LSTORE系列指令使用
*/
func _lstore(frame *rtda.Frame, index uint) {
	val := frame.OperandStack().PopLong()
	frame.LocalVars().SetLong(index, val)
}

/**
LSTORE指令的索引来自操作数
*/
func (self *LSTORE) Execute(frame *rtda.Frame) {
	_lstore(frame, uint(self.Index))
}

/**
其余3条指令的索引隐含在操作码中
*/
func (self *LSTORE_0) Execute(frame *rtda.Frame) {
	_lstore(frame, 0)
}
func (self *LSTORE_1) Execute(frame *rtda.Frame) {
	_lstore(frame, 1)
}
func (self *LSTORE_2) Execute(frame *rtda.Frame) {
	_lstore(frame, 2)
}
func (self *LSTORE_3) Execute(frame *rtda.Frame) {
	_lstore(frame, 3)
}
