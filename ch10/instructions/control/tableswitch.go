package control

/**

https://docs.oracle.com/javase/specs/jvms/se8/html/jvms-3.html#jvms-3.10
*/
import (
	"../../rtda"
	"../base"
)

type TABLE_SWITCH struct {
	defaultOffset int32
	low           int32
	high          int32
	jumpOffsets   []int32
}

/**
https://www.cnblogs.com/heart-king/p/5390246.html

0: iload_1
 1: tableswitch   {		//tableswitch字节码后面有0-3字节的padding，原因如下：
         default: 42	//default在字节码中的地址，永远是4的倍数
             min: 0
             max: 4
               0: 36
               1: 38
               2: 42	//default
               3: 42	//default
               4: 40
    }
36: iconst_3
37: ireturn
38: iconst_2
39: ireturn
40: iconst_1
41: ireturn
42: iconst_m1
43: ireturn
*/

func (self *TABLE_SWITCH) FetchOperands(reader *base.ByteCodeReader) {
	reader.SkipPadding()
	self.defaultOffset = reader.ReadInt32()
	self.low = reader.ReadInt32()
	self.high = reader.ReadInt32()
	jumpOffsetsCount := self.high - self.low + 1
	self.jumpOffsets = reader.ReadInt32s(jumpOffsetsCount)
}

func (self *TABLE_SWITCH) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	index := stack.PopInt()
	var offset int
	if index >= self.low && index <= self.high {
		offset = int(self.jumpOffsets[index-self.low])
	} else {
		offset = int(self.defaultOffset)
	}
	base.Branch(frame, offset)
}
