package main

import (
	"fmt"
	"sort"
)

/*
Slice 切片（可以理解为自动扩容数组，即：不设置元素数量的数组就是切片。）
1、切片（Slice）是一个拥有相同类型元素的可变长度的序列。它是基于数组类型做的一层封装。它非常灵活，支持自动扩容。
2、切片是一个引用类型，它的内部结构包含地址、长度和容量。切片一般用于快速地操作一块数据集合。
3、切片的长度和容量。切片拥有自己的长度和容量，我们可以通过使用内置的len()函数求长度，使用内置的cap()函数求切片的容量。
4、切片表达式。切片表达式从字符串、数组、指向数组或切片的指针构造子字符串或切片。它有两种变体：一种指定low和high两个索引界限值的简单的形式，另一种是除了low和high索引界限值外还指定容量的完整的形式。
5、切片的本质就是对底层数组的封装，它包含了三个信息：底层数组的指针、切片的长度（len）和切片的容量（cap）。
6、判断切片是否为空。请始终使用len(s) == 0来判断，而不应该使用s == nil来判断。

7、切片的定义：
var name []T
其中，
name:表示变量名
T:表示切片中的元素类型
PS：数组的定义是：var name [元素数量]T，即：不设置元素数量的数组就是切片。

使用make()函数构造切片
make([]T, size, cap)
其中，
T:切片的元素类型
size:切片中元素的数量
cap:切片的容量

为什么使用推荐使用切片，而不是数组？
数组的长度是固定的并且数组长度属于类型的一部分，所以数组有很多的局限性。 例如：

	func arraySum(x [3]int) int{
	    sum := 0
	    for _, v := range x{
	        sum = sum + v
	    }
	    return sum
	}

这个求和函数只能接受[3]int类型，其他的都不支持。 再比如，
a := [3]int{1, 2, 3}
数组a中已经有三个元素了，我们不能再继续往数组a中添加新元素了。

8、切片的扩容策略：
可以通过查看$GOROOT/src/runtime/slice.go源码，其中扩容相关代码如下：

	newcap := old.cap
	doublecap := newcap + newcap

	if cap > doublecap {
		newcap = cap
	} else {

		if old.len < 1024 {
			newcap = doublecap
		} else {
			// Check 0 < newcap to detect overflow
			// and prevent an infinite loop.
			for 0 < newcap && newcap < cap {
				newcap += newcap / 4
			}
			// Set newcap to the requested cap when
			// the newcap calculation overflowed.
			if newcap <= 0 {
				newcap = cap
			}
		}
	}

	从上面的代码可以看出以下内容：
	一、首先判断，如果新申请容量（cap）大于2倍的旧容量（old.cap），最终容量（newcap）就是新申请的容量（cap）。
	二、否则判断，如果旧切片的长度小于1024，则最终容量(newcap)就是旧容量(old.cap)的两倍，即（newcap=doublecap），
	三、否则判断，如果旧切片长度大于等于1024，则最终容量（newcap）从旧容量（old.cap）开始循环增加原来的1/4，即（newcap=old.cap,for {newcap += newcap/4}）直到最终容量（newcap）大于等于新申请的容量(cap)，即（newcap >= cap）
	四、如果最终容量（cap）计算值溢出，则最终容量（cap）就是新申请容量（cap）。
	需要注意的是，切片扩容还会根据切片中元素的类型不同而做不同的处理，比如int和string类型的处理方式就不一样。
*/
func Slice() {
	fmt.Println("============ 切片 ============")

	// 声明切片类型
	var a []string              //声明一个字符串切片
	var b = []int{}             //声明一个整型切片并初始化
	var c = []bool{false, true} //声明一个布尔切片并初始化
	var d = []bool{false, true} //声明一个布尔切片并初始化
	fmt.Println(a)              //[]
	fmt.Println(b)              //[]
	fmt.Println(c)              //[false true]
	fmt.Println(d)              //[false true]
	fmt.Println(a == nil)       //true
	fmt.Println(b == nil)       //false
	fmt.Println(c == nil)       //false
	//fmt.Println(c == d)   //切片是引用类型，不支持直接比较，只能和nil比较

	/*
		简单切片表达式
		切片的底层就是一个数组，所以我们可以基于数组通过切片表达式得到切片。
		切片表达式中的low和high表示一个索引范围（左包含，右不包含），也就是下面代码中从数组a中选出1<=索引值<4的元素组成切片s，得到的切片长度=high-low，容量等于得到的切片的底层数组的容量。
		为了方便起见，可以省略切片表达式中的任何索引。省略了low则默认为0；省略了high则默认为切片操作数的长度:
		a[2:]  // 等同于 a[2:len(a)]
		a[:3]  // 等同于 a[0:3]
		a[:]   // 等同于 a[0:len(a)]
	*/
	fmt.Println("简单切片表达式")
	a2 := [6]int{5, 4, 3, 2, 1, 0}
	fmt.Println("a2", a2)
	//即：获取数组下标[1,3)中的元素
	s := a2[1:3] // s := a[low:high]
	fmt.Printf("a2[1:3]表示获取数组a2下标[1,3)中的元素。s:%v len(s):%v cap(s):%v\n", s, len(s), cap(s))
	/*
		对于数组或字符串，如果0 <= low <= high <= len(a)，则索引合法，否则就会索引越界（out of range）。
		对切片再执行切片表达式时（切片再切片），high的上限边界是切片的容量cap(a)，而不是长度。
		常量索引必须是非负的，并且可以用int类型的值表示;对于数组或常量字符串，常量索引也必须在有效范围内。
		如果low和high两个指标都是常数，它们必须满足low <= high。
		如果索引在运行时超出范围，就会发生运行时panic。
	*/
	fmt.Println("对切片再执行切片表达式时（切片再切片），high的上限边界是切片的容量cap(a3)，而不是长度。")
	a3 := [5]int{1, 2, 3, 4, 5}
	fmt.Printf("a3:%v len(a3):%v cap(a3):%v\n", a3, len(a3), cap(a3))
	//s3 := a3[1:3] 表示切掉数组a3从下标0开始的一个元素得到s3，即切掉：1。保留 3 - 1 = 2 个元素，即：长度(2)：{2,3}，容量(4)：{2, 3, 4, 5}
	s3 := a3[1:3] // s3 := a3[low:high]
	fmt.Printf("s3对a3进行a3[1:3]切片。s3:%v len(s3):%v cap(s3):%v\n", s3, len(s3), cap(s3))
	fmt.Println("s3 := a3[1:3] 表示切掉数组a3从下标0开始的一个元素得到s3，即切掉：1。保留 3 - 1 = 2 个元素，即：长度(2)：{2,3}，容量(4)：{2, 3, 4, 5}")
	//s4 := s3[3:4] 切掉s3切片从0开始的三个元素，即切掉：2, 3, 4。保留 4 - 3 = 1 个元素，即：长度(1)：{5}，容量(1)：{5}
	s4 := s3[3:4] // 索引的上限是cap(s3)，即：4。而不是len(s3)，即：2
	fmt.Printf("s4对s3再次进行s3[3:4]切片。s4:%v len(s4):%v cap(s4):%v\n", s4, len(s4), cap(s4))
	fmt.Println("s4 := s3[3:4] 切掉s3切片从0开始的三个元素，即切掉：2, 3, 4。保留 4 - 3 = 1 个元素，即：长度(1)：{5}，容量(1)：{5}")

	/*
		完整切片表达式
		对于数组，指向数组的指针，或切片a(注意不能是字符串)支持完整切片表达式：
		a4[low : high : max]
		上面的代码会构造与简单切片表达式a3[low: high]相同类型、相同长度和元素的切片。
		另外，它会将得到的结果切片的容量设置为 max - low 。在完整切片表达式中只有第一个索引值（low）可以省略；它默认为0。
		完整切片表达式需要满足的条件是0 <= low <= high <= max <= cap(a4)，其他条件和简单切片表达式相同。
	*/
	a4 := [5]int{1, 2, 3, 4, 5}
	//切片：长度 = high - low，容量 = max - low。
	t := a4[1:3:5]                                              //low = 1，high = 3，max = 5。得到切片：{2, 3}，长度【high - low】(3 - 1 = 2)：2，容量【max - low】(5 - 1 = 4)：4
	fmt.Printf("t:%v len(t):%v cap(t):%v\n", t, len(t), cap(t)) //t:[2 3] len(t):2 cap(t):4
	fmt.Println("t := a4[1:3:5]，low = 1，high = 3，max = 5。得到切片：{2, 3}，长度【high - low】(3 - 1 = 2)：2，容量【max - low】(5 - 1 = 4)：4")

	/*
		使用make()函数构造切片
		我们上面都是基于数组来创建的切片，如果需要动态的创建一个切片，我们就需要使用内置的make()函数，格式如下：
		make([]T, size, cap)
		其中：
		T:切片的元素类型
		size:切片中元素的数量
		cap:切片的容量
	*/
	a5 := make([]int, 2, 10)
	fmt.Println("使用make()函数构造切片：格式：make([]T, size, cap)")
	fmt.Printf("a5 := make([]int, 2, 10) 其中：2是长度，10是容量。a5:%v len(a5):%v cap(a5):%v\n", a5, len(a5), cap(a5)) //a5:[0 0] len(a5):2 cap(a5):10
	//上面代码中a5的内部存储空间已经分配了10个，但实际上只用了2个。 容量并不会影响当前元素的个数，所以len(a5)返回2，cap(a5)则返回该切片的容量。

	//要检查切片是否为空，请始终使用len(s) == 0来判断，而不应该使用s == nil来判断。
	isEmpty := len(a5) == 0
	fmt.Printf("要检查切片是否为空，请始终使用len(s) == 0来判断，而不应该使用s == nil来判断。a5 = %v， len(a5) = %v ，判空：len(a5) == 0 ? %v\n", a5, len(a5), isEmpty)

	/*
		切片之间是不能比较的，我们不能使用==操作符来判断两个切片是否含有全部相等元素。
		切片唯一合法的比较操作是和nil比较。
		一个nil值的切片并没有底层数组，一个nil值的切片的长度和容量都是0。
		但是我们不能说一个长度和容量都是0的切片一定是nil
		所以要判断一个切片是否是空的，要是用len(s) == 0来判断，不应该使用s == nil来判断。
	*/

	/*
		切片的赋值拷贝
		下面的代码中演示了拷贝前后两个变量共享底层数组，对一个切片的修改会影响另一个切片的内容，这点需要特别注意。
	*/
	s1 := make([]int, 3) //[0 0 0]
	s2 := s1             //将s1直接赋值给s2，s1和s2共用一个底层数组
	fmt.Println("未赋值前的s1", s1)
	s2[0] = 100
	fmt.Println("拷贝前后两个变量共享底层数组，对一个切片的修改会影响另一个切片的内容，这点需要特别注意。如果需要独立s1和s2的修改，则需要使用copy()函数，这样就不会造成修改同一个共享底层数组。")
	fmt.Println("s1 = ", s1)                                 //[100 0 0]
	fmt.Println("s2[0] = 100 后s1和s2都会改变，因为共享底层数组，s2 = ", s2) //[100 0 0]

	//切片遍历方式和数组是一致的，支持索引遍历和for range遍历。
	s5 := []int{1, 3, 5}
	forSlice(s5)
	forRangeSlice(s5)

	/*
		append()方法为切片添加元素
		Go语言的内建函数append()可以为切片动态添加元素。
		可以一次添加一个元素，可以添加多个元素，也可以添加另一个切片中的元素（后面加…）。
	*/
	var s6 []int
	s6 = append(s6, 1)       // [1]
	s6 = append(s6, 2, 3, 4) // [1 2 3 4]
	s7 := []int{5, 6, 7}     // [5, 6, 7]
	s6 = append(s6, s7...)   // [1 2 3 4 5 6 7]
	fmt.Printf("s6 = %v，s7 = %v\n", s6, s7)

	/*
		通过var声明的零值切片可以在append()函数直接使用，无需初始化。
		正确示例如下：
		var s []int
		s = append(s, 1, 2, 3)

		错误示例如下：
		s := []int{}  // 没有必要初始化
		s = append(s, 1, 2, 3)

		var s = make([]int)  // 没有必要初始化
		s = append(s, 1, 2, 3)
	*/

	/*
		每个切片会指向一个底层数组，这个数组的容量够用就添加新增元素。
		当底层数组不能容纳新增的元素时，切片就会自动按照一定的策略进行“扩容”，此时该切片指向的底层数组就会更换。
		“扩容”操作往往发生在append()函数调用时，所以我们通常都需要用原变量接收append函数的返回值。
	*/
	//追加一个元素
	var numSlice []int
	appendSlice(numSlice)
	//追加单个元素、多个元素、切片和扩容
	var citySlice []string
	appendSlices(citySlice)

	//使用copy()函数复制切片。copy(destSlice, srcSlice []T)，其中：destSlice: 目标切片【结果】，srcSlice: 数据来源切片【需要复制的切片】
	copySlice()

	/*
		从切片中删除元素
		Go语言中并没有删除切片元素的专用方法，我们可以使用切片本身的特性来删除元素。
		总结一下就是：要从切片a中删除索引为index的元素，操作方法是a = append(a[:index], a[index+1:]...)
	*/
	removeSlice()

	//切片排序
	sortSlice()

	fmt.Println("============ 切片 ============")
}

