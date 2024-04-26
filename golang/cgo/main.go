package main

/*
#cgo CFLAGS: -I.
#cgo LDFLAGS: -L.
#cgo windows CFLAGS: -DCGO_OS_WINDOWS=1
#cgo darwin CFLAGS: -DCGO_OS_DARWIN=1
#cgo linux CFLAGS: -DCGO_OS_LINUX=1
#include "test_lib.h"

#if defined(CGO_OS_WINDOWS)
    const char* os = "windows";
#elif defined(CGO_OS_DARWIN)
    const char* os = "darwin";
#elif defined(CGO_OS_LINUX)
    const char* os = "linux";
#else
#	error(unknown os)
#endif
*/
import "C"
import "fmt"

// 紧跟在 import "C" 前的注释是特殊语法，包含正常的C代码（一段.c，不只能用于引入宏和头文件）
// 另外类型系统比较诡异，比如这里的C.CString 类型其实展开后是 main.C.CString，因此从不同 package 导过来
// 但同样是 CString 的类型是不兼容的

// #cgo 设置编译阶段和链接阶段的相关参数
// 在库文件的检索目录中可以通过 ${SRCDIR} 变量表示当前包目录的绝对路径：#cgo LDFLAGS: -L/go/src/foo/libs -lfoo
// 支持条件选择

// 对于一个启用 CGO 特性的程序，CGO 会构造一个虚拟的 C 包。通过这个虚拟的 C 包可以调用 C 语言函数。

func main() {
	// success
	var name string = "Bocabbage"
	C.hello(C.CString(name))
	C.hello(C.os)

	// failed
	_, err := C.hello(C.CString("")) // 没有返回值也还是要取第一个 v
	if err != nil {
		fmt.Printf("Error: %v", err)
	}
}
