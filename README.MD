# 如何实现解析class文件

## 首先摆出class文件的格式
```

ClassFile {
    u4             magic; //Class 文件的标志
    u2             minor_version;//Class 的小版本号
    u2             major_version;//Class 的大版本号
    u2             constant_pool_count;//常量池的数量
    cp_info        constant_pool[constant_pool_count-1];//常量池
    u2             access_flags;//Class 的访问标记
    u2             this_class;//当前类
    u2             super_class;//父类
    u2             interfaces_count;//接口
    u2             interfaces[interfaces_count];//一个类可以实现多个接口
    u2             fields_count;//Class 文件的字段属性
    field_info     fields[fields_count];//一个类会可以有个字段
    u2             methods_count;//Class 文件的方法数量
    method_info    methods[methods_count];//一个类可以有个多个方法
    u2             attributes_count;//此类的属性表中的属性数
    attribute_info attributes[attributes_count];//属性表集合
}
```
## 解析第一步
   将class文件转换为byte数组。
## 非集合的字段有哪些？
   magic，minor_version, major_version, access_flags, this_class, super_class, interfaces_count, fields_count, methods_count, attributes_count
### 非集合字段长度如何计算？
答 ：非集合字段都是定长的，所以只需要依次读取相应字节长度的字段即可

在以上的字段中：**interfaces_count, fields_count, methods_count, attributes_count** 是用来指定其他集合字段中的元素个数的。

## 集合字段constant_pool
对于集合字段constant_pool中的常量，Java虚拟机规范给出了统一的常量结构。
### 常量字段的规律

1. **两类常量：** 字面量（literal)和符号引用（symbolic reference），其中字面量包括数字常量和字符串常量。
2. **数字常量结构（除去tag部分）可以根据长度分为2类：** “u2”, "u4"
3. **结构可以根据是否具有索引分为2类：** 数字常量和CONSTANT_utf8_info不包含索引，除此之外的其他常量都通过索引。**有人看到这里会有疑问：有没有索引和长度有关系吗？还真有，索引的长度都固定是u2的, 而且每个具有索引的结构会包含1-2个索引(u2或u4)。** 
4. **总结：** u2或u4

### 解析常量的方法已经很明显了

1. **设计常量结构:** 用无符号16位(u2)或无符号32位整数(u4)保存结构体中的字段，并为其每种结构实现相应的字段读取方法。
2. **遍历循环constant_pool_count次：** 每次循环只需要做三件事。
* 第一，读取tag字段。
* 第二，根据tag字段，创建相应的常量。
* 第三，调用常量的字段读取方法。

### 这一段内容和本文无关，纯粹给自己提个醒

1. **表头给的常量池大小比实际大1。** 常量池大小是constant_pool_count - 1
2. 有效的常量池索引是1~n-1，**0是无效索引**，代表不指向任何常量。
3. CONSTANT_Long_info和CONSTANT_Double_info各占两个位置，如果常量池中**存在这种占两个位置的常量，将会有更多的无效索引**，实际的常量数量会比n-1还少。

## 集合字段fields和methods
**额，为什么把这两个集合字段放在一起？** 原因：这两个集合的元素分别是field_info和method_info，而这两个结构的 **！！！不同之处只是结构名！！！**
```
field_info {
    u2             access_flags;
    u2             name_index;
    u2             descriptor_index;
    u2             attributes_count;
    attribute_info attributes[attributes_count];
}
method_info {
    u2             access_flags;
    u2             name_index;
    u2             descriptor_index;
    u2             attributes_count;
    attribute_info attributes[attributes_count];
}
```
对，就是一模一样（我不是故意在浪费篇幅）。

### 解析方法
**为了更好的理解，我必须先介绍attribute_info，因为它和cp_info不一样。**
（JVM内存区**资料讲了半天字节码怎么执行，也没告诉你字节码在哪儿（好气哦）**，我告诉你：字节码就在一个叫Code的attribute_info里存放着以字节数组的形式存放。）

## attribute_info
**这不是一个具体的结构**，因为各种属性表达的属性不同。除了JVMS(Java SE 8)中预定义的23项属性以外，不同的虚拟机还可以定义自己的属性类型。
### 不具体，但算是半具体结构
```
attribute_info {
    u2 attribute_name_index;
    u4 attribute_length;
    u1 info[attribute_length];
}
```
引用**官方解释**：
>The value of the attribute_length item indicates the length of the subsequent information in bytes. The length does not include the initial six bytes that contain the attribute_name_index and attribute_length items.

**人工翻译：**
attribute_length表明后续信息占用的字节数，它不包括attribute_name_index和attribute_length所占用的6个字节。

### attribute_info值得注意的地方

1. 属性名不允许重复
2. 没了。。。。

### 解析attribute_info

1. **读属性名索引**。读取attribute_name_index字段
2. **取属性名**。根据attribute_name_index从常量池中取到CONSTANT_utf8_info中的字节数组，反编码成字符串（实际是MUTF-8编码）
3. 根据属性名，**创建相应的属性**。
4. 属性调用自己的读取字段方法。

## 回过头来，解析attributes
**属性字段应用在3个地方：**，属性、方法、类。
**3类attributes，有同一个解析方法：**

1. 读取attributes_count
2. 循环读取attributes_count次，每次都执行前文提到的解析attribute_info的操作