// for 遍历切片
func forSlice(s []int) {
	fmt.Println("for 遍历 slice 切片")
	for i := 0; i < len(s); i++ {
		fmt.Println(i, s[i])
	}
}

// for range 遍历切片
func forRangeSlice(s []int) {
	fmt.Println("for range 遍历 slice 切片")
	for index, value := range s {
		fmt.Println(index, value)
	}
}

// append()添加元素和切片扩容
func appendSlice(numSlice []int) {
	fmt.Println("append()添加元素和切片扩容")
	for i := 0; i < 10; i++ {
		numSlice = append(numSlice, i)
		fmt.Printf("%v  len:%d  cap:%d  ptr:%p\n", numSlice, len(numSlice), cap(numSlice), numSlice)
	}
	/*
		从上面的结果可以看出：
		1、append()函数将元素追加到切片的最后并返回该切片。
		2、切片numSlice的容量按照1，2，4，8，16这样的规则自动进行扩容，每次扩容后都是扩容前的2倍。
	*/
	fmt.Println(`从上面的结果可以看出：
1、append()函数将元素追加到切片的最后并返回该切片。
2、切片numSlice的容量按照1，2，4，8，16这样的规则自动进行扩容，每次扩容后都是扩容前的2倍。`)
}

// append()函数一次性追加多个元素和扩容
func appendSlices(citySlice []string) {
	// 追加一个元素
	citySlice = append(citySlice, "北京")
	// 追加多个元素
	citySlice = append(citySlice, "上海", "广州", "深圳")
	// 追加切片
	a := []string{"成都", "重庆"}
	citySlice = append(citySlice, a...)
	fmt.Println("append()函数一次性追加多个元素和扩容", citySlice) //[北京 上海 广州 深圳 成都 重庆]
}

