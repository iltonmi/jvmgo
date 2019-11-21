package classfile

/*
LineNumberTable 属性 表 存放 方法 的 行号 信息， LocalVariableTable 属性 表中 存放 方法 的 局部 变量 信息。
这 两种 属性 和 前面 介绍 的 SourceFile 属性 都 属于 调试 信息， 都不 是 运行时 必需 的。
在 使用 javac 编译器 编译 Java 程序 时， 默认 会在 class 文件 中 生成 这些 信息。
可以 使用 javac 提供 的- g： none 选项 来 关闭 这些 信息 的 生成，
LineNumberTable_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
    u2 line_number_table_length;
    {   u2 start_pc;
        u2 line_number;
    } line_number_table[line_number_table_length];
}
*/
type LineNumberTableAttribute struct {
	lineNumberTable []*LineNumberTableEntry
}

type LineNumberTableEntry struct {
	startPc    uint16
	lineNumber uint16
}

func (self *LineNumberTableAttribute) readInfo(reader *ClassReader) {
	lineNumberTableLength := reader.readUint16()
	self.lineNumberTable = make([]*LineNumberTableEntry, lineNumberTableLength)
	for i := range self.lineNumberTable {
		self.lineNumberTable[i] = &LineNumberTableEntry{
			startPc:    reader.readUint16(),
			lineNumber: reader.readUint16(),
		}
	}
}
