package heap

import (
	"../../classfile"
)

type InterfaceMethodRef struct {
	MemberRef
	method *Method
}

func newInterfaceMethodRef(cp *ConstantPool, refInfo *classfile.ConstantInterfaceMethodRefInfo) *InterfaceMethodRef {
	ref := &InterfaceMethodRef{}
	ref.cp = cp
	ref.copyMemberRefInfo(&refInfo.ConstantMemberRefInfo)
	return ref
}

func (imr *InterfaceMethodRef) ResolvedInterfaceMethod() *Method {
	if imr.method == nil {
		imr.resolveInterfaceMethodRef()
	}
	return imr.method
}

func (imr *InterfaceMethodRef) resolveInterfaceMethodRef() {
	currIface := imr.cp.class
	c := imr.ResolvedClass()
	if !c.IsInterface() {
		panic("java.lang.IncompatibleClassChangeError")
	}
	method := lookUpInterfaceMethod(c, imr.name, imr.descriptor)
	if method == nil {
		panic("java.lang.NoSuchMethodError")
	}
	if !method.isAccessibleTo(currIface) {
		panic("java.lang.IllegalAccessError")
	}
	imr.method = method
}

func lookUpInterfaceMethod(iface *Class, name, descriptor string) *Method {
	for _, method := range iface.methods {
		if method.name == name && method.descriptor == descriptor {
			return method
		}
	}
	return LookUpMethodInInterfaces(iface.interfaces, name, descriptor)
}
