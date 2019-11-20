package classfile

/*
//定长属性，只会出现在field_info结构，用于表示常量表达式的值
ConstantValue_attribute {
    u2 attribute_name_index;
//长度必须是2
    u4 attribute_length;
    u2 constantvalue_index;
}
*/
type ConstantValueAttribute struct {
	constantValueIndex uint16
}

func (self *ConstantValueAttribute) readInfo(reader *ClassReader) {
	self.constantValueIndex = reader.readUint16()
}

func (self *ConstantValueAttribute) ConstantValueIndex() uint16 {
	return self.constantValueIndex
}
