package classfile

type MarkerAttribute struct{}

func (self *MarkerAttribute) readInfo(reader *ClassReader) {

}

//这两种结构仅仅起标记作用，不包含任何数据
/*
Deprecated_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
}
*/
type DeprecatedAttribute struct {
	MarkerAttribute
}

/*
标记源文件中不存在、由编译器生成的类成员，
主要是为了支持嵌套类和嵌套接口
Synthetic_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
}
*/
type SyntheticAttribute struct {
	MarkerAttribute
}
