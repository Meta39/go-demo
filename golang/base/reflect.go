package main

import (
	"errors"
	"fmt"
	"golang/base/structs"
	"reflect"
	"sync"
	"sync/atomic"
)

/*
Reflect 反射
反射是指在程序运行期间对程序本身进行访问和修改的能力。
程序在编译时，变量被转换为内存地址，变量名不会被编译器写入到可执行部分。
在运行程序时，程序无法获取自身的信息。
支持反射的语言可以在程序编译期间将变量的反射信息，如字段名称、类型信息、结构体信息等整合到可执行文件中，并给程序提供接口访问反射信息。
这样就可以在程序运行期间获取类型的反射信息，并且有能力修改它们。
Go程序在运行期间使用reflect包访问程序的反射信息。
空接口可以存储任意类型的变量，那我们如何知道这个空接口保存的数据是什么呢？反射就是在运行时动态的获取一个变量的类型信息和值信息。

reflect包
在Go语言的反射机制中，任何接口值都由是一个具体类型和具体类型的值两部分组成的
在Go语言中反射的相关功能由内置的reflect包提供，任意接口值在反射中都可以理解为由reflect.Type和reflect.Value两部分组成；
并且reflect包提供了reflect.TypeOf和reflect.ValueOf两个函数来获取任意对象的Value和Type。

反射是把双刃剑
反射是一个强大并富有表现力的工具，能让我们写出更灵活的代码。但是反射不应该被滥用，原因有以下三个。
一、基于反射的代码是极其脆弱的，反射中的类型错误会在真正运行的时候才会引发panic，那很可能是在代码写完的很长时间之后。
二、大量使用反射的代码通常难以理解。
三、反射的性能低下，基于反射实现的代码通常比正常代码运行速度慢一到两个数量级。
PS：一般通过反射处理过的结构体应该缓存起来，如果下次还使用它的话，可以直接从缓存取，这样就可以避免重复进行反射操作，影响性能。
*/
var (
	userCache         = structs.NewConcurrentHashMap() //缓存反射实例化的User结构体
	userCache2        sync.Map                         //缓存反射实例化的User结构体第2种做法
	countNewUserTimes int32                            // 必须用 int32/int64 计算同一个User数据是否重复使用反射实例化
)

