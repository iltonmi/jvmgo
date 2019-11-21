package math

import (
	"../../rtda"
	"../base"
)

// Increment local variable by constant
type IINC struct {
	//局部变量表索引
	Index uint
	//常量值
	Const int32
}

func (self *IINC) FetchOperands(reader *base.ByteCodeReader) {
	self.Index = uint(reader.ReadUInt8())
	self.Const = int32(reader.ReadInt8())
}

func (self *IINC) Execute(frame *rtda.Frame) {
	localVars := frame.LocalVars()
	val := localVars.GetInt(self.Index)
	val += self.Const
	localVars.SetInt(self.Index, val)
}
