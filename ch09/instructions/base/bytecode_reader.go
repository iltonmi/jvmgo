package base

type ByteCodeReader struct {
	//存放字节码
	code []byte
	//记录读取到哪个字节
	pc int
}

/**
作用：避免每次解码指令都新创建一个BytecodeReader
*/
func (self *ByteCodeReader) Reset(code []byte, pc int) {
	self.code = code
	self.pc = pc
}

func (self *ByteCodeReader) SkipPadding() {
	for self.pc%4 != 0 {
		self.ReadUInt8()
	}
}

func (self *ByteCodeReader) ReadUInt8() uint8 {
	i := self.code[self.pc]
	self.pc++
	return i
}

func (self *ByteCodeReader) ReadInt8() int8 {
	return int8(self.ReadUInt8())
}

func (self *ByteCodeReader) ReadUInt16() uint16 {
	byte1 := uint16(self.ReadUInt8())
	byte2 := uint16(self.ReadUInt8())
	return (byte1 << 8) | byte2
}

func (self *ByteCodeReader) ReadInt16() int16 {
	return int16(self.ReadUInt16())
}

func (self *ByteCodeReader) ReadInt32() int32 {
	byte1 := int32(self.ReadUInt8())
	byte2 := int32(self.ReadUInt8())
	byte3 := int32(self.ReadUInt8())
	byte4 := int32(self.ReadUInt8())
	return (byte1 << 24) | (byte2 << 16) | (byte3 << 8) | byte4
}

func (self *ByteCodeReader) ReadInt32s(n int32) []int32 {
	ints := make([]int32, n)
	for i := range ints {
		ints[i] = self.ReadInt32()
	}
	return ints
}

func (self *ByteCodeReader) PC() int {
	return self.pc
}
