package heap

//todo 找到的方法万一是抽象方法怎么办,这和接口方法同理
func LookUpMethodInClass(class *Class, name, descriptor string) *Method {
	for c := class; c != nil; c = class.superClass {
		for _, method := range class.methods {
			if method.name == name && method.descriptor == descriptor {
				return method
			}
		}
	}
	return nil
}

//todo java8之前,接口方法都是抽象方法,找到的方法体是空的，这样找到的方法有用吗？
func LookUpMethodInInterfaces(ifaces []*Class, name, descriptor string) *Method {
	for _, iface := range ifaces {
		for _, method := range iface.methods {
			if method.name == name && method.descriptor == descriptor {
				return method
			}
		}
		method := LookUpMethodInInterfaces(iface.interfaces, name, descriptor)
		if method != nil {
			return method
		}
	}
	return nil
}
