package constants

import "../base"
import "../../rtda"

/**
统一扩展成int型再推入栈顶
*/
type BIBUSH struct {
	val int8
}

type SIBUSH struct {
	val int16
}

func (self *BIBUSH) FetchOperands(reader *base.ByteCodeReader) {
	self.val = reader.ReadInt8()
}

func (self *BIBUSH) Execute(frame *rtda.Frame) {
	i := int32(self.val)
	frame.OperandStack().PushInt(i)
}

func (self *SIBUSH) FetchOperands(reader *base.ByteCodeReader) {
	self.val = reader.ReadInt16()
}

func (self *SIBUSH) Execute(frame *rtda.Frame) {
	i := int32(self.val)
	frame.OperandStack().PushInt(i)
}
