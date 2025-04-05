package test

import (
	"fmt"
	"strconv"
	"testing"
)

/*
strconv 基本数据类型和字符串处理
常用函数： Atoi()、Itoa()、parse系列、format系列、append系列。
a的典故：这是C语言遗留下的典故。C语言中没有string类型而是用字符数组(array)表示字符串，所以Itoa对很多C系的程序员很好理解。
*/
func TestStrconv(t *testing.T) {
	//常用
	stringAndIntegerSwapping()     //Atoi()【string转int】，Itoa()【int转string】
	parseBoolOrIntOrUnitOrFloat()  //Parse系列函数，将string转为指定类型
	formatBoolOrIntOrUnitOrFloat() //Format系列函数，将指定类型转为string
	appendBoolOrIntOrUnitOrFloat() //Append系列函数，用于高效地将各种类型的数据转换为字符串后，追加到现有的字节切片（[]byte）中。
	/*
		其他
		isPrint()
		func IsPrint(r rune) bool
		返回一个字符是否是可打印的，和unicode.IsPrint一样，r必须是：字母（广义）、数字、标点、符号、ASCII空格。

		CanBackquote()
		func CanBackquote(s string) bool
		返回字符串s是否可以不被修改的表示为一个单行的、没有空格和tab之外控制字符的反引号字符串。

		除上面的函数外，strconv包中还有Append系列、Quote系列等函数。具体用法可查看官方文档。
	*/
}

// string与int类型相互转换
func stringAndIntegerSwapping() {
	//字符串转int
	s1 := "100"
	i1, _ := strconv.Atoi(s1)
	fmt.Println("字符串转int成功i：", i1)
	//int转字符串
	i2 := 200
	s2 := strconv.Itoa(i2)
	fmt.Println("int转字符串成功string：", s2)
}

/*
Parse系列函数
用于转换字符串为给定类型的值：ParseBool()、ParseFloat()、ParseInt()、ParseUint()。

ParseBool()
func ParseBool(str string) (value bool, err error)
返回字符串表示的bool值。它接受1、0、t、f、T、F、true、false、True、False、TRUE、FALSE；否则返回错误。

ParseInt()
func ParseInt(s string, base int, bitSize int) (i int64, err error)
返回字符串表示的整数值，接受正负号。
base指定进制（2到36），如果base为0，则会从字符串前置判断，“0x"是16进制，“0"是8进制，否则是10进制；
bitSize指定结果必须能无溢出赋值的整数类型，0、8、16、32、64 分别代表 int、int8、int16、int32、int64；
返回的err是*NumErr类型的，如果语法有误，err.Error = ErrSyntax；如果结果超出类型范围err.Error = ErrRange。

ParseUnit()
func ParseUint(s string, base int, bitSize int) (n uint64, err error)
ParseUint类似ParseInt但不接受正负号，用于无符号整型。

ParseFloat()
func ParseFloat(s string, bitSize int) (f float64, err error)
解析一个表示浮点数的字符串并返回其值。
如果s合乎语法规则，函数会返回最为接近s表示值的一个浮点数（使用IEEE754规范舍入）。
bitSize指定了期望的接收类型，32是float32（返回值可以不改变精确值的赋值给float32），64是float64；
返回值err是*NumErr类型的，语法有误的，err.Error=ErrSyntax；结果超出表示范围的，返回值f为±Inf，err.Error= ErrRange。
*/
func parseBoolOrIntOrUnitOrFloat() {
	fmt.Println("Parse系列函数")
	//些函数都有两个返回值，第一个返回值是转换后的值，第二个返回值为转化失败的错误信息。
	strBool := "true"
	b, _ := strconv.ParseBool(strBool)
	fmt.Printf("%s字符串转bool:%T\n", strBool, b)

	strFloat := "3.1415926"
	f, _ := strconv.ParseFloat(strFloat, 64)
	fmt.Printf("%s字符串转float:%T\n", strFloat, f)

	strInt := "-2"
	i, _ := strconv.ParseInt(strInt, 10, 64)
	fmt.Printf("%s字符串转int:%T\n", strInt, i)

	strUint := "2"
	u, _ := strconv.ParseUint(strUint, 10, 64)
	fmt.Printf("%s字符串转uint:%T\n", strUint, u)
}

