package classfile

/**
CONSTANT_Class_info {
    u1 tag;
//只存放常量池索引，指向一个Constant—_Utf8_info常量
    u2 name_index;
}
*/

type ConstantClassInfo struct {
	cp        ConstantPool
	nameIndex uint16
}

func (self *ConstantClassInfo) readInfo(reader *ClassReader) {
	self.nameIndex = reader.readUint16()
}

func (self *ConstantClassInfo) Name() string {
	return self.cp.getUtf8(self.nameIndex)
}
