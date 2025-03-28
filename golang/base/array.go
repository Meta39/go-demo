package main

import "fmt"

/*
Array 数组
1、数组是同一种数据类型元素的集合。 在Go语言中，数组从声明时就确定，使用时可以修改数组成员，但是数组大小不可变化。
2、数组的长度必须是常量，并且长度是数组类型的一部分。一旦定义，长度不能变。[5]int和[10]int是不同的类型。
3、数组可以通过下标进行访问，下标是从0开始，最后一个元素下标是：len-1，访问越界（下标在合法范围之外），则触发访问越界，会panic。
4、数组是值类型，赋值和传参会复制整个数组。因此改变副本的值，不会改变本身的值。
数组定义：
var 数组变量名 [元素数量]T
基本语法：
// 定义一个长度为3元素类型为int的数组a
var a [3]int

//数组长度不一样，不能相互赋值
var a [3]int
var b [4]int
a = b //不可以这样做，因为此时a和b是不同的类型
*/
func Array() {
	fmt.Println("============ 数组 ============")

	//数组的初始化
	fmt.Println("方法一：初始化数组时可以使用初始化列表来设置数组元素的值。")
	var testArray [3]int                        //数组会初始化为int类型的零值
	var numArray = [3]int{1, 2}                 //使用指定的初始值完成初始化
	var cityArray = [3]string{"北京", "上海", "深圳"} //使用指定的初始值完成初始化
	fmt.Println(testArray)                      //[0 0 0]
	fmt.Println(numArray)                       //[1 2 0]
	fmt.Println(cityArray)                      //[北京 上海 深圳]

	fmt.Println("方法二：按照上面的方法每次都要确保提供的初始值和数组长度一致，一般情况下我们可以让编译器根据初始值的个数自行推断数组的长度")
	var testArray2 [3]int
	var numArray2 = [...]int{1, 2}
	var cityArray2 = [...]string{"北京", "上海", "深圳"}
	fmt.Println(testArray2)                          //[0 0 0]
	fmt.Println(numArray2)                           //[1 2]
	fmt.Printf("type of numArray:%T\n", numArray2)   //type of numArray:[2]int
	fmt.Println(cityArray2)                          //[北京 上海 深圳]
	fmt.Printf("type of cityArray:%T\n", cityArray2) //type of cityArray:[3]string

	fmt.Println("方法三：我们还可以使用指定索引值的方式来初始化数组")
	a := [...]int{1: 1, 3: 5}
	fmt.Println(a)                  // [0 1 0 5]
	fmt.Printf("type of a:%T\n", a) //type of a:[4]int

	//数组的遍历
	var a2 = [...]string{"北京", "上海", "深圳"}
	fmt.Println("方法1：for循环遍历")
	for i := 0; i < len(a2); i++ {
		fmt.Println(a2[i])
	}

	fmt.Println("方法2：for range遍历（推荐）")
	for index, value := range a2 {
		fmt.Println(index, value)
	}

	/*
		多维数组。Go语言是支持多维数组的，我们这里以二维数组为例（数组中又嵌套数组）。
		多维数组只有第一层可以使用...来让编译器推导数组长度。例如：

		//支持的写法
		a := [...][2]string{
			{"北京", "上海"},
			{"广州", "深圳"},
			{"成都", "重庆"},
		}

		//不支持多维数组的内层使用...
		b := [3][...]string{
			{"北京", "上海"},
			{"广州", "深圳"},
			{"成都", "重庆"},
		}
	*/
	fmt.Println("二维数组的定义")
	a3 := [3][2]string{
		{"北京", "上海"},
		{"广州", "深圳"},
		{"成都", "重庆"},
	}
	fmt.Println(a3)       //[[北京 上海] [广州 深圳] [成都 重庆]]
	fmt.Println(a3[2][1]) //支持索引取值:重庆

	fmt.Println("二维数组的遍历")
	for _, v1 := range a3 {
		//fmt.Printf("%s\t", v1)
		for _, v2 := range v1 {
			/*if v2 == "北京" {
				continue
			}*/
			fmt.Printf("%s\t", v2)
		}
		fmt.Println()
	}

	//数组是值类型，赋值和传参会复制整个数组。因此改变副本的值，不会改变本身的值。
	arrayDemo()

	//如果直接赋值，不传给其它函数，是会直接修改数组对应下标的值
	a4 := [...]int{1, 2, 3}
	a4[1] = 20
	fmt.Println("如果直接赋值，不传给其它函数，是会直接修改数组对应下标的值", a4)

	fmt.Println("============ 数组 ============")
}

func modifyArrayA(x [3]int) {
	x[0] = 100
	fmt.Println("modifyArrayA中的数组a的值为：", x) //[100 20 30]
}

func modifyArrayB(x [3][2]int) {
	x[2][0] = 100
	fmt.Println("modifyArrayB中的数组b的值为：", x) //[[1 1] [1 1] [100 1]]
}

/*
1、数组支持 “=="、”!=" 操作符，因为内存总是被初始化过的。
2、[n]*T表示指针数组，*[n]T表示数组指针 。
一、指针数组适用于需要存储多个独立指针的场景，每个指针可以指向不同的内存地址。
二、数组指针适用于需要传递整个固定长度数组的场景，避免复制数组内容的开销。
*/
func arrayDemo() {
	fmt.Println("数组是值类型，赋值和传参会复制整个数组。因此改变副本的值，不会改变本身的值。")
	a := [3]int{10, 20, 30}
	modifyArrayA(a) //在modify中修改的是a的副本x
	fmt.Println(a)  //[10 20 30]
	b := [3][2]int{
		{1, 1},
		{1, 1},
		{1, 1},
	}
	modifyArrayB(b) //在modify中修改的是b的副本x
	fmt.Println(b)  //[[1 1] [1 1] [1 1]]
}
