package main

import (
	"./instructions"
	"./instructions/base"
	"./rtda"
	"./rtda/heap"
	"fmt"
)

func interpret(method *heap.Method) {
	thread := rtda.NewThread()
	frame := thread.NewFrame(method)
	thread.PushFrame(frame)
	defer catchErr(frame)
	loop(thread, method.Code())
}

//打印局部变量表和操作数栈的内容
func loop(thread *rtda.Thread, bytecode []byte) {
	frame := thread.PopFrame()
	reader := &base.ByteCodeReader{}
	for {
		pc := frame.NextPC()
		thread.SetPC(pc)
		//decode
		reader.Reset(bytecode, pc)
		opcode := reader.ReadUInt8()
		inst := instructions.NewInstruction(opcode)
		inst.FetchOperands(reader)
		frame.SetNextPC(reader.PC())
		//execute
		//fmt.Printf("operand stack before execute: %v\n", frame.OperandStack())
		fmt.Printf("pc:%2d inst:%T %v\n", pc, inst, inst)
		inst.Execute(frame)
	}
}

func catchErr(frame *rtda.Frame) {
	if r := recover(); r != nil {
		fmt.Printf("LocalVars:%v\n", frame.LocalVars())
		fmt.Printf("OperandStack:%v\n", frame.OperandStack())
		panic(r)
	}
}
