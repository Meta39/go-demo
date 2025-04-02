package main

/*
main 函数的package一定是main包，否则无法执行
*/
func main() {
	//1、调用在keywords包下的Keywords函数里输出Go语言中25个关键字
	Keywords()
	//2、变量
	Variable()
	//3、常量
	Const()
	//4、基本数据类型
	BasicDataTypes()
	//5、运算符
	Operator()
	//6、流程控制
	ProcessControl()
	//7、数组【var 数组变量名 [元素数量]T】（因为数组的长度是固定的，不够灵活，因此在大多数情况下，推荐使用切片。切片提供了更灵活和强大的功能，是处理动态数据的首选方式。）
	Array()
	//8、切片【var 切片变量名 []T】（可以理解为自动扩容数组，即：不设置元素数量的数组就是切片。）
	Slice() //使用make函数构造切片。make([]T, size切片中元素的数量, cap切片的容量)
	//9、map【make(map[KeyType]ValueType, [cap])】
	Map()
	//10、函数
	Func()
	//11、指针、new和make关键字的作用
	Pointer()
	//12、结构体
	Struct()
	//13、包
	Package()
	//14、接口
	Interface()
	//15、错误
	Error()
	//16、反射
	Reflect()
	//17、并发
	Concurrent()
	//18、泛型
	Any()
}
