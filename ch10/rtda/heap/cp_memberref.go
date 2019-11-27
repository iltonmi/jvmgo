package heap

import (
	"../../classfile"
)

/**
存放字段和方法的符号引用共有的信息
*/
type MemberRef struct {
	SymRef
	name string
	//存放字段描述符的原因是：站在java虚拟机的角度，一个类是完全可以有多个同名的字段，只要他们的类型互不相同即可
	descriptor string
}

func (self *MemberRef) copyMemberRefInfo(refInfo *classfile.ConstantMemberRefInfo) {
	self.className = refInfo.ClassName()
	self.name, self.descriptor = refInfo.NameAndDesctiptor()
}

func (self *MemberRef) Name() string {
	return self.name
}
func (self *MemberRef) Descriptor() string {
	return self.descriptor
}