func Reflect() {
	fmt.Println("============ 反射 ============")
	/*
		TypeOf
		在Go语言中，使用reflect.TypeOf()函数可以获得任意值的类型对象（reflect.Type），程序通过类型对象可以访问任意值的类型信息。
	*/
	var a float32 = 3.14
	reflectType(a) // type:float32,value:3.14
	var b int64 = 100
	reflectType(b) // type:int64,value:100

	/*
		type name 和 type kind
		在反射中关于类型还划分为两种：类型（Type）和种类（Kind）。
		因为在Go语言中我们可以使用type关键字构造很多自定义类型，而种类（Kind）就是指底层的类型。
		但在反射中，当需要区分指针、结构体等大品种的类型时，就会用到种类（Kind）。
		Go语言的反射中像数组、切片、Map、指针等类型的变量，它们的.Name()都是返回空。
		在reflect包中定义的type.go Kind类型如下：
		type Kind uint
		const (
			Invalid Kind = iota  // 非法类型
			Bool                 // 布尔型
			Int                  // 有符号整型
			Int8                 // 有符号8位整型
			Int16                // 有符号16位整型
			Int32                // 有符号32位整型
			Int64                // 有符号64位整型
			Uint                 // 无符号整型
			Uint8                // 无符号8位整型
			Uint16               // 无符号16位整型
			Uint32               // 无符号32位整型
			Uint64               // 无符号64位整型
			Uintptr              // 指针
			Float32              // 单精度浮点数
			Float64              // 双精度浮点数
			Complex64            // 64位复数类型
			Complex128           // 128位复数类型
			Array                // 数组
			Chan                 // 通道
			Func                 // 函数
			Interface            // 接口
			Map                  // 映射
			Ptr                  // 指针
			Slice                // 切片
			String               // 字符串
			Struct               // 结构体
			UnsafePointer        // 底层指针
		)
		举个例子，我们定义了两个指针类型和两个结构体类型，通过反射查看它们的类型和种类。
	*/
	var a2 *float32 // 指针
	var b2 MyInt    // 自定义类型
	var c2 rune     // 类型别名
	var d2 = person{
		name: "沙河小王子",
		age:  18,
	}
	var e2 = book{title: "《跟小王子学Go语言》"}
	reflectType2(a2) // type:,kind:ptr,value:<nil>
	reflectType2(b2) // type:int,kind:int,value:0
	reflectType2(c2) // type:int32,kind:int32,value:0
	reflectType2(d2) // type:person,kind:struct,value:{沙河小王子 18}
	reflectType2(e2) // type:book,kind:struct,value:{《跟小王子学Go语言》}

	/*
		ValueOf
		reflect.ValueOf()返回的是reflect.Value类型，其中包含了原始值的值信息。reflect.Value与原始值之间可以互相转换。
		reflect.Value类型提供的获取原始值的方法如下：
		方法							说明
		Interface() interface {}	将值以 interface{} 类型返回，可以通过类型断言转换为指定类型
		Int() int64					将值以 int 类型返回，所有有符号整型均可以此方式返回
		Uint() uint64				将值以 uint 类型返回，所有无符号整型均可以此方式返回
		Float() float64				将值以双精度（float64）类型返回，所有浮点数（float32、float64）均可以此方式返回
		Bool() bool					将值以 bool 类型返回
		Bytes() []bytes				将值以字节数组 []bytes 类型返回
		String() string				将值以字符串类型返回
	*/
	fmt.Printf("通过反射获取值")
	var a3 float32 = 3.14
	var b3 int64 = 100
	reflectValue(a3) // type is float32, value is 3.140000
	reflectValue(b3) // type is int64, value is 100
	// 将int类型的原始值转换为reflect.Value类型
	c3 := reflect.ValueOf(10)
	fmt.Printf("c3 type:%T,kind:%v\n", c3, c3.Kind()) // c3 type:reflect.Value,kind:int

	/*
		通过反射设置变量的值
		想要在函数中通过反射修改变量的值，需要注意函数参数传递的是值拷贝，必须传递变量地址才能修改变量值。
		而反射中使用专有的Elem()方法来获取指针对应的值。
	*/
	fmt.Println("通过反射设置变量的值")
	var a4 int64 = 100
	//reflectSetValue1(a4) //方法错误，传值也错误 panic: reflect: reflect.Value.SetInt using unaddressable value
	reflectSetValue2(&a4) //传入的是指针地址
	fmt.Println("原：100 => 修改后：200", a4)

	/*
		isNil()和isValid()
		IsNil()报告v持有的值是否为nil。v持有的值的分类必须是通道、函数、接口、映射、指针、切片之一；否则IsNil函数会导致panic。
		func (v Value) IsNil() bool
		IsValid()返回v是否持有一个值。如果v是Value零值会返回假，此时v除了IsValid、String、Kind之外的方法都会导致panic。
		func (v Value) IsValid() bool
		举个例子
		IsNil()常被用于判断指针是否为空；IsValid()常被用于判定返回值是否有效。
	*/
	// *int类型空指针
	var a5 *int
	fmt.Println("var a5 *int IsNil:", reflect.ValueOf(a5).IsNil())
	// nil值
	fmt.Println("nil IsValid:", reflect.ValueOf(nil).IsValid())
	// 实例化一个匿名结构体
	b5 := struct{}{}
	// 尝试从结构体中查找"abc"字段
	fmt.Println("不存在的结构体成员:", reflect.ValueOf(b5).FieldByName("abc").IsValid())
	// 尝试从结构体中查找"abc"方法
	fmt.Println("不存在的结构体方法:", reflect.ValueOf(b5).MethodByName("abc").IsValid())
	// map
	c5 := map[string]int{}
	// 尝试从map中查找一个不存在的键
	fmt.Println("map中不存在的键：", reflect.ValueOf(c5).MapIndex(reflect.ValueOf("娜扎")).IsValid())

	/*
		结构体反射
		任意值通过reflect.TypeOf()获得反射对象信息后，如果它的类型是结构体，可以通过反射值对象（reflect.Type）的NumField()和Field()方法获得结构体成员的详细信息。
		reflect.Type中与获取结构体成员相关的的方法如下表所示。
		方法																说明
		Field(i int) StructField										根据索引，返回索引对应的结构体字段的信息。
		NumField() int													返回结构体成员字段数量。
		FieldByName(name string) (StructField, bool)					根据给定字符串返回字符串对应的结构体字段的信息。
		FieldByIndex(index []int) StructField							多层成员访问时，根据 []int 提供的每个结构体的字段索引，返回字段的信息。
		FieldByNameFunc(match func(string) bool) (StructField,bool)		根据传入的匹配函数匹配需要的字段。
		NumMethod() int													返回该类型的方法集中方法的数目
		Method(int) Method												返回该类型方法集中的第i个方法
		MethodByName(string)(Method, bool)								根据方法名返回该类型方法集中的方法

		StructField类型用来描述结构体中的一个字段的信息。
		StructField的定义如下：
		type StructField struct {
			// Name是字段的名字。PkgPath是非导出字段的包路径，对导出字段该字段为""。
			// 参见http://golang.org/ref/spec#Uniqueness_of_identifiers
			Name    string
			PkgPath string
			Type      Type      // 字段的类型
			Tag       StructTag // 字段的标签
			Offset    uintptr   // 字段在结构体中的字节偏移量
			Index     []int     // 用于Type.FieldByIndex时的索引切片
			Anonymous bool      // 是否匿名字段
		}
	*/
	//结构体反射示例。当我们使用反射得到一个结构体数据之后可以通过索引依次获取其字段信息，也可以通过字段名去获取指定的字段信息。
	stu1 := structs.Student{
		ID:     1,
		Name:   "小王子",
		Gender: "男",
	}

	t := reflect.TypeOf(stu1)
	fmt.Println(t.Name(), t.Kind()) // Student struct
	// 通过for循环遍历结构体的所有字段信息
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fmt.Printf("name:%s index:%d type:%v json tag:%v\n", field.Name, field.Index, field.Type, field.Tag.Get("json"))
	}

	// 通过字段名获取指定结构体字段信息
	if genderField, ok := t.FieldByName("Gender"); ok {
		fmt.Printf("通过字段名获取指定结构体字段信息 name:%s index:%d type:%v json tag:%v\n", genderField.Name, genderField.Index, genderField.Type, genderField.Tag.Get("json"))
	}

	printMethod(d2)

	//通过反射实例化结构体，并且如果是第一次创建，则放入缓存，第二次发现重复的结构体，直接从缓存取。提升性能。
	user := newPerson("Meta39", "男", "广东省", "广州市")
	fmt.Printf("user %+v\n", user)

	user2 := newPerson("Meta39", "男", "广东省", "广州市")
	fmt.Printf("user2 %+v\n", user2)

	fmt.Printf("相同结构体创建structs.User次数为%v\n", countNewUserTimesGet())
	if user3, ok := userCache.Get("Meta39男广东省广州市"); ok {
		fmt.Printf("从缓存userCache读取structs.User %+v\n", user3)
	}

	countNewUserTimesZero()
	fmt.Println("countNewUserTimesGet原子归零。countNewUserTimes:", countNewUserTimes)

	user4 := newPerson2("Meta39", "男", "广东省", "广州市")
	fmt.Printf("user4 %+v\n", user4)

	user5 := newPerson2("Meta39", "男", "广东省", "广州市")
	fmt.Printf("user5 %+v\n", user5)
	fmt.Printf("newPerson2相同结构体创建structs.User次数为%v\n", countNewUserTimesGet())
	if user6, ok2 := userCache2.Load("Meta39男广东省广州市"); ok2 {
		fmt.Printf("从缓存userCache2读取structs.User %+v\n", user6.(*LazyValue).value.(*structs.User))
	}

	fmt.Println("============ 反射 ============")
}

