package heap

import (
	"../../classfile"
)

type MethodRef struct {
	MemberRef
	method *Method
}

func newMethodRef(cp *ConstantPool, refInfo *classfile.ConstantMethodRefInfo) *MethodRef {
	ref := &MethodRef{}
	ref.cp = cp
	ref.copyMemberRefInfo(&refInfo.ConstantMemberRefInfo)
	return ref
}

func (mr *MethodRef) ResolvedMethod() *Method {
	if mr.method == nil {
		mr.resolveMethodRef()
	}
	return mr.method
}

func (mr *MethodRef) resolveMethodRef() {
	currClass := mr.cp.class
	c := mr.ResolvedClass()
	if c.IsInterface() {
		panic("java.lang.IncompatibleClassChangeError")
	}
	method := lookupMethod(c, mr.name, mr.descriptor)
	if method == nil {
		panic("java.lang.NoSuchMethodError")
	}
	if !method.isAccessibleTo(currClass) {
		panic("java.lang.IllegalAccessError")
	}
	mr.method = method
}

func lookupMethod(class *Class, name, descriptor string) *Method {
	method := LookUpMethodInClass(class, name, descriptor)
	if method == nil {
		method = LookUpMethodInInterfaces(class.interfaces, name, descriptor)
	}
	return method
}
