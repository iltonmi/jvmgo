package classfile

import "encoding/binary"

/**
[]byte类型的包装，用作读取数据
*/
type ClassReader struct {
	data []byte
}

/**
读入u1类型数据
*/
func (self *ClassReader) readUint8() uint8 {
	val := self.data[0]
	//跳过已经读取的数据
	self.data = self.data[1:]
	return val
}

/**
读入u2类型数据
*/
func (self *ClassReader) readUint16() uint16 {
	val := binary.BigEndian.Uint16(self.data)
	self.data = self.data[2:]
	return val
}

/**
读入u4类型数据
*/
func (self *ClassReader) readUint32() uint32 {
	val := binary.BigEndian.Uint32(self.data)
	self.data = self.data[4:]
	return val
}

/**
读入uint64类型数据(JVMS没有规定u8)
*/
func (self *ClassReader) readUint64() uint64 {
	val := binary.BigEndian.Uint64(self.data)
	self.data = self.data[8:]
	return val
}

/**
读入uint16表,表的大小由开头的uint16数据指出
*/
func (self *ClassReader) readUint16s() []uint16 {
	n := self.readUint16() //interfaces_count
	s := make([]uint16, n)
	for i := range s {
		s[i] = self.readUint16()
	}
	return s
}

/**
读取指定数量的字节
*/
func (self *ClassReader) readBytes(n uint32) []byte {
	bytes := self.data[:n]
	self.data = self.data[n:]
	return bytes
}
