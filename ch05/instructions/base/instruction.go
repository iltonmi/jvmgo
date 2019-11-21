package base

import "../../rtda"

type Instruction interface {
	FetchOperands(reader *ByteCodeReader)
	Execute(frame *rtda.Frame)
}

/**
无操作数指令
*/
type NoOperandsInstruction struct{}

func (self *NoOperandsInstruction) FetchOperands(reader *ByteCodeReader) {

}

/**
跳转指令
*/
type BranchInstruction struct {
	//跳转偏移量
	Offset int
}

func (self *BranchInstruction) FetchOperands(reader *ByteCodeReader) {
	self.Offset = int(reader.ReadInt16())
}

/**
给出局部变量表索引
*/
type Index8Instruction struct {
	//局部变量表索引
	Index uint
}

func (self *Index8Instruction) FetchOperands(reader *ByteCodeReader) {
	self.Index = uint(reader.ReadUInt8())
}

/**
给出常量池索引
*/
type Index16Instruction struct {
	//局部变量表索引
	Index uint
}

func (self *Index16Instruction) FetchOperands(reader *ByteCodeReader) {
	self.Index = uint(reader.ReadUInt16())
}
