package test

import (
	"flag"
	"fmt"
	"os"
	"testing"
	"time"
)

/*
flag【要用main函数，否则无法使用命令行输入参数】
1、实现命令行参数的解析
2、开发命令行工具简单
*/
func TestFlag(t *testing.T) {
	getArgs() //获取命令行参数
	/*
		flag参数类型
		flag包支持的命令行参数类型有bool、int、int64、uint、uint64、float float64、string、duration。

		flag参数			有效值
		字符串flag		合法字符串
		整数flag			1234、0664、0x1234等类型，也可以是负数。
		浮点数flag		合法浮点数
		bool类型flag		1, 0, t, f, T, F, true, false, TRUE, FALSE, True, False。
		时间段flag		任何合法的时间段字符串。如"300ms"、"-1.5h"、“2h45m”。合法的单位有"ns"、“us” /“µs”、“ms”、“s”、“m”、“h”。

		定义命令行参数：
		flag.Type：flag.Type(flag名, 默认值, 帮助信息)*Type
		flag.TypeVar()：flag.TypeVar(Type指针, flag名, 默认值, 帮助信息)

		flag.Parse()
		通过以上两种方法定义好命令行flag参数后，需要通过调用flag.Parse()来对命令行参数进行解析。
		支持的命令行参数格式有以下几种：
			-flag xxx （使用空格，一个-符号）
			--flag xxx （使用空格，两个-符号）
			-flag=xxx （使用等号，一个-符号）
			--flag=xxx （使用等号，两个-符号）
		其中，布尔类型的参数必须使用等号的方式指定。
		Flag解析在第一个非flag参数（单个"-“不是flag参数）之前停止，或者在终止符”–“之后停止。

		flag其他函数
		flag.Args()  ////返回命令行参数后的其他参数，以[]string类型
		flag.NArg()  //返回命令行参数后的其他参数个数
		flag.NFlag() //返回使用的命令行参数个数
	*/
	//定义命令行参数方式1
	var name string
	var age int
	var married bool
	var delay time.Duration
	flag.StringVar(&name, "name", "张三", "姓名")
	flag.IntVar(&age, "age", 18, "年龄")
	flag.BoolVar(&married, "married", false, "婚否")
	flag.DurationVar(&delay, "d", 0, "延迟的时间间隔")

	//解析命令行参数
	flag.Parse()
	fmt.Println(name, age, married, delay)
	//返回命令行参数后的其他参数
	fmt.Println(flag.Args())
	//返回命令行参数后的其他参数个数
	fmt.Println(flag.NArg())
	//返回使用的命令行参数个数
	fmt.Println(flag.NFlag())

}

// 获取命令行参数
func getArgs() {
	//os.Args是一个[]string。os.Args是一个存储命令行参数的字符串切片，它的第一个元素是执行文件的名称。
	if len(os.Args) > 0 {
		for index, arg := range os.Args {
			fmt.Printf("args[%d]=%v\n", index, arg)
		}
	}
}
