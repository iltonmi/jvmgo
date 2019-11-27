package constants

import "../base"
import "../../rtda"

// Do nothing
type NOP struct{ base.NoOperandsInstruction }

func (instr *NOP) Execute(frame *rtda.Frame) {
	// really do nothing
}