// 使用copy()函数复制切片，独立各自的共享底层数组。
func copySlice() {
	//不使用copy()函数，修改共享底层数组，底层数组会收到影响。（一般也不会定义a，然后又把a赋给b这么干~一般都是取出来直接用。）
	a := []int{1, 2, 3, 4, 5}
	b := a
	fmt.Println("不使用copy()函数，影响共享底层数组。由于切片是引用类型，所以a和b其实都指向了同一块内存地址。修改b的同时a的值也会发生变化。")
	fmt.Println("修改元素前的a = ", a) //[1 2 3 4 5]
	fmt.Println("修改元素前的b = ", b) //[1 2 3 4 5]
	b[0] = 1000
	fmt.Println("修改元素后的a = ", a) //[1000 2 3 4 5]
	fmt.Println("修改元素后的b = ", b) //[1000 2 3 4 5]

	fmt.Println("使用copy()函数，独立各自的共享底层数组，互不影响。Go语言内建的copy()函数可以迅速地将一个切片的数据复制到另外一个切片空间中，copy()函数的使用格式如下：copy(destSlice, srcSlice []T)")
	// copy()复制切片
	c := []int{1, 2, 3, 4, 5}
	d := make([]int, len(c), cap(c))
	fmt.Println("使用copy()函数将切片c中的元素复制到切片d")
	copy(d, c)                   //使用copy()函数将切片c中的元素复制到切片d
	fmt.Println("修改元素前的c = ", c) //[1 2 3 4 5]
	fmt.Println("修改元素前的d = ", d) //[1 2 3 4 5]
	d[0] = 1000
	fmt.Println("修改d切片元素后的c = ", c) //[1 2 3 4 5]
	fmt.Println("修改d切片元素后的d = ", d) //[1000 2 3 4 5]
}

// 从切片中删除元素。总结一下就是：要从切片a中删除索引为index的元素，操作方法是a = append(a[:index], a[index+1:]...)
func removeSlice() {
	// 从切片中删除元素
	a := []int{30, 31, 32, 33, 34, 35, 36, 37}
	fmt.Printf("删除a = %v中的索引为2的元素，即：%v。\n", a, a[2])
	// 要删除索引为2的元素
	a = append(a[:2], a[3:]...)
	fmt.Println("删除a[2]后的切片a =", a) //[30 31 33 34 35 36 37]
}

// 切片排序
func sortSlice() {
	fmt.Println("切片排序：")
	a := []int{3, 7, 8, 9, 1}
	fmt.Println("排序前的a = ", a)
	//默认ASC排序
	sort.Ints(a)
	fmt.Println("ASC 排序后的a = ", a)
	//没有提供DESC排序的方法，因此可以通过比较Slice函数中调整逻辑 (a[i] > a[j]) 来实现降序排序。
	sort.Slice(a, func(i, j int) bool {
		return a[i] > a[j]
	})
	fmt.Println("DESC 排序后的a = ", a)
}
