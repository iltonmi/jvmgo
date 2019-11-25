package heap

import (
	"unicode/utf16"
)

//字符串池,key是Go字符串,value是Java字符串
var internedStrings = map[string]*Object{}

/*
java.lang.String的构造函数,直接用hack方式创建实例
*/
func JString(loader *ClassLoader, goStr string) *Object {
	//若字符串已在池中,直接返回,
	if internedStr, ok := internedStrings[goStr]; ok {
		return internedStr
	}
	chars := stringToUtf16(goStr)
	//utf16格式字符数组
	jChars := &Object{loader.LoadClass("[C"), chars}
	//创建java字符串实例
	jStr := loader.LoadClass("java/lang/String").NewObject()
	//将java.lang.String的value字段设置为字符数组
	jStr.SetRefVar("value", "[C", jChars)
	internedStrings[goStr] = jStr
	return jStr
}

func stringToUtf16(s string) []uint16 {
	//go字符串在内存中是utf8编码, 强转为utf32
	runes := []rune(s) //utf32
	//调用函数编码成utf16
	return utf16.Encode(runes)
}

func GoString(jStr *Object) string {
	charArr := jStr.GetRefVar("value", "[C")
	return utf16ToString(charArr.Chars())
}

func utf16ToString(s []uint16) string {
	runes := utf16.Decode(s) //先转换utf8
	return string(runes)
}
