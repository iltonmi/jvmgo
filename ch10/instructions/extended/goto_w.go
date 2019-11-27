package extended

import (
	"../../rtda"
	"../base"
)

// Branch always (wide index)
type GOTO_W struct {
	offset int
}

func (self *GOTO_W) FetchOperands(reader *base.ByteCodeReader) {
	//和goto的唯一区别：索引从2字节变成4字节
	self.offset = int(reader.ReadInt32())
}
func (self *GOTO_W) Execute(frame *rtda.Frame) {
	base.Branch(frame, self.offset)
}
