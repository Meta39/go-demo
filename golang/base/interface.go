package main

import (
	"fmt"
	"golang/base/interfaces"
	"golang/base/interfaces/implements/pay"
	"golang/base/interfaces/implements/sing"
)

/*
Interface 接口。定义了一个对象的行为规范，只定义规范不实现，由具体的对象来实现规范的细节。
在Go语言中接口（interface）是一种类型，一种抽象的类型。
相较于之前章节中讲到的那些具体类型（字符串、切片、结构体等）更注重“我是谁”，接口类型更注重“我能做什么”的问题。
接口类型就像是一种约定——概括了一种类型应该具备哪些方法，在Go语言中提倡使用面向接口的编程方式实现解耦。
PS：和java差不多，只是go接口里没有默认方法。但go接口里可以有私有方法，java接口只有公有方法。go接口实现类不需要实现所有接口，java实现类必须实现所有接口。

接口类型：
一、接口是一种由程序员来定义的类型，一个接口类型就是一组方法的集合，它规定了需要实现的所有方法。
二、相较于使用结构体类型，当我们使用接口类型说明相比于它是什么更关心它能做什么。

接口的定义：
每个接口类型由任意个方法签名组成，接口的定义格式如下：

	type 接口类型名 interface {
	    方法名1( 参数列表1 ) 返回值列表1
	    方法名2( 参数列表2 ) 返回值列表2
	    …
	}

其中：
一、接口类型名：Go语言的接口在命名时，一般会在单词后面添加er，如有写操作的接口叫Writer，有关闭操作的接口叫closer等。接口名最好要能突出该接口的类型含义。
二、方法名：当方法名首字母是大写且这个接口类型名首字母也是大写时，这个方法可以被接口所在的包（package）之外的代码访问。
三、参数列表、返回值列表：参数列表和返回值列表中的参数变量名可以省略。

举个例子，定义一个包含Write方法的Writer接口。

	type Writer interface {
	    Write([]byte) error
	}

当你看到一个Writer接口类型的值时，你不知道它是什么，唯一知道的就是可以通过调用它的Write方法来做一些事情。

实现接口的条件：
接口就是规定了一个需要实现的方法列表，在 Go 语言中一个类型只要实现了接口中规定的所有方法，那么我们就称它实现了这个接口。

注意事项：
由于接口类型变量能够动态存储不同类型值的特点，所以很多初学者会滥用接口类型（特别是空接口）来实现编码过程中的便捷。
只有当有两个或两个以上的具体类型必须以相同的方式进行处理时才需要定义接口。切记不要为了使用接口类型而增加不必要的抽象，导致不必要的运行时损耗。
在 Go 语言中接口是一个非常重要的概念和特性，使用接口类型能够实现代码的抽象和解耦，也可以隐藏某个功能的内部实现，但是缺点就是在查看源码的时候，不太方便查找到具体实现接口的类型。
相信很多读者在刚接触到接口类型时都会有很多疑惑，请牢记接口是一种类型，一种抽象的类型。
区别于我们在之前章节提到的那些具体类型（整型、数组、结构体类型等），它是一个只要求实现特定方法的抽象类型。
*/
func Interface() {
	fmt.Println("============ 接口 ============")

	//实例化 Bird 结构体，然后调用 Bird 实现 Singer 接口的 Sing() 方法
	fmt.Println("实例化 Bird 结构体，然后调用 Bird 实现 Singer 接口的 Sing() 方法")
	bird := sing.Bird{}
	bird.Sing()

	/*
		面向接口编程
		PHP、Java等语言中也有接口的概念，不过在PHP和Java语言中需要显式声明一个类实现了哪些接口，在Go语言中使用隐式声明的方式实现接口。
		只要一个类型实现了接口中规定的所有方法，那么它就实现了这个接口。
		Go语言中的这种设计符合程序开发中抽象的一般规律，例如在下面的代码示例中，我们的电商系统最开始只设计了支付宝一种支付方式：
		1.定义Payer支付接口
		2.定义ZhiFuBao结构体并实现Payer接口中的Pay方法
		3.定义WeChat结构体并实现Payer接口中的Pay方法
		4.在Payer支付接口里定义一个Checkout结账函数，防止写重复 obj.Pay(金额) 代码
	*/
	fmt.Println("面向接口编程：定义支付接口，分别实现支付宝、微信支付")
	interfaces.Checkout(&pay.ZhiFuBao{})
	interfaces.Checkout(&pay.WeChat{})

	/*
		接口类型变量
		那实现了接口又有什么用呢？一个接口类型的变量能够存储所有实现了该接口的类型变量。
		例如在上面的示例中，Dog和Bird类型均实现了 Singer 接口，此时一个 Singer 类型的变量就能够接收Bird和Dog类型的变量。
	*/
	fmt.Println("接口类型变量：")
	var x interfaces.Singer
	a := sing.Bird{} // 声明一个 Bird 类型变量a
	b := sing.Dog{}  // 声明一个 Dog 类型变量b
	x = a            // 可以把 Bird 类型变量直接赋值给x
	x.Sing()         // 叽叽喳喳
	x = b            // 可以把 Dog 类型变量直接赋值给x
	x.Sing()         // 汪汪汪

	/*
		值接收者和指针接收者
		特性			值接收者						指针接收者
		调用对象类型	值或指针均可调用（Go 自动解引用）	必须通过指针调用（否则编译错误）
		接口实现规则	值类型和指针类型均实现接口		只有指针类型实现接口
		数据修改能力	无法修改原结构体				可以修改原结构体
		性能开销		复制整个结构体（大对象时性能低）	仅复制指针（高效）
		零值可用性	安全（操作副本）				需注意 nil 指针

		优先使用值接收者的场景：
		一、不可变数据：方法不需要修改原结构体（纯计算或读取操作）。
		二、小型结构体：结构体很小（如基本类型、少量字段），复制开销可忽略。
		三、并发安全：避免意外修改共享数据（副本隔离）。

		优先使用指针接收者的场景：
		一、需要修改原数据：方法需修改结构体内部状态
		二、大型结构体：避免复制大对象的性能开销（如包含数组或嵌套结构）。
		三、实现某些标准接口：如 fmt.Stringer、error 等常用指针接收者。
		四、结构体包含引用类型：如切片、映射、通道等，需通过指针统一修改。

		总结：
		默认优先使用指针接收者（除非明确不需要修改数据且结构体很小）。
		保持一致性：同一结构体的所有方法尽量统一接收者类型。
	*/
	fmt.Println("值接收者：")
	var x2 interfaces.Singer
	d2 := sing.Dog{} // d2是值类型
	x2 = d2          // 可以将d2赋值给变量x2
	x2.Sing()
	d3 := &sing.Dog{} //d3是指针类型
	x2 = d3
	x2.Sing() // 也可以将d3赋值给变量x2
	//从上面的代码中我们可以发现，使用值接收者实现接口之后，不管是结构体类型还是对应的结构体指针类型的变量都可以赋值给该接口变量。

	fmt.Println("指针接收者：")
	var x3 interfaces.Singer
	var c1 = &sing.Cat{} // c1是*Cat类型
	x3 = c1              // 可以将c1当成Singer类型
	x3.Sing()
	var c2 = sing.Cat{} // c2是Cat类型
	//x3 = c2             //不能将c2当成Singer类型
	x3 = &c2 //只能取指针地址赋值给x3
	x3.Sing()
	//由于Go语言中有对指针求值的语法糖，对于值接收者实现的接口，无论使用值类型还是指针类型都没有问题。但是我们并不总是能对一个值求址，所以对于指针接收者实现的接口要额外注意。

	//一个类型实现多个接口，而接口间彼此独立，不知道对方的实现。
	fmt.Println("一个类型实现多个接口，而接口间彼此独立，不知道对方的实现。")
	var d4 = sing.Dog{}
	var s4 interfaces.Singer = d4 //PS：用java来理解就是 Map<K,V> = new HashMap<>()。s4表示接口，d4表示实现类。
	var j4 interfaces.Jumper = d4 //同理

	s4.Sing() //调用实现类的方法
	//s4.Jump() //即使d4实现了Jumper接口的Jump方法也无法使用，因为此时s4是Singer接口，Singer不包含Jump方法
	j4.Jump() //调用实现类的方法
	//j4.Sing() //同理。这就表示接口间彼此独立，不知道对方的实现。
	//因为d4分别实现了2个Singer、Jumper接口的Sing、Jump方法，因此d4可以同时使用Sing、Jump方法
	fmt.Println("因为d4分别实现了2个Singer、Jumper接口的Sing、Jump方法，因此d4可以同时使用Sing、Jump方法")
	d4.Sing()
	d4.Jump()

	/*
		多种类型实现同一接口
		Go语言中不同的类型还可以实现同一接口。
		一个接口的所有方法，不一定需要由一个类型完全实现，接口的方法可以通过在类型中嵌入其他类型或者结构体来实现。
		如：WashingMachine接口定义了wash、dry方法。dryer结构体只实现了dry方法，没有实现wash方法。haier结构体只实现了wash方法，没有实现dry方法。
	*/
	fmt.Println("多种类型实现同一接口：WashingMachine接口定义了wash、dry方法。dryer结构体只实现了dry方法，没有实现wash方法。haier结构体只实现了wash方法，没有实现dry方法。")
	var washingMachine interfaces.WashingMachine
	fmt.Println(washingMachine)

	/*
		接口组合
		接口与接口之间可以通过互相嵌套形成新的接口类型
		对于这种由多个接口类型组合形成的新接口类型，同样只需要实现新接口类型中规定的所有方法就算实现了该接口类型。
		通过在结构体中嵌入一个接口类型，从而让该结构体类型实现了该接口类型，并且还可以改写该接口的方法。
	*/
	fmt.Println("接口组合")
	var singJumper interfaces.SingJumper
	fmt.Println(singJumper)

	fmt.Println("通过在结构体中嵌入一个接口类型，从而让该结构体类型实现了该接口类型，并且还可以改写该接口的方法。")
	superCat := sing.SuperCat{}
	//调用重写方法Sing。输出的不再是：喵喵喵
	superCat.Sing() //喵喵喵？只有低等猫才会喵喵喵。人类！我是超级猫！我会说人话！你会说猫话吗？

	/*
		空接口
		空接口是指没有定义任何方法的接口类型。因此任何类型都可以视为实现了空接口。
		也正是因为空接口类型的这个特性，空接口类型的变量可以存储任意类型的值。
		通常我们在使用空接口类型时不必使用type关键字声明，可以像下面的代码一样直接使用interface{}。
	*/
	fmt.Println("空接口类型的变量可以存储任意类型的值。")
	var anyer interfaces.Any
	anyer = "你好"
	fmt.Printf("type:%T value:%v\n", anyer, anyer)
	anyer = 100 // int型
	fmt.Printf("type:%T value:%v\n", anyer, anyer)
	anyer = true // 布尔型
	fmt.Printf("type:%T value:%v\n", anyer, anyer)
	anyer = sing.Dog{} // 结构体类型
	fmt.Printf("type:%T value:%v\n", anyer, anyer)

	var anyer2 interface{} // 声明一个空接口类型变量x
	fmt.Printf("type:%T value:%v\n", anyer2, anyer2)

	/*
		空接口的应用（PS：可以粗略理解为泛型变量，只是作用范围比较小。只支持下面2种情况。）
		1.空接口作为函数的参数：使用空接口实现可以接收任意类型的函数参数。
		2.空接口作为map的值：使用空接口实现可以保存任意值的字典。
	*/
	fmt.Println("空接口作为函数的参数")
	show(1)
	show("a")

	fmt.Println("空接口作为map的值")
	showMap()

	/*
		接口值
		由于接口类型的值可以是任意一个实现了该接口的类型值，所以接口值除了需要记录具体值之外，还需要记录这个值属于的类型。
		也就是说接口值由“类型”和“值”组成，鉴于这两部分会根据存入值的不同而发生变化，我们称之为接口的动态类型和动态值。
		接口值是支持相互比较的，当且仅当接口值的动态类型和动态值都相等时才相等。
	*/
	var bird2 interfaces.Singer = &sing.Bird{}
	var dog2 interfaces.Singer = &sing.Dog{}
	fmt.Println(bird2 == dog2) // false

	//但是有一种特殊情况需要特别注意，如果接口值保存的动态类型相同，但是这个动态类型不支持互相比较（比如切片），那么对它们相互比较时就会引发panic。
	var z interface{} = []int{1, 2, 3}
	fmt.Println(z)
	//fmt.Println(z == z) // panic: runtime error: comparing uncomparable type []int

	/*
		类型断言
		其语法格式：
		x.(T)
		其中：
		x：表示接口类型的变量
		T：表示断言x可能是的类型。
		该语法返回两个参数，第一个参数是x转化为T类型后的变量，第二个值是一个布尔值，若为true则表示断言成功，为false则表示断言失败。
	*/
	var n interfaces.Singer = &sing.Dog{}
	v, ok := n.(*sing.Dog)
	if ok {
		fmt.Println("类型断言成功。调用Sing方法")
		v.Sing()
	} else {
		fmt.Println("类型断言失败")
	}

	//如果对一个接口值有多个实际类型需要判断，推荐使用switch语句来实现。
	fmt.Println("如果对一个接口值有多个实际类型需要判断，推荐使用switch语句来实现。")
	justifyType("字符串")

	fmt.Println("============ 接口 ============")
}

// 空接口作为函数参数
func show(a interface{}) {
	fmt.Printf("type:%T value:%v\n", a, a)
}

func showMap() {
	// 空接口作为map值
	var studentInfo = make(map[string]interface{})
	studentInfo["name"] = "沙河娜扎"
	studentInfo["age"] = 18
	studentInfo["married"] = false
	fmt.Println(studentInfo)
}

// justifyType 对传入的空接口类型变量x进行类型断言
func justifyType(x interface{}) {
	switch v := x.(type) {
	case string:
		fmt.Printf("x is a string，value is %v\n", v)
	case int:
		fmt.Printf("x is a int is %v\n", v)
	case bool:
		fmt.Printf("x is a bool is %v\n", v)
	default:
		fmt.Println("unsupport type！")
	}
}
