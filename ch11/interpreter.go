package main

import (
	"./instructions"
	"./instructions/base"
	"./rtda"
	"fmt"
)

func interpret(thread *rtda.Thread, logInst bool) {
	defer catchErr(thread)
	loop(thread, logInst)
}

//打印局部变量表和操作数栈的内容
func loop(thread *rtda.Thread, logInst bool) {
	reader := &base.ByteCodeReader{}
	for {
		frame := thread.CurrentFrame()
		pc := frame.NextPC()
		thread.SetPC(pc)
		//decode
		reader.Reset(frame.Method().Code(), pc)
		opcode := reader.ReadUInt8()
		inst := instructions.NewInstruction(opcode)
		inst.FetchOperands(reader)
		frame.SetNextPC(reader.PC())
		if logInst {
			logInstruction(frame, inst)
		}
		//execute
		//fmt.Printf("current method:" + frame.Method().Name() + "\n")
		//fmt.Printf("local variable table before execute: %v\n", frame.LocalVars())
		//fmt.Printf("operand stack before execute: %v\n", frame.OperandStack())
		//fmt.Printf("pc:%2d inst:%T %v\n", pc, inst, inst)
		inst.Execute(frame)
		if thread.IsStackEmpty() {
			break
		}
	}
}

func catchErr(thread *rtda.Thread) {
	if r := recover(); r != nil {
		logFrames(thread)
		panic(r)
	}
}

func logFrames(thread *rtda.Thread) {
	for !thread.IsStackEmpty() {
		frame := thread.PopFrame()
		method := frame.Method()
		className := method.Class().Name()
		fmt.Printf(">> pc:%4d %v.%v%v \n",
			frame.NextPC(), className, method.Name(), method.Descriptor())
	}
}

func logInstruction(frame *rtda.Frame, inst base.Instruction) {
	method := frame.Method()
	className := method.Class().Name()
	methodName := method.Name()
	pc := frame.Thread().PC()
	fmt.Printf("%v.%v #%2d %T %v\n", className, methodName, pc, inst, inst)
}
