package references

import (
	"../../rtda"
	"../../rtda/heap"
	"../base"
)

// Invoke a class (static) method
type INVOKE_STATIC struct{ base.Index16Instruction }

func (is *INVOKE_STATIC) Execute(frame *rtda.Frame) {
	cp := frame.Method().Class().ConstantPool()
	methodRef := cp.GetConstant(is.Index).(*heap.MethodRef)
	resolvedMethod := methodRef.ResolvedMethod()
	if !resolvedMethod.IsStatic() {
		panic("java.lang.IncompatibleClassChangeError")
	}
	class := resolvedMethod.Class()
	if !class.InitSatated() {
		frame.RevertNextPC()
		base.InitClass(frame.Thread(), class)
		return
	}
	base.InvokeMethod(frame, resolvedMethod)
}
