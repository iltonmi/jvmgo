package classfile

import "fmt"

type ClassFile struct {
	magic        uint32
	minorVersion uint16
	majorVersion uint16
	constantPool ConstantPool
	accessFlags  uint16
	thisClass    uint16
	superClass   uint16
	interfaces   []uint16
	fields       []*MemberInfo
	methods      []*MemberInfo
	attributes   []AttributeInfo
}

/**
将[]byte解析成ClassFile结构体
*/
func Parse(classData []byte) (cf *ClassFile, err error) {
	defer func() {
		//panic-recover机制
		//recover的返回值表示当前goroutine是否有panic行为
		if r := recover(); r != nil {
			var ok bool
			_, ok = r.(error)
			if !ok {
				_ = fmt.Errorf("%v", r)
			}
		}
	}()
	cr := &ClassReader{classData}
	cf = &ClassFile{}
	cf.read(cr)
	return
}

/**
依次调用其他方法，解析class文件
*/
func (self *ClassFile) read(reader *ClassReader) {
	self.readAndCheckMagic(reader)
	self.readAndCheckVersion(reader)
	self.constantPool = readConstantPool(reader)
	self.accessFlags = reader.readUint16()
	self.thisClass = reader.readUint16()
	self.superClass = reader.readUint16()
	self.interfaces = reader.readUint16s()
	self.fields = readMembers(reader, self.constantPool)
	self.methods = readMembers(reader, self.constantPool)
	self.attributes = readAttributes(reader, self.constantPool)
}

//getter
func (self *ClassFile) MajorVersion() uint16 {
	return self.majorVersion
}

//getter
func (self *ClassFile) MinorVersion() uint16 {
	return self.minorVersion
}

//getter
func (self *ClassFile) ConstantPool() ConstantPool {
	return self.constantPool
}

//getter
func (self *ClassFile) AccessFlags() uint16 {
	return self.accessFlags
}

//getter
func (self *ClassFile) Fields() []*MemberInfo {
	return self.fields
}

//getter
func (self *ClassFile) Methods() []*MemberInfo {
	return self.methods
}

/**
从常量池查找类名
*/
func (self *ClassFile) ClassName() string {
	return self.constantPool.getClassName(self.thisClass)
}

/**
从常量池查找超类名
*/
func (self *ClassFile) SuperClassName() string {
	if self.superClass > 0 {
		return self.constantPool.getClassName(self.superClass)
	}
	return "" //只要java.lang.Object没有超类
}

/**
从常量池查找接口名
*/
func (self *ClassFile) InterfaceNames() []string {
	interfaceNames := make([]string, len(self.interfaces))
	for i, cpIndex := range self.interfaces {
		interfaceNames[i] = self.constantPool.getClassName(cpIndex)
	}
	return interfaceNames
}

/**
class文件的魔数是0xCAFEBABE
现在还没有抛出异常的功能，调用panic方法直接终止程序执行
*/
func (self *ClassFile) readAndCheckMagic(reader *ClassReader) {
	magic := reader.readUint32()
	if magic != 0xCAFEBABE {
		panic("java.lang.ClassFormatError: magic!")
	}
}

/**
只支持版本号为45.0-52.0的class文件，即java1.0.2-1.1-1.4-5.0-8
遇到不支持的版本号则抛出java.lang.UnsupportedClassVersionError异常
*/
func (self *ClassFile) readAndCheckVersion(reader *ClassReader) {
	self.minorVersion = reader.readUint16()
	self.majorVersion = reader.readUint16()
	switch self.majorVersion {
	case 45:
		return
		//次版本号在1.2之后基本没用过了
	case 46, 47, 48, 49, 50, 51, 52:
		if self.minorVersion == 0 {
			return
		}
	}
	panic("java.lang.UnsupportedClassVersionError")
}

func (self *ClassFile) SourceFileAttribute() *SourceFileAttribute {
	for _, attrInfo := range self.attributes {
		switch attrInfo.(type) {
		case *SourceFileAttribute:
			return attrInfo.(*SourceFileAttribute)
		}
	}
	return nil
}
