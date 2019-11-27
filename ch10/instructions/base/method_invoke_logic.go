package base

import (
	"../../rtda"
	"../../rtda/heap"
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
}