type person struct {
	name string
	age  int
}

func (s person) Study() string {
	msg := "好好学习，天天向上。"
	fmt.Println(msg)
	return msg
}

func (s person) Sleep() string {
	msg := "好好睡觉，快快长大。"
	fmt.Println(msg)
	return msg
}

type book struct {
	title string
}

// 根据传入的变量值，通过反射获得变量的类型对象
func reflectType(x interface{}) {
	t := reflect.TypeOf(x)
	fmt.Printf("根据传入的变量，通过反射获取变量的类型 type:%v,value:%v\n", t, x)
}

// 根据传入的变量值，通过反射获得变量的类型对象（类型、种类）
func reflectType2(x interface{}) {
	t := reflect.TypeOf(x)
	fmt.Printf("type:%v,kind:%v,value:%v\n", t.Name(), t.Kind(), x)
}

// 根据传入的变量值，通过反射强转为其它类型
func reflectValue(x interface{}) {
	v := reflect.ValueOf(x)
	k := v.Kind()
	switch k {
	case reflect.Int64:
		// v.Int()从反射中获取整型的原始值，然后通过int64()强制类型转换
		fmt.Printf("type is int64, value is %d\n", v.Int())
	case reflect.Float32:
		// v.Float()从反射中获取浮点型的原始值，然后通过float32()强制类型转换
		fmt.Printf("type is float32, value is %f\n", float32(v.Float()))
	case reflect.Float64:
		// v.Float()从反射中获取浮点型的原始值，然后通过float64()强制类型转换
		fmt.Printf("type is float64, value is %f\n", v.Float())
	default:
		err := errors.New("无法识别的类型或种类")
		fmt.Println(err)
	}
}

