package heap

// jvms8 6.5.instanceof
// jvms8 6.5.checkcast
func (self *Class) isAssignableFrom(other *Class) bool {
	varClass, assignedToClass := other, self

	if varClass == assignedToClass {
		return true
	}

	if !varClass.IsArray() {
		if !varClass.IsInterface() {
			if !assignedToClass.IsInterface() {
				// both are not type of array or interface
				varClass.IsSubClassOf(assignedToClass)
			} else {
				// varClass is not type of array or interface, assignedToClass is interface type
				varClass.isSubInterfaceOf(assignedToClass)
			}
		} else {
			if !assignedToClass.IsInterface() {
				// varClass is interface type, assignedToClass is not interface type or array type
				//result depends on whether assignedToClass is type of java.lang.Object
				return assignedToClass.isJlObject()
			} else {
				// both are interface type
				return assignedToClass.isSubInterfaceOf(varClass)
			}
		}
	} else {
		if !assignedToClass.IsArray() {
			if !assignedToClass.IsInterface() {
				// varClass is array, assignedToClass is not interface type or array type
				//result depends on whether assignedToClass is type of java.lang.Object
				return assignedToClass.isJlObject()
			} else {
				// both are interface type
				//result depends on whether assignedToClass is type of
				// java.lang.Cloneable or java.io.Serializable
				return assignedToClass.isJlCloneable() || assignedToClass.isJioSerializable()
			}
		} else {
			// both are array
			//result depends on
			//1. both are same primitive type
			//2. both are reference type and varArrComponent can be converted to assignedToArrComponent
			varArrComponent := varClass.ComponentClass()
			assignedToArrComponent := assignedToClass.ComponentClass()
			return varArrComponent == assignedToArrComponent || assignedToArrComponent.isAssignableFrom(varArrComponent)
		}
	}
	return false
}

// self extends c
func (self *Class) IsSubClassOf(other *Class) bool {
	for c := self.superClass; c != nil; c = c.superClass {
		if c == other {
			return true
		}
	}
	return false
}

// self implements iface
func (self *Class) IsImplements(iface *Class) bool {
	for c := self; c != nil; c = c.superClass {
		for _, i := range c.interfaces {
			if i == iface || i.isSubInterfaceOf(iface) {
				return true
			}
		}
	}
	return false
}

// self extends iface
func (self *Class) isSubInterfaceOf(iface *Class) bool {
	for _, superInterface := range self.interfaces {
		if superInterface == iface || superInterface.isSubInterfaceOf(iface) {
			return true
		}
	}
	return false
}

// c extends self
func (self *Class) IsSuperClassOf(other *Class) bool {
	return other.IsSubClassOf(self)
}
