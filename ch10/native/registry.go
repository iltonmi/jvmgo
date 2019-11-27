package native

import "../rtda"

type NativeMethod func(frame *rtda.Frame)

var registry = map[string]NativeMethod{}

/**
查找本地方法实现
*/
func Register(className, methodName, methodDesctiptor string, method NativeMethod) {
	//类名、方法名和方法描述名组合在一起唯一确定一个方法
	key := className + "~" + methodName + "~" + methodDesctiptor
	registry[key] = method
}

func FindNativeMethod(className, methodName, methodDescriptor string) NativeMethod {
	key := className + "~" + methodName + "~" + methodDescriptor
	if method, ok := registry[key]; ok {
		return method
	}
	if methodDescriptor == "()V" && methodName == "registerNatives" {
		return emptyNativeMethod
	}
	return nil
}

func emptyNativeMethod(frame *rtda.Frame) {
	//do nothing
}