// 使用反射的修改值（错误用法）
func reflectSetValue1(x interface{}) {
	v := reflect.ValueOf(x)
	if v.Kind() == reflect.Int64 {
		v.SetInt(200) //修改的是副本，reflect包会引发panic
	}
}

// 使用反射的修改值（正确用法）
func reflectSetValue2(x interface{}) { //传入的x是指针，即：&x
	v := reflect.ValueOf(x)
	// 反射中使用 Elem()方法获取指针对应的值
	if v.Elem().Kind() == reflect.Int64 {
		v.Elem().SetInt(200)
	}
}

// 通过反射获取结构体里的方法并调用其中的方法
func printMethod(x interface{}) {
	t := reflect.TypeOf(x)
	v := reflect.ValueOf(x)

	fmt.Printf("通过返回获取到结构体包含%v个方法\n", t.NumMethod())
	for i := 0; i < v.NumMethod(); i++ {
		methodType := v.Method(i).Type()
		fmt.Printf("method name:%s\n", t.Method(i).Name)
		fmt.Printf("method:%s\n", methodType)
		// 通过反射调用方法传递的参数必须是 []reflect.Value 类型
		var args []reflect.Value
		v.Method(i).Call(args)
	}
}

// 通过反射实例化结构体，并放入缓存
func newPerson(name string, gender string, province string, city string) *structs.User {
	//先判断缓存 userCache 是否存在，如果存在，直接取，不存在，则放入缓存。
	// 并发安全地获取或计算值
	key := name + gender + province + city
	val := userCache.ComputeIfAbsent(key, func() any {
		countNewUserTimesAdd() //计数器 + 1

		employeeType := reflect.TypeOf(structs.User{})
		employeePtr := reflect.New(employeeType)
		employeeValue := employeePtr.Elem()

		// 设置嵌套字段
		employeeValue.FieldByName("Name").SetString(name)
		employeeValue.FieldByName("Gender").SetString(gender)
		address := employeeValue.FieldByName("Address")
		address.FieldByName("Province").SetString(province)
		address.FieldByName("City").SetString(city)

		//realEmployee := employeePtr.Interface().(*structs.User)
		//return realEmployee
		return employeePtr.Interface()
	})
	return val.(*structs.User)
}

/*
通过反射实例化结构体，并放入缓存。sync.Map 适合读多写少的场景，频繁写入时性能可能不如 sync.RWMutex + map。
显然我们就是这种读多写少的情况。
*/
func newPerson2(name string, gender string, province string, city string) *structs.User {
	//先判断缓存 userCache 是否存在，如果存在，直接取，不存在，则放入缓存。
	key := name + gender + province + city

	val, _ := userCache2.LoadOrStore(key, &LazyValue{})
	computedVal := val.(*LazyValue).Get(func() interface{} {
		countNewUserTimesAdd() //计数器 + 1

		employeeType := reflect.TypeOf(structs.User{})
		employeePtr := reflect.New(employeeType)
		employeeValue := employeePtr.Elem()

		// 设置嵌套字段
		employeeValue.FieldByName("Name").SetString(name)
		employeeValue.FieldByName("Gender").SetString(gender)
		address := employeeValue.FieldByName("Address")
		address.FieldByName("Province").SetString(province)
		address.FieldByName("City").SetString(city)

		//realEmployee := employeePtr.Interface().(*structs.User)
		//return realEmployee
		return employeePtr.Interface()
	})
	return computedVal.(*structs.User)
}

// 原子归零
func countNewUserTimesZero() {
	atomic.StoreInt32(&countNewUserTimes, 0) // 将 counter 重置为 0
}

func countNewUserTimesAdd() {
	atomic.AddInt32(&countNewUserTimes, 1) // 原子性 +1
}

func countNewUserTimesGet() int32 {
	return atomic.LoadInt32(&countNewUserTimes) // 原子性读取
}

type LazyValue struct {
	once  sync.Once
	value interface{}
}

func (lv *LazyValue) Get(computeFunc func() interface{}) interface{} {
	lv.once.Do(func() {
		lv.value = computeFunc()
	})
	return lv.value
}