/*
Format系列函数，实现将给定类型数据格式化为string类型。

FormatBool()
func FormatBool(b bool) string
根据b的值返回"true"或"false”。

FormatInt()
func FormatInt(i int64, base int) string
返回i的base进制的字符串表示。base 必须在2到36之间，结果中会使用小写字母’a’到’z’表示大于10的数字。

FormatUint()
func FormatUint(i uint64, base int) string
是FormatInt的无符号整数版本。

FormatFloat()
func FormatFloat(f float64, fmt byte, prec, bitSize int) string
函数将浮点数表示为字符串并返回。
bitSize表示f的来源类型（32：float32、64：float64），会据此进行舍入。
fmt表示格式：‘f’（-ddd.dddd）、‘b’（-ddddp±ddd，指数为二进制）、’e’（-d.dddde±dd，十进制指数）；
‘E’（-d.ddddE±dd，十进制指数）、‘g’（指数很大时用’e’格式，否则’f’格式）、‘G’（指数很大时用’E’格式，否则’f’格式）。
prec控制精度（排除指数部分）：对’f’、’e’、‘E’，它表示小数点后的数字个数；对’g’、‘G’，它控制总的数字个数。
如果prec 为-1，则代表使用最少数量的、但又必需的数字来表示f。
*/
func formatBoolOrIntOrUnitOrFloat() {
	fmt.Println("Format系列函数")
	b := true
	s1 := strconv.FormatBool(b)
	fmt.Printf("%v %T转为string:%v\n", b, b, s1)

	f := 3.1415
	s2 := strconv.FormatFloat(f, 'E', -1, 64)
	fmt.Printf("%v %T转为string:%v\n", f, f, s2)

	var i int64 = -2
	s3 := strconv.FormatInt(i, 16)
	fmt.Printf("%v %T转为string:%v\n", i, i, s3)

	var u uint64 = 2
	s4 := strconv.FormatUint(u, 16)
	fmt.Printf("%v %T转为string:%v\n", u, u, s4)

}

/*
Append 系列函数用于高效地将各种类型的数据转换为字符串后，追加到现有的字节切片（[]byte）中。
这些函数通常用于高性能场景（如日志记录、网络协议处理），避免频繁分配内存。

函数																				功能
func AppendBool(dst []byte, b bool) []byte										追加布尔值
func AppendInt(dst []byte, i int64, base int) []byte							追加整数（可指定进制）
func AppendUint(dst []byte, i uint64, base int) []byte							追加无符号整数
func AppendFloat(dst []byte, f float64, fmt byte, prec, bitSize int) []byte		追加浮点数
func AppendQuote(dst []byte, s string) []byte									追加带双引号的字符串（转义特殊字符）
func AppendQuoteRune(dst []byte, r rune) []byte									追加带单引号的 Unicode 字符

核心特点
1.零分配（Zero-allocation）直接操作传入的 []byte，避免 fmt.Sprintf 或 strconv.Itoa 的临时内存分配。
2.链式调用，返回新的切片，可连续调用多个 Append 函数。
3.性能优先，比 fmt.Sprintf 快 2-5 倍。
*/
func appendBoolOrIntOrUnitOrFloat() {
	//1.基本类型追加
	fmt.Println("Append系列函数")
	buf := make([]byte, 0, 32) // 预分配容量

	// 追加布尔值
	buf = strconv.AppendBool(buf, true)
	buf = append(buf, '|') // 添加分隔符

	// 追加十进制整数
	buf = strconv.AppendInt(buf, -42, 10)
	buf = append(buf, '|')

	// 追加十六进制整数
	buf = strconv.AppendUint(buf, 42, 16)
	buf = append(buf, '|')

	// 追加浮点数（保留2位小数）
	buf = strconv.AppendFloat(buf, 3.14159, 'f', 2, 64)

	fmt.Println("1.基本类型追加:", string(buf)) // 输出: true|-42|2a|3.14

	//2.字符串和字符转义
	bufString := make([]byte, 0, 50)
	bufString = append(bufString, "Log: "...)

	// 追加带引号的字符串（自动转义）
	bufString = strconv.AppendQuote(bufString, "Hello\tWorld!")
	bufString = append(bufString, '\n')

	// 追加 Unicode 字符
	bufString = strconv.AppendQuoteRune(bufString, '☺')

	fmt.Println("2.字符串和字符转义", string(bufString))

	//3.高性能日志拼接
	bytes := logEvent(123, "login", true)
	fmt.Println("3.高性能日志拼接:", string(bytes))
}

// strconv.Append 高性能日志拼接
func logEvent(userID int, action string, success bool) []byte {
	buf := make([]byte, 0, 128)
	buf = append(buf, "user="...)
	buf = strconv.AppendInt(buf, int64(userID), 10)
	buf = append(buf, ", action="...)
	buf = strconv.AppendQuote(buf, action)
	buf = append(buf, ", success="...)
	buf = strconv.AppendBool(buf, success)
	return buf
}
