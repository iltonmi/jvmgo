package classpath

import (
	"os"
	"path/filepath"
)

//存放三种类路径
type Classpath struct {
	bootClasspath Entry
	extClasspath  Entry
	userClasspath Entry
}

/**
使用-Xjre选项解析启动类路径和扩展类路径
使用-classpath/-cp选项解析用户类路径
*/
func Parse(jreOption, cpOption string) *Classpath {
	cp := &Classpath{}
	cp.parseBootAndExtClasspath(jreOption)
	cp.parseUserClasspath(cpOption)
	return cp
}

func (self *Classpath) parseBootAndExtClasspath(jreOption string) {
	jreDir := getJreDir(jreOption)
	//jre/lib/*
	jreLibPath := filepath.Join(jreDir, "lib", "*")
	self.bootClasspath = newWildcardEntry(jreLibPath)
	//jre/lib/ext/*
	jreExtPath := filepath.Join(jreDir, "lib", "ext", "*")
	self.extClasspath = newWildcardEntry(jreExtPath)
}

/**
优先使用-Xjre选项作为jre目录。若无此选项，则在当前目录下寻找jre目录。
如果找不到，尝试使用JAVA_HOME环境变量
*/
func getJreDir(jreOption string) string {
	if jreOption != "" && exists(jreOption) {
		return jreOption
	}
	if exists("./jre") {
		return "./jre"
	}
	if jh := os.Getenv("JAVA_HOME"); jh != "" {
		return filepath.Join(jh, "jre")
	}
	panic("Can not find jre folder!")
}

func exists(path string) bool {
	//获取文件结构
	if _, err := os.Stat(path); err != nil {
		//若错误是因为文件或目录不存在，则返回false
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

/**
若用户没有提供-classpath或-cp选项，则使用当前目录作为用户类路径
*/
func (self *Classpath) parseUserClasspath(cpOption string) {
	if cpOption == "" {
		cpOption = "."
	}
	self.userClasspath = newEntry(cpOption)
}

/**
依次从启动类路径、扩展类路径和用户类路径中搜索class文件
*/
func (self *Classpath) ReadClass(className string) ([]byte, Entry, error) {
	className = className + ".class"
	if data, entry, err := self.bootClasspath.readClass(className); err == nil {
		return data, entry, err
	}
	if data, entry, err := self.extClasspath.readClass(className); err == nil {
		return data, entry, err
	}
	return self.userClasspath.readClass(className)
}

func (self *Classpath) String() string {
	return self.userClasspath.String()
}
