package base

import (
	"../../rtda"
	"../../rtda/heap"
	"fmt"
)

func InvokeMethod(invokerFrame *rtda.Frame, method *heap.Method) {
	thread := invokerFrame.Thread()
	newFrame := thread.NewFrame(method)
	//将新创建的帧推入Java虚拟机栈
	thread.PushFrame(newFrame)
	//传递参数
	argsSlotCount := int(method.ArgSlotCount())
	if argsSlotCount > 0 {
		for i := argsSlotCount - 1; i >= 0; i-- {
			slot := invokerFrame.OperandStack().PopSlot()
			newFrame.LocalVars().SetSlot(uint(i), slot)
		}
	}
	if method.IsNative() {
		if method.Name() == "registerNatives" {
			thread.PopFrame()
		} else if method.Name() == "currentTimeMillis" {
			thread.PopFrame()
			thread.PopFrame()
			thread.PopFrame()
		} else {
			panic(fmt.Sprintf("native method: %v.%v%v\n",
				method.Class().Name(), method.Name(), method.Descriptor()))
		}
	}
}
