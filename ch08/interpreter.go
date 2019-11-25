package main

import (
	"./instructions"
	"./instructions/base"
	"./rtda"
	"./rtda/heap"
	"fmt"
)

func interpret(method *heap.Method, logInst bool, args []string) {
	thread := rtda.NewThread()
	frame := thread.NewFrame(method)
	thread.PushFrame(frame)
	jArgs := createArgsArray(method.Class().Loader(), args)
	frame.LocalVars().SetRef(0, jArgs)
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

func createArgsArray(loader *heap.ClassLoader, args []string) *heap.Object {
	stringClass := loader.LoadClass("java/lang/String")
	argsArr := stringClass.ArrayClass().NewArray(uint(len(args)))
	jArgs := argsArr.Refs()
	for i, arg := range args {
		jArgs[i] = heap.JString(loader, arg)
	}
	return argsArr
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
