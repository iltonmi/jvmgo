package constants

import "../base"
import "../../rtda"

/**
统一扩展成int型再推入栈顶
*/
type BIPUSH struct {
	val int8
}

type SIPUSH struct {
	val int16
}

func (self *BIPUSH) FetchOperands(reader *base.ByteCodeReader) {
	self.val = reader.ReadInt8()
}

func (self *BIPUSH) Execute(frame *rtda.Frame) {
	i := int32(self.val)
	frame.OperandStack().PushInt(i)
}

func (self *SIPUSH) FetchOperands(reader *base.ByteCodeReader) {
	self.val = reader.ReadInt16()
}

func (self *SIPUSH) Execute(frame *rtda.Frame) {
	i := int32(self.val)
	frame.OperandStack().PushInt(i)
}
