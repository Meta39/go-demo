package main

import (
	"errors"
	"fmt"
	"golang/base/customerror"
	"golang/base/structs"
	"os"
	"strings"
)

/*
Error Go 语言中的错误处理与其他语言不太一样，它把错误当成一种值来处理，更强调判断错误、处理错误，而不是一股脑的 catch 捕获异常。
*/
func Error() {
	fmt.Println("============ 错误 ============")
	/*
		Error接口和错误处理：
		Error 接口
		Go 语言中把错误当成一种特殊的值来处理，不支持其他语言中使用try/catch捕获异常的方式。
		Go 语言中使用一个名为 error 接口来表示错误类型。

		type error interface {
			Error() string
		}

		error 接口只包含一个方法——Error，这个函数需要返回一个描述错误信息的字符串。
		当一个函数或方法需要返回错误时，我们通常是把错误作为最后一个返回值。例如下面标准库 os 中打开文件的函数。

		func Open(name string) (*File, error) {
			return OpenFile(name, O_RDONLY, 0)
		}
		由于 error 是一个接口类型，默认零值为nil。所以我们通常将调用函数返回的错误与nil进行比较，以此来判断函数是否返回错误。例如你会经常看到类似下面的错误判断代码。
		当我们使用fmt包打印错误时会自动调用 error 类型的 Error 方法，也就是会打印出错误的描述信息。
	*/
	//获取当前工作目录（调试用）
	wd, _ := os.Getwd()
	fmt.Println("获取当前工作目录（调试用）:", wd) // 检查路径是否匹配
	file, err := os.Open("./base/func.go")
	if err != nil {
		fmt.Println("打开文件失败,customerror:", err)
		return
	}
	// 确保函数退出时关闭文件，防止资源泄漏（如文件描述符耗尽、数据未刷新等问题）
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println("关闭文件失败", err)
		}
	}(file)

	/*
		创建错误
		我们可以根据需求自定义 error，最简单的方式是使用errors 包提供的New函数创建一个错误。
		customerror.New
		函数签名如下，
		func New(text string) error
		它接收一个字符串参数返回包含该字符串的错误。我们可以在函数返回时快速创建一个错误。
	*/
	fmt.Println("创建错误")
	name, err2 := queryByName("name")
	if err2 != nil {
		fmt.Println(err2)
		return
	}
	fmt.Println(name)

	//当我们需要传入格式化的错误描述信息时，使用fmt.Errorf是个更好的选择。
	_ = fmt.Errorf("查询数据库失败，customerror:%v", err2)

	/*
		但是上面的方式会丢失原有的错误类型，只拿到错误描述的文本信息。
		为了不丢失函数调用的错误链，使用fmt.Errorf时搭配使用特殊的格式化动词%w，可以实现基于已有的错误再包装得到一个新的错误。
	*/
	_ = fmt.Errorf("查询数据库失败，customerror:%w", err2)

	/*
		对于这种二次包装的错误，errors包中提供了以下三个方法。
		func Unwrap(customerror error) error                 // 获得err包含下一层错误
		func Is(customerror, target error) bool              // 判断err是否包含target
		func As(customerror error, target interface{}) bool  // 判断err是否为target类型
	*/

	/*
		错误结构体类型（PS：自定义错误，类似Java自定义异常）
		此外我们还可以自己定义结构体类型，实现error接口。
	*/
	err3 := customerror.New("自定义运行时异常")
	if err3 != nil {
		fmt.Println(err3)
		//return//一般情况下是不让程序往下执行的
	}

	fmt.Println("============ 错误 ============")
}

// customerror.New快速创建一个错误
func queryByName(name string) (p *structs.Person, err error) {
	if strings.TrimSpace(name) == "" {
		return nil, errors.New("无效的name")
	}
	return p, nil
}
