package structs

import (
	"errors"
	"fmt"
)

/*
MyStruct 自定义结构体
这样我们就拥有了一个 MyStruct 的自定义类型，它有name、city、age三个字段，分别表示姓名、城市和年龄。
这样我们使用这个 MyStruct 结构体就能够很方便的在程序中表示和存储人信息了。
语言内置的基础数据类型是用来描述一个值的，而结构体是用来描述一组值的。
比如一个人有名字、年龄和居住城市等，本质上是一种聚合型的数据类型
结构体字段的可见性：结构体中字段大写开头表示可公开访问，小写表示私有（仅在定义当前结构体的包中可访问）。
*/
type MyStruct struct {
	Name       string //变量名如果是小写，包外面就无法访问了，类似private，大写的话，包外也能访问，类似public
	City       string
	Age        int8
	Sex, IsDel bool //同样类型的字段也可以写在一行（但是不推荐！！！）
}

/*
NewMyStruct 构造函数
Go语言的结构体没有构造函数，我们可以自己实现。
例如，下方的代码就实现了一个 MyStruct 的构造函数。
因为 struct 是值类型，如果结构体比较复杂的话，值拷贝性能开销会比较大，所以该构造函数返回的是结构体指针类型。
*/
func NewMyStruct(name, city string, age int8, sex, isDel bool) *MyStruct {
	return &MyStruct{
		Name:  name,
		City:  city,
		Age:   age,
		Sex:   sex,
		IsDel: isDel,
	}
}

/*
Dream 方法
Go语言中的方法（Method）是一种作用于特定类型变量的函数。
这种特定类型变量叫做接收者（Receiver）。
接收者的概念就类似于其他语言中的this或者 self。
方法的定义格式如下：

	func (接收者变量 接收者类型) 方法名(参数列表) (返回参数) {
	    函数体
	}

其中，
1、接收者变量：接收者中的参数变量名在命名时，官方建议使用接收者类型名称首字母的小写，而不是self、this之类的命名。
例如，Person类型的接收者变量应该命名为 p，Connector类型的接收者变量应该命名为c等。
2、接收者类型：接收者类型和参数类似，可以是指针类型和非指针类型。
3、方法名、参数列表、返回参数：具体格式与函数定义相同。
*/
func (p MyStruct) Dream(param string) (result string, err error) {
	fmt.Println("调用方法")
	fmt.Printf("%s的梦想是学好Go语言！\n", p.Name)
	if param != "我有一个梦想" {
		return "", errors.New("你让我没有爱啊，易小川！")
	}
	return "我是返回值。获取的参数是：" + param, nil
}

/*
SetAge 指针类型的接收者
指针类型的接收者由一个结构体的指针组成，由于指针的特性，调用方法时修改接收者指针的任意成员变量，在方法结束后，修改都是有效的。
这种方式就十分接近于其他语言中面向对象中的this或者self。
例如我们为 MyStruct 添加一个 SetAge 方法，来修改实例变量的年龄。
*/
func (p *MyStruct) SetAge(age int8) {
	p.Age = age
}

/*
SetAge2 值类型的接收者
当方法作用于值类型接收者时，Go语言会在代码运行时将接收者的值复制一份。
在值类型接收者的方法中可以获取接收者的成员值，但修改操作只是针对副本，无法修改接收者变量本身。
*/
func (p MyStruct) SetAge2(age int8) {
	p.Age = age
}
