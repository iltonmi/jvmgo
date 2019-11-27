package misc

import (
	"../.."
	"../../../instructions/base"
	"../../../rtda"
)

func init() {
	native.Register("sun/misc/VM", "initialize", "()V", initialize)
}

// private static native void initialize();
func initialize(frame *rtda.Frame) { // hack: just make VM.savedProps nonempty
	classLoader := frame.Method().Class().Loader()
	jlSysClass := classLoader.LoadClass("java/lang/System")
	initSysClass := jlSysClass.GetStaticMethod("initializeSystemClass", "()V")
	base.InvokeMethod(frame, initSysClass)
}
