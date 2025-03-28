package main

import (
	"encoding/json"
	"fmt"
	"golang/base/structs"
	"unsafe"
)

/*
Struct 结构体
1、Go语言中没有“类”的概念，也不支持“类”的继承等面向对象的概念。
2、Go语言中通过结构体的内嵌再配合接口比面向对象具有更高的扩展性和灵活性。
3、类型别名和自定义类型：

一、自定义类型
（1）在Go语言中有一些基本的数据类型，如string、整型、浮点型、布尔等数据类型， Go语言中可以使用type关键字来定义自定义类型。
（2）自定义类型是定义了一个全新的类型。我们可以基于内置的基本类型定义，也可以通过struct定义。
//将MyInt定义为int类型。通过type关键字的定义，MyInt就是一种新的类型，它具有int的特性。
type MyInt int

二、类型别名
（1）类型别名是Go1.9版本添加的新功能。
（2）类型别名规定：TypeAlias只是Type的别名，本质上TypeAlias与Type是同一个类型。
就像一个孩子小时候有小名、乳名，上学后用学名，英语老师又会给他起英文名，但这些名字都指的是他本人。
type TypeAlias = Type

rune和byte就是类型别名，他们的定义如下：
type byte = uint8
type rune = int32

三、结构体
Go语言中的基础数据类型可以表示一些事物的基本属性，但是当我们想表达一个事物的全部或部分属性时，这时候再用单一的基本数据类型明显就无法满足需求了。
Go语言提供了一种自定义数据类型，可以封装多个基本数据类型，这种数据类型叫结构体，英文名称struct。
也就是我们可以通过struct来定义自己的类型了。
Go语言中通过struct来实现面向对象。
PS：可以粗略的理解就是java的class，但又不完全是class，因为struct不能继承，而class可以继承。

结构体的定义：
使用type和struct关键字来定义结构体，具体代码格式如下：

	type 类型名 struct {
	    字段名 字段类型
	    字段名 字段类型
	    …
	}

其中：
（1）类型名：标识自定义结构体的名称，在同一个包内不能重复。
（2）字段名：表示结构体字段名。结构体中的字段名必须唯一。
（3）字段类型：表示结构体字段的具体类型。
*/
func Struct() {
	fmt.Println("============ 结构体 ============")
	/*
		类型定义与类型别名的区别
		结果显示a的类型是main.NewInt，表示main包下定义的NewInt类型。
		b的类型是int。MyInt类型只会在代码中存在，编译完成时并不会有MyInt类型。
	*/
	var a NewInt                        //类型定义，表示一个新的类型，本质已经发生改变
	var b MyInt                         //类型别名，表示一个新的名字，本质没有发生改变
	fmt.Printf("类型定义type of a:%T\n", a) //type of a:main.NewInt
	fmt.Printf("类型别名type of b:%T\n", b) //type of b:int

	/*
		结构体实例化
		1、只有当结构体实例化时，才会真正地分配内存。也就是必须实例化后才能使用结构体的字段。
		2、结构体本身也是一种类型，我们可以像声明内置类型一样使用var关键字声明结构体类型。
		PS：跟java是一样的，只有实例化了才能进行赋值操作。只不过java是通过new关键字进行实例化，而go是用var关键字声明了结构体后就已经实例化了。
		go也可以使用new关键字进行实例化（创建指针类型结构体）
	*/
	fmt.Println("自定义结构体MyStruct，实例化并进行赋值操作，没有等号，因此是go自动赋值的。")
	var myStruct structs.MyStruct
	//var myInnerStruct = structs.innerStruct//无法使用，因为innerStruct首字母是小写开头，只能在structs包里使用。更不用说里面的变量了。
	myStruct.Name = "Meta39"
	myStruct.City = "广州"
	myStruct.Age = 18
	myStruct.Sex = true
	myStruct.IsDel = false
	fmt.Printf("var关键字常规实例化myStruct=%v\n", myStruct)  //myStruct={Meta39 广州 18 true false}
	fmt.Printf("var关键字常规实例化myStruct=%#v\n", myStruct) //myStruct=structs.MyStruct{Name:"Meta39", City:"广州", Age:18, Sex:true, IsDel:false}

	//匿名结构体
	var user struct {
		Name string
		Age  int
	}
	user.Name = "小王子"
	user.Age = 18
	fmt.Printf("匿名结构体user=%#v\n", user)

	//创建指针类型结构体。我们还可以通过使用new关键字对结构体进行实例化，得到的是结构体的地址。
	fmt.Println("通过new关键字进行实例化并赋值给变量myStruct2，因为有等号，等号表示赋值")
	var myStruct2 = new(structs.MyStruct)
	myStruct2.Name = "Meta40"
	myStruct2.City = "佛山"
	myStruct2.Age = 19
	myStruct2.Sex = false
	myStruct2.IsDel = true
	fmt.Printf("new关键字实例化myStruct2类型%T\n", myStruct2) //*structs.MyStruct
	fmt.Printf("new关键字实例化myStruct2=%#v\n", myStruct2) //myStruct2=&structs.MyStruct{Name:"Meta40", City:"佛山", Age:19, Sex:false, IsDel:true}
	//从打印的结果中我们可以看出myStruct2是一个结构体指针。需要注意的是在Go语言中支持对结构体指针直接使用.来访问结构体的成员。

	/*
		取结构体的地址实例化
		1、使用&对结构体进行取地址操作相当于对该结构体类型进行了一次new实例化操作。
		2、p3.name = "七米"其实在底层是(*p3).name = "七米"，这是Go语言帮我们实现的语法糖。
	*/
	p3 := &structs.MyStruct{}
	p3.Name = "七米"
	p3.Age = 30
	p3.City = "成都"
	fmt.Printf("取结构体的地址实例化p3=%#v\n", p3) //&structs.MyStruct{Name:"七米", City:"成都", Age:30, Sex:false, IsDel:false}

	/*
		结构体初始化
		没有初始化的结构体，其成员变量都是对应其类型的零值。
	*/
	var p4 structs.MyStruct
	fmt.Printf("实例化却没有初始化(赋值)的结构体，其成员变量都是对应其类型的零值。p4=%#v\n", p4)

	/*
			使用键值对初始化。
			PS：相当于java里的new关键字实例化并赋值。如下所示：

			import lombok.Data;

			@Data
			public class MyClass {
				private String name;
			}

			//实例化并赋值
			MyClass myClass = new MyClass(){{
				this.setName("Meta39");
		    }};
	*/
	p5 := structs.MyStruct{
		Name: "小王子",
		City: "北京",
		Age:  18,
	}
	fmt.Printf("键值对初始化p5=%#v\n", p5) //structs.MyStruct{Name:"小王子", City:"北京", Age:18, Sex:false, IsDel:false}

	//结构体指针键值对初始化
	p6 := &structs.MyStruct{
		Name: "小王子",
		City: "北京",
		Age:  18,
	}
	fmt.Printf("结构体指针键值对初始化p6=%#v\n", p6)

	//当某些字段没有初始值的时候，该字段可以不写。此时，没有指定初始值的字段的值就是该字段类型的零值。
	p7 := &structs.MyStruct{
		City: "北京",
	}
	fmt.Printf("当某些字段没有初始值的时候，该字段可以不写。此时，没有指定初始值的字段的值就是该字段类型的零值。p7=%#v\n", p7)
	//p7=&structs.MyStruct{Name:"", City:"北京", Age:0, Sex:false, IsDel:false}

	/*
		使用值的列表初始化。初始化结构体的时候可以简写，也就是初始化的时候不写键，直接写值。
		但是GoLand编译器并不推荐这种写法，因为可能会写错值，导致增加排查问题的难度。（推荐：结构体指针键值对初始化）
		使用这种格式初始化时，需要注意：
		1、必须初始化结构体的所有字段。
		2、初始值的填充顺序必须与字段在结构体中的声明顺序一致。
		3、该方式不能和键值初始化方式混用。
		PS：相当于java里class的全参构造函数，但是go里面是自带无参和全参的。
		java需要自己手写全参构造函数，如果手写了全参构造，还要写无参构造。
		因为java里面默认实现无参构造，如果有其它构造函数，则默认的无参构造会失效，因此还要手动定义无参构造。
	*/
	p8 := &structs.MyStruct{
		"沙河娜扎",
		"北京",
		28,
		false,
		false,
	}
	fmt.Printf("使用值的列表初始化。初始化结构体的时候可以简写，也就是初始化的时候不写键，直接写值。p8=%#v\n", p8)
	//p8=&structs.MyStruct{Name:"沙河娜扎", City:"北京", Age:28, Sex:false, IsDel:false}

	//结构体内存布局。结构体占用一块连续的内存。
	fmt.Printf("结构体占用一块连续的内存。")
	n := testStruct{
		1, 2, 3, 4,
	}
	fmt.Printf("n.a %p\n", &n.a)
	fmt.Printf("n.b %p\n", &n.b)
	fmt.Printf("n.c %p\n", &n.c)
	fmt.Printf("n.d %p\n", &n.d)

	//空结构体。不占用空间。
	var n2 testStruct2
	var n3 testStruct
	fmt.Println("n2空结构体，不占用空间", unsafe.Sizeof(n2)) // 0
	fmt.Println("n3非空结构体，占用空间", unsafe.Sizeof(n3)) // 4

	//构造函数
	p9 := structs.NewMyStruct("张三", "沙河", 90, true, false)
	fmt.Printf("%#v\n", p9) //&structs.MyStruct{Name:"张三", City:"沙河", Age:90, Sex:true, IsDel:false}

	//调用方法。方法与函数的区别是，函数不属于任何类型，方法属于特定的类型。
	res, err := p9.Dream("我有一个梦想")
	if err != nil {
		fmt.Println("报错啦~", err)
		return //报错了，应阻止当前函数继续往下执行。因为后面的代码都是正常的逻辑。
	}
	fmt.Println("res:", res)

	/*
		指针类型的接收者和值类型的接收者

		特性			值接收者			指针接收者
		修改原始值	不能				可以
		内存开销		每次调用创建副本	共享同一内存
		并发安全		安全				需要额外同步
		适用大小		小型结构体		大型结构体
		nil检查		不需要			需要

		什么时候应该使用指针类型接收者？
		1、需要修改接收者中的值
		2、接收者是拷贝代价比较大的大对象
		3、保证一致性，如果有某个方法使用了指针接收者，那么其他的方法也应该使用指针接收者。

		什么时候应该使用值类型的接收者？
		1、当类型是不可变的小对象时
		2、当方法不需要修改接收者时
		3、当类型是基本类型或小型结构体时
		4、当需要保证方法调用不会修改原始值时

		记住：当不确定时，优先使用指针接收者，特别是对于结构体类型。
	*/

	//指针类型的接收者
	fmt.Println("指针类型的接收者修改前的年龄：", p9.Age) // 90
	p9.SetAge(20)
	fmt.Println("指针类型的接收者修改后的年龄：", p9.Age) // 20

	//值类型的接收者（无法修改接收者变量本身）
	fmt.Println("值类型的接收者修改前的年龄：", p9.Age) // 20
	p9.SetAge2(25)
	fmt.Println("值类型的接收者修改后的年龄：", p9.Age) // 20

	//结构体的匿名字段
	p10 := structs.UseAnonymousStructParam()
	fmt.Printf("结构体的匿名字段%#v\n", p10) //&structs.AnonymousStructParam{string:"小王子", int:18}

	//嵌套结构体（类似java里的对象包对象）
	user1 := structs.User{
		Name:   "小王子",
		Gender: "男",
		Address: structs.Address{
			Province: "山东",
			City:     "威海",
		},
	}
	fmt.Printf("嵌套结构体 user1=%#v\n", user1) //structs.User{Name:"小王子", Gender:"男", Address:structs.Address{Province:"山东", City:"威海"}}

	//嵌套匿名字段(上面user结构体中嵌套的Address结构体也可以采用匿名字段的方式)
	var user2 structs.User2
	user2.Name = "小王子"
	user2.Gender = "男"
	user2.Address.Province = "山东"           // 匿名字段默认使用类型名作为字段名
	user2.City = "威海"                       // 匿名字段可以省略，即：user2.Address.City =》 user2.City
	fmt.Printf("嵌套匿名字段 user2=%#v\n", user2) //structs.User2{Name:"小王子", Gender:"男", Address:structs.Address{Province:"山东", City:"威海"}}

	/*
		嵌套结构体的字段名冲突
		嵌套结构体内部可能存在相同的字段名。在这种情况下为了避免歧义需要通过指定具体的内嵌结构体字段名。
		即：全路径赋值
	*/
	var user3 structs.User3
	user3.Name = "沙河娜扎"
	user3.Gender = "男"
	//user3.CreateTime = "2019" //Ambiguous reference 'CreateTime'，引用“CreateTime”不明确。不知道是Address2.CreateTime、Email.CreateTime中的哪一个
	user3.Address2.CreateTime = "1000" //指定Address2结构体中的CreateTime
	user3.Email.CreateTime = "2000"    //指定Email结构体中的CreateTime
	fmt.Printf("嵌套结构体的字段名冲突 user3=%#v\n", user2)

	//结构体的“继承”。Go语言中使用结构体也可以实现其他编程语言中面向对象的继承。
	d1 := &structs.Dog{
		Feet: 4,
		Animal: &structs.Animal{ //注意嵌套的是结构体指针
			Name: "旺财",
		},
	}
	d1.Wang() //旺财会狗叫
	d1.Move() //旺财会行走

	/*
		结构体与JSON序列化
		JSON(JavaScript Object Notation) 是一种轻量级的数据交换格式。易于人阅读和编写。同时也易于机器解析和生成。
		JSON键值对是用来保存JS对象的一种方式，键/值对组合中的键名写在前面并用双引号""包裹，使用冒号:分隔，然后紧接着值；多个键值之间使用英文,分隔。
	*/
	c := &structs.Class{
		Title:    "101",
		Students: make([]structs.Student, 0, 200),
	}
	for i := 0; i < 10; i++ {
		stu := structs.Student{
			Name:   fmt.Sprintf("stu%02d", i),
			Gender: "男",
			ID:     i,
		}
		c.Students = append(c.Students, stu)
	}
	//JSON序列化：结构体-->JSON格式的字符串
	data, err2 := json.Marshal(c)
	if err2 != nil {
		fmt.Println("json marshal failed")
		return
	}
	fmt.Printf("JSON序列化：结构体-->JSON格式的字符串 json:%s\n", data)

	//JSON反序列化：JSON格式的字符串-->结构体
	str := `{"Title":"101","Students":[{"ID":0,"Gender":"男","Name":"stu00"},{"ID":1,"Gender":"男","Name":"stu01"},{"ID":2,"Gender":"男","Name":"stu02"},{"ID":3,"Gender":"男","Name":"stu03"},{"ID":4,"Gender":"男","Name":"stu04"},{"ID":5,"Gender":"男","Name":"stu05"},{"ID":6,"Gender":"男","Name":"stu06"},{"ID":7,"Gender":"男","Name":"stu07"},{"ID":8,"Gender":"男","Name":"stu08"},{"ID":9,"Gender":"男","Name":"stu09"}]}`
	c1 := &structs.Class{}
	err = json.Unmarshal([]byte(str), c1)
	if err != nil {
		fmt.Println("json unmarshal failed!")
		return
	}
	fmt.Printf("JSON反序列化：JSON格式的字符串-->结构体 struct:%#v\n", c1)

	/*
		结构体标签（Tag）
		1、Tag是结构体的元信息，可以在运行的时候通过反射的机制读取出来。
		2、Tag在结构体字段的后方定义，由一对反引号包裹起来，具体的格式如下：
			`key1:"value1" key2:"value2"`
		3、结构体tag由一个或多个键值对组成。键与值使用冒号分隔，值用双引号括起来。
		4、同一个结构体字段可以设置多个键值对tag，不同的键值对之间使用空格分隔。
		注意事项： 为结构体编写Tag时，必须严格遵守键值对的规则。结构体标签的解析代码的容错能力很差，一旦格式写错，编译和运行时都不会提示任何错误，通过反射也无法正确取值。
		例如不要在key和value之间添加空格。
		例如我们为 Student2 结构体的每个字段定义json序列化时使用的Tag
	*/
	student2 := structs.Student2{
		ID:     1,
		Gender: "Meta39",
	}
	marshalStudent2, err3 := json.Marshal(student2)
	if err3 != nil {
		fmt.Println("json unmarshal failed!")
		return
	}

	//ID => id，name不序列化
	fmt.Printf("结构体标签（Tag）JSON序列化：结构体-->JSON格式的字符串 json:%s\n", marshalStudent2) //{"id":1,"Gender":"Meta39"}

	//结构体和方法补充知识点。因为slice和map这两种数据类型都包含了指向底层数据的指针，因此我们在需要复制它们时要特别注意。
	//如果需要data2赋值完p11里面的Dreams内容，后续修改data2，同时也要修改Dreams，推荐使用SetDreams
	p11 := structs.Person{Name: "小王子", Age: 18}
	data2 := []string{"吃饭", "睡觉", "打豆豆"}
	p11.SetDreams(data2)                     //SetDreams传递的是切片，因此共享底层切片。这里修改切片的内容，SetDreams里面的切片也会受到影响
	fmt.Println("SetDreams data2", data2)    //[吃饭 睡觉 打豆豆]
	fmt.Println("SetDreams p11", p11.Dreams) //[吃饭 睡觉 打豆豆]
	data2[1] = "不睡觉"                         //data2修改时，也会修改p11的Dreams数据。因为共享底层切片。
	fmt.Println("修改data2后 data2", data2)     //[吃饭 不睡觉 打豆豆]
	fmt.Println("修改data2后 p11", p11.Dreams)  //[吃饭 不睡觉 打豆豆]

	//如果需要data3赋值完p12里面的Dreams内容不受影响，推荐使用SetDreams2
	p12 := structs.Person{Name: "小王子", Age: 18}
	data3 := []string{"吃饭", "睡觉", "打豆豆"}
	p12.SetDreams2(data3)                     //data3修改，不会修改p12的数据。各自独立
	fmt.Println("SetDreams2 data3", data3)    //[吃饭 睡觉 打豆豆]
	fmt.Println("SetDreams2 p12", p12.Dreams) //[吃饭 睡觉 打豆豆]
	data3[1] = "不睡觉"
	fmt.Println("修改data3后 data3", data3)    //[吃饭 不睡觉 打豆豆]
	fmt.Println("修改data3后 p12", p12.Dreams) //[吃饭 睡觉 打豆豆]

	fmt.Println("============ 结构体 ============")
}

// NewInt 类型定义，没有等号，说明是一个全新的类型
type NewInt int

// MyInt 类型别名，有等号，说明还是原来的int类型
type MyInt = int

type testStruct struct {
	a int8
	b int8
	c int8
	d int8
}

// 空结构体
type testStruct2 struct{}
