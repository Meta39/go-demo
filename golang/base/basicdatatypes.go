package main

import (
	"fmt"
	"math"
)

/*
BasicDataTypes 基本数据类型
Go语言中有丰富的数据类型，除了基本的整型、浮点型、布尔型、字符串外，还有数组、切片、结构体、函数、map、通道（channel）等。
Go 语言的基本类型和其他语言大同小异。
*/

func BasicDataTypes() {
	fmt.Println("============ 基本数据类型 ============")

	/*
		整型
		分为以下两个大类： 按长度分为：int8、int16、int32、int64 对应的无符号整型：uint8、uint16、uint32、uint64
		其中，uint8就是我们熟知的byte型，int16对应C语言中的short型，int64对应C语言中的long型。
		uint8	无符号 8位整型 (0 到 255)
		uint16	无符号 16位整型 (0 到 65535)
		uint32	无符号 32位整型 (0 到 4294967295)
		uint64	无符号 64位整型 (0 到 18446744073709551615)
		int8	有符号 8位整型 (-128 到 127)
		int16	有符号 16位整型 (-32768 到 32767)
		int32	有符号 32位整型 (-2147483648 到 2147483647)
		int64	有符号 64位整型 (-9223372036854775808 到 9223372036854775807)

		特殊整型
				uint	32位操作系统上就是uint32，64位操作系统上就是uint64
				int	32位操作系统上就是int32，64位操作系统上就是int64
				uintptr	无符号整型，用于存放一个指针
			注意： 在使用int和 uint类型时，不能假定它是32位或64位的整型，而是考虑int和uint可能在不同平台上的差异。
			注意事项 获取对象的长度的内建len()函数返回的长度可以根据不同平台的字节长度进行变化。
			实际使用中，切片或 map 的元素数量等都可以用int来表示。在涉及到二进制传输、读写文件的结构描述时，为了保持文件的结构不会受到不同编译目标平台字节长度的影响，不要使用int和 uint。
	*/

	/*
		浮点型
			Go语言支持两种浮点型数：float32和float64。
			这两种浮点型数据格式遵循IEEE 754 标准：
			float32 的浮点数的最大范围约为 3.4e38，可以使用常量定义：math.MaxFloat32。
			float64 的浮点数的最大范围约为 1.8e308，可以使用一个常量定义：math.MaxFloat64。
			打印浮点数时，可以使用fmt包配合动词%f，代码如下：
	*/
	fmt.Printf("%f\n", math.Pi)
	fmt.Printf("%.2f\n", math.Pi)

	//复数complex64和complex128。复数有实部和虚部，complex64的实部和虚部为32位，complex128的实部和虚部为64位。
	var c1 complex64
	c1 = 1 + 2i
	var c2 complex128
	c2 = 2 + 3i
	fmt.Println(c1)
	fmt.Println(c2)

	/*
				布尔值
			Go语言中以bool类型进行声明布尔型数据，布尔型数据只有true（真）和false（假）两个值。
		注意：
		1、布尔类型变量的默认值为false。
		2、Go 语言中不允许将整型强制转换为布尔型.
		3、布尔型无法参与数值运算，也无法与其他类型进行转换。
	*/
	boolean := true
	fmt.Println("打印布尔值", boolean)

	/*
			字符串 string，如果只是声明了一个字符串变量 s，但没有给它赋值。当我们检查 s 是否等于空字符串时，结果是 true，这说明 s 的零值是一个空字符串。
			字符串的常用操作
		方法			介绍
		len(str)	求长度
		+或fmt.Sprintf	拼接字符串
		strings.Split	分割
		strings.contains	判断是否包含
		strings.HasPrefix,strings.HasSuffix	前缀/后缀判断
		strings.Index(),strings.LastIndex()	子串出现的位置
		strings.Join(a[]string, sep string)	join操作
	*/
	var s string
	fmt.Println("s == \"\" ?", s == "") //true

	/*
		多行字符串
		Go语言中要定义一个多行字符串时，就必须使用反引号字符：
		反引号间换行将被作为字符串中的换行，但是所有的转义字符均无效，文本将会原样输出。
	*/
	stringBlocks := `第一行
第二行
第三行`
	fmt.Println(stringBlocks)

	/*
				byte和rune类型
			组成每个字符串的元素叫做“字符”，可以通过遍历或者单个获取字符串元素获得字符。 字符用单引号（’）包裹起来，如：
			var a = '中'
			var b = 'x'
		Go 语言的字符有以下两种：
		uint8类型，或者叫 byte 型，代表一个ASCII码字符。
		rune类型，代表一个 UTF-8字符。
		当需要处理中文、日文或者其他复合字符时，则需要用到rune类型。rune类型实际是一个int32。
		Go 使用了特殊的 rune 类型来处理 Unicode，让基于 Unicode 的文本处理更为方便，也可以使用 byte 型进行默认字符串处理，性能和扩展性都有照顾。
	*/
	var a = '中'
	var b = 'x'
	fmt.Println(a, b)

	traversalString()

	/*
		修改字符串
		要修改字符串，需要先将其转换成[]rune或[]byte，完成后再转换为string。无论哪种转换，都会重新分配内存，并复制字节数组。
	*/
	s1 := "big"
	// 强制类型转换
	byteS1 := []byte(s1)
	byteS1[0] = 'p'
	fmt.Println(string(byteS1))

	s2 := "白萝卜"
	runeS2 := []rune(s2)
	runeS2[0] = '红'
	fmt.Println(string(runeS2))

	//替换原始字符串内容
	//var s3 string = "原始字符串"
	s3 := "原始字符串"
	fmt.Println("初始化s3字符串内容为", s3)
	s3 = "修改后的字符串"
	fmt.Println("重新赋值s3字符串内容为", s3)

	/*
		类型转换
		Go语言中只有强制类型转换，没有隐式类型转换。该语法只能在两个类型之间支持相互转换的时候使用。
		强制类型转换的基本语法如下：
		T(表达式)
		其中，T表示要转换的类型。表达式包括变量、复杂算子和函数返回值等.
		比如计算直角三角形的斜边长时使用math包的Sqrt()函数，该函数接收的是float64类型的参数，而变量a和b都是int类型的，这个时候就需要将a和b强制类型转换为float64类型。
	*/
	sqrtDemo()

	//单个字符串强转字符(byte)，需要转为byte[]数组
	s4 := "a"
	b4 := []byte(s4)
	//%v 用于以默认格式打印变量的值，适用于大多数类型。
	//%c 用于将整数解释为 Unicode 码点并打印对应的字符，适用于整数或字符类型。
	fmt.Printf("string转为byte[]后的值为%v，Unicode 码点并打印对应的字符为%c\n", b4, b4)

	fmt.Println("============ 基本数据类型 ============")
}

// 遍历字符串
/*
byte：104(h) 101(e) 108(l) 108(l) 111(o) 230(æ) 178(²) 153() 230(æ) 178(²) 179(³)
rune：104(h) 101(e) 108(l) 108(l) 111(o) 27801(沙) 27827(河)
因为UTF8编码下一个中文汉字由3~4个字节组成，所以我们不能简单的按照字节去遍历一个包含中文的字符串，否则就会出现上面输出中第一行的结果。
字符串底层是一个byte数组，所以可以和[]byte类型相互转换。
字符串是不能修改的 字符串是由byte字节组成，所以字符串的长度是byte字节的长度。
rune类型用来表示utf8字符，一个rune字符由一个或多个byte组成。
*/
func traversalString() {
	s := "hello米特"
	for i := 0; i < len(s); i++ { //byte
		fmt.Printf("%v(%c) ", s[i], s[i])
	}
	fmt.Println()
	for _, r := range s { //rune，处理中文
		fmt.Printf("%v(%c) ", r, r)
	}
	fmt.Println()
}

func sqrtDemo() {
	var a, b = 3, 4
	var c int
	// math.Sqrt()接收的参数是float64类型，需要强制转换
	c = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(c)
}
