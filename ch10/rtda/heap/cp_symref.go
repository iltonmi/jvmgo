package heap

//symbolic reference
/**
对于类符号引用，只要有类名就可以解析符号引用
对于字段和方法的符号引用，需要同伙类符号引用得到类数据，再用名和描述符查找
*/
type SymRef struct {
	cp *ConstantPool
	//全限定名
	className string
	//缓存类结构体指针，类符号引用只需解析一次，后续可以直接使用缓存值
	class *Class
}

func (sr *SymRef) ResolvedClass() *Class {
	if sr.class == nil {
		sr.resolveClassRef()
	}
	return sr.class
}

func (sr *SymRef) resolveClassRef() {
	d := sr.cp.class
	c := d.loader.LoadClass(sr.className)
	if !c.isAccessibleTo(d) {
		panic("java.lang.IllegalAccessError")
	}
	sr.class = c
}
