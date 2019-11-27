package classfile

/**

 */
type MemberInfo struct {
	cp              ConstantPool //保存常量池指针，后面会用到
	accessFlags     uint16
	nameIndex       uint16
	descriptorIndex uint16          //描述符用来描述字段的数据类型、方法的参数列表（包括数量、类型以及顺序）和返回值
	attributes      []AttributeInfo //假如某个字段是final static int m = 123, 可能会存在一项ConstantValue的属性，值指向123
}

/**
读取字段表或方法表
*/
func readMembers(reader *ClassReader, cp ConstantPool) []*MemberInfo {
	memberCount := reader.readUint16() //class文件中的fields_count
	members := make([]*MemberInfo, memberCount)
	for i := range members {
		members[i] = readMember(reader, cp)
	}
	return members
}

/**
读取字段或方法数据
*/
func readMember(reader *ClassReader, cp ConstantPool) *MemberInfo {
	return &MemberInfo{
		cp:              cp,
		accessFlags:     reader.readUint16(),
		nameIndex:       reader.readUint16(),
		descriptorIndex: reader.readUint16(),
		attributes:      readAttributes(reader, cp),
	}
}

func (self *MemberInfo) CodeAttribute() *CodeAttribute {
	for _, attrInfo := range self.attributes {
		switch attrInfo.(type) {
		case *CodeAttribute:
			return attrInfo.(*CodeAttribute)
		}
	}
	return nil
}

/**
从常量池查找字段或方法的名称
*/
func (self *MemberInfo) Name() string {
	return self.cp.getUtf8(self.nameIndex)
}

/**
从常量池查找字段或方法的描述符
*/
func (self *MemberInfo) Descriptor() string {
	return self.cp.getUtf8(self.descriptorIndex)
}

func (self *MemberInfo) ConstantValueAttribute() *ConstantValueAttribute {
	//只有常量才会有这个属性
	for _, attrInfo := range self.attributes {
		switch attrInfo.(type) {
		case *ConstantValueAttribute:
			return attrInfo.(*ConstantValueAttribute)
		}
	}
	return nil
}
func (self *MemberInfo) AccessFlags() uint16 {
	return self.accessFlags
}
