package test

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
	"testing"
)

/*
fmt
fmt包实现了类似C语言printf和scanf的格式化I/O。主要分为向外输出内容和获取输入内容两大部分。
*/
func TestFmt(t *testing.T) {
	fmt.Println("============ fmt ============")
	outPrint()             //向外输出
	formattedPlaceholder() //格式化占位符
	getInput()             //获取输入【需要在main函数中执行，否则无法达到控制台输入的效果】

	fmt.Println("============ fmt ============")
}

// 向外输出
func outPrint() {
	fmt.Println("向外输出")
	/*
		Print
		func Print(a ...interface{}) (n int, err error) //将内容输出到系统的标准输出
		func Printf(format string, a ...interface{}) (n int, err error) //支持格式化输出字符串
		func Println(a ...interface{}) (n int, err error) //输出内容的结尾添加一个换行符
	*/
	fmt.Print("在终端打印该信息。")
	name := "Meta39"
	fmt.Printf("%s", name)
	fmt.Println("在终端打印单独一行显示")

	/*
		Fprint
		将内容输出到一个io.Writer接口类型的变量w中，我们通常用这个函数往文件中写入内容。
		func Fprint(w io.Writer, a ...interface{}) (n int, err error)
		func Fprintf(w io.Writer, format string, a ...interface{}) (n int, err error)
		func Fprintln(w io.Writer, a ...interface{}) (n int, err error)
		注意：只要满足io.Writer接口的类型都支持写入。
	*/
	// 向标准输出写入内容
	fmt.Fprintln(os.Stdout, "向标准输出写入内容")
	//打开 Fprint.txt 文件，如果文件不存在，则创建并授予读写权限。
	fileObj, err := os.OpenFile("../../Fprint.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	defer fileObj.Close() //关闭文件流
	if err != nil {
		fmt.Println("打开文件出错，err:", err)
		return
	}
	name2 := "Meta39"
	// 向打开的文件句柄中写入内容
	fmt.Fprintf(fileObj, "往文件中写如信息：%s", name2)

	/*
		Sprint
		把传入的数据生成并返回一个字符串。
		func Sprint(a ...interface{}) string
		func Sprintf(format string, a ...interface{}) string
		func Sprintln(a ...interface{}) string
	*/
	s1 := fmt.Sprint("Meta39")
	name3 := "Meta39"
	age := 18
	s2 := fmt.Sprintf("name:%s,age:%d", name3, age)
	s3 := fmt.Sprintln("Meta39")
	fmt.Println(s1, s2, s3) //Meta39 name:Meta39,age:18 Meta39

	/*
		Errorf
		根据format参数生成格式化字符串并返回一个包含该字符串的错误。
		func Errorf(format string, a ...interface{}) error
	*/
	e := errors.New("错误啦~")
	//%w占位符用来生成一个可以包裹Error的Wrapping Error。
	w := fmt.Errorf("errorf:%w", e)
	fmt.Println(w) //errorf:错误啦~
}

// 格式化占位符
func formattedPlaceholder() {
	/*
		格式化占位符
		占位符		说明

		【通用占位符】
		%v			值的默认格式表示
		%+v			类似%v，但输出结构体时会添加字段名
		%#v			值的Go语法表示
		%T			打印值的类型
		%%			百分号

		【布尔值】
		%t			布尔值，true或false

		【整型】
		%b			表示为二进制
		%c			该值对应的unicode码值
		%d			表示为十进制
		%o			表示为八进制
		%x			表示为十六进制，使用a-f
		%X			表示为十六进制，使用A-F
		%U			表示为Unicode格式：U+1234，等价于"U+%04X"
		%q			该值对应的单引号括起来的go语法字符字面值，必要时会采用安全的转义表示

		【浮点数与复数】
		%b			无小数部分、二进制指数的科学计数法，如-123456p-78
		%e			科学计数法，如-1234.456e+78
		%E			科学计数法，如-1234.456E+78
		%f			有小数部分但无指数部分，如123.456
		%F			等价于%f
		%g			根据实际情况采用%e或%f格式（以获得更简洁、准确的输出）
		%G			根据实际情况采用%E或%F格式（以获得更简洁、准确的输出）

		【字符串和[]byte】
		%s			直接输出字符串或者[]byte
		%q			该值对应的双引号括起来的go语法字符串字面值，必要时会采用安全的转义表示
		%x			每个字节用两字符十六进制数表示（使用a-f
		%X			每个字节用两字符十六进制数表示（使用A-F）

		【指针】
		%p			表示为十六进制，不带前导的0x
		%#p			表示为十六进制，并加上前导的0x

		【宽度标识符】
		%f			默认宽度，默认精度
		%9f			宽度9，默认精度
		%.2f		默认宽度，精度2
		%9.2f		宽度9，精度2
		%9.f		宽度9，精度0

		【其他flag】
		‘+’			总是输出数值的正负号；对%q（%+q）会生成全部是ASCII字符的输出（通过转义）；
		’ '			对数值，正数前加空格而负数前加负号；对字符串采用%x或%X时（% x或% X）会给各打印的字节之间加空格
		‘-’			在输出右边填充空白而不是默认的左边（即从默认的右对齐切换为左对齐）；
		‘#’			八进制数前加0（%#o），十六进制数前加0x（%#x）或0X（%#X），指针去掉前面的0x（%#p）对%q（%#q），对%U（%#U）会输出空格和单引号括起来的go字面值；
		‘0’			使用0而不是空格填充，对于数值类型会把填充的0放在正负号后面；
	*/
	fmt.Println("【通用占位符】")
	o := struct{ name string }{"Meta39"}           //匿名结构体
	fmt.Printf("%%v值的默认格式表示:%v\n", o)              //{Meta39}
	fmt.Printf("%%+v类似%%v，但输出结构体时会添加字段名:%+v\n", o) //{name:Meta39}
	fmt.Printf("%%#v值的Go语法表示:%#v\n", o)            //struct { name string }{name:"Meta39"}
	fmt.Printf("%%T打印值的类型:%T\n", o)                //struct { name string }
	fmt.Printf("%%%%百分号:100%%\n")                  //100%
	fmt.Printf("%%t布尔值:%t\n", false)               //false

	fmt.Println("【整型】")
	n := 111
	fmt.Printf("%%b二进制:%b\n", n)           //1101111
	fmt.Printf("%%c值对应的unicode码值:%c\n", n) //o
	fmt.Printf("%%d十进制:%d\n", n)           //111
	fmt.Printf("%%o八进制:%o\n", n)           //157
	fmt.Printf("%%x十六进制，使用a-f:%x\n", n)    //6f
	fmt.Printf("%%X十六进制，使用A-F:%X\n", n)    //6F

	fmt.Println("【浮点数与复数】")
	f := 3.141592653589793
	fmt.Printf("%%b无小数部分、二进制指数的科学计数法:%b\n", f) //7074237752028440p-51
	fmt.Printf("%%e科学计数法（小写字母）:%e\n", f)       //3.141593e+00
	fmt.Printf("%%E科学计数法（大写字母）:%E\n", f)       //3.141593E+00
	fmt.Printf("%%f有小数部分但无指数部分:%f\n", f)       //3.141593
	fmt.Printf("%%F等价于%%f有小数部分但无指数部分:%F\n", f) //3.141593
	fmt.Printf("%%g根据实际情况采用%%e或%%f格式:%g\n", f) //3.141592653589793
	fmt.Printf("%%G根据实际情况采用%%E或%%F格式:%G\n", f) //3.141592653589793

	fmt.Println("【字符串和[]byte】")
	s := "Meta39"
	b := []byte{'a', 1, 'A', 'X', 2}
	fmt.Printf("%%s直接输出字符串:%s 或者[]byte:%s\n", s, b)               //Meta39 aAX
	fmt.Printf("%%q该值对应的双引号括起来的go语法字符串字面值，必要时会采用安全的转义表示:%q\n", s) //"Meta39"
	fmt.Printf("%%x每个字节用两字符十六进制数表示（使用a-f）:%x\n", s)               //4d6574613339
	fmt.Printf("%%X每个字节用两字符十六进制数表示（使用A-F）:%X\n", s)               //4D6574613339

	fmt.Println("【指针】")
	p := 10
	fmt.Printf("%%p表示为十六进制，并加上前导的0x：【%p】\n", &p)  //0xc00000a458
	fmt.Printf("%%#p表示为十六进制，不带前导的0x：【%#p】\n", &p) //c00000a458

	fmt.Println("【宽度标识符】")
	w := 12.345678910
	fmt.Printf("%%f默认宽度，默认精度:%f\n", w)     //12.345679
	fmt.Printf("%%9f宽度9，默认精度:%9f\n", w)    //12.345679
	fmt.Printf("%%.2f默认宽度，精度2:%.2f\n", w)  //12.35
	fmt.Printf("%%9.2f宽度9，精度2:%9.2f\n", w) //    12.35
	fmt.Printf("%%9.f宽度9，精度0:%9.f\n", w)   //       12

}

// 获取输入【需要在main函数中执行，否则无法达到控制台输入的效果】
func getInput() {
	//Go语言fmt包下有fmt.Scan、fmt.Scanf、fmt.Scanln三个函数，可以在程序运行过程中从标准输入获取用户的输入。
	fmt.Println(">:获取输入【需要在main函数中执行，否则无法达到控制台输入的效果】")
	fmt.Println("fmt.Scan")
	/*
		fmt.Scan
		func Scan(a ...interface{}) (n int, err error)
		Scan从标准输入扫描文本，读取由空白符分隔的值保存到传递给本函数的参数中，换行符视为空白符。
		本函数返回成功扫描的数据个数和遇到的任何错误。如果读取的数据个数比提供的参数少，会返回一个错误报告原因。
		fmt.Scan从标准输入中扫描用户输入的数据，将以空白符分隔的数据分别存入指定的参数。
	*/
	var (
		name    string
		age     int
		married bool
	)
	fmt.Scan(&name, &age, &married)
	//将代码编译后在终端执行，在终端依次输入Meta39、18和false使用空格分隔，不用按回车。
	fmt.Printf("扫描结果 name:%s age:%d married:%t \n", name, age, married)

	/*
		fmt.Scanf
		func Scanf(format string, a ...interface{}) (n int, err error)
		Scanf从标准输入扫描文本，根据format参数指定的格式去读取由空白符分隔的值保存到传递给本函数的参数中。
		本函数返回成功扫描的数据个数和遇到的任何错误。
		fmt.Scanf不同于fmt.Scan简单的以空格作为输入数据的分隔符，fmt.Scanf为输入数据指定了具体的输入内容格式，只有按照格式输入数据才会被扫描并存入对应变量。
	*/
	fmt.Println("fmt.Scanf")
	fmt.Scanf("1:%s 2:%d 3:%t", &name, &age, &married)
	//将上面的代码编译后在终端执行，在终端按照指定的格式依次输入1:Meta39 2:18 3:false
	fmt.Printf("扫描结果 name:%s age:%d married:%t \n", name, age, married)

	/*
		fmt.Scanln
		func Scanln(a ...interface{}) (n int, err error)
		Scanln类似Scan，它在遇到换行时才停止扫描。最后一个数据后面必须有换行或者到达结束位置。
		本函数返回成功扫描的数据个数和遇到的任何错误。
		fmt.Scanln遇到回车就结束扫描了，这个比较常用。
	*/
	fmt.Println("fmt.Scanln")
	fmt.Scanln(&name, &age, &married)
	//将上面的代码编译后在终端执行，在终端依次输入Meta39、18和false使用空格分隔，然后按回车键。
	fmt.Printf("扫描结果 name:%s age:%d married:%t \n", name, age, married)

	/*
		bufio.NewReader
		有时候我们想完整获取输入的内容，而输入的内容可能包含空格，这种情况下可以使用bufio包来实现。
	*/
	reader := bufio.NewReader(os.Stdin) // 从标准输入生成读对象
	fmt.Print("请输入内容：")
	text, _ := reader.ReadString('\n') // 读到换行，即：回车键或者换行符
	text = strings.TrimSpace(text)
	fmt.Printf("%#v\n", text)

	/*
		Fscan系列
		这几个函数功能分别类似于fmt.Scan、fmt.Scanf、fmt.Scanln三个函数，只不过它们不是从标准输入中读取数据而是从io.Reader中读取数据。
		func Fscan(r io.Reader, a ...interface{}) (n int, err error)
		func Fscanln(r io.Reader, a ...interface{}) (n int, err error)
		func Fscanf(r io.Reader, format string, a ...interface{}) (n int, err error)
	*/

	/*
		Sscan系列
		这几个函数功能分别类似于fmt.Scan、fmt.Scanf、fmt.Scanln三个函数，只不过它们不是从标准输入中读取数据而是从指定字符串中读取数据。
		func Sscan(str string, a ...interface{}) (n int, err error)
		func Sscanln(str string, a ...interface{}) (n int, err error)
		func Sscanf(str string, format string, a ...interface{}) (n int, err error)
	*/

}
