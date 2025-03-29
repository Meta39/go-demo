package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

/*
Map Go语言中提供的映射关系容器为map，其内部使用散列表（hash）实现并。
1、map是一种无序的基于key-value的数据结构，Go语言中的map是引用类型，必须初始化才能使用。
2、遍历map时的元素顺序与添加键值对的顺序无关。

map定义
map[KeyType]ValueType
其中，
KeyType:表示键的类型。
ValueType:表示键对应的值的类型。

map类型的变量默认初始值为nil，需要使用make()函数来分配内存。语法为：
make(map[KeyType]ValueType, [cap])
其中cap表示map的容量，该参数虽然不是必须的，但是我们应该在初始化map的时候就为其指定一个合适的容量。这点跟java差不多的
*/
func Map() {
	fmt.Println("============ map ============")

	//map中的数据都是成对出现的，map的基本使用示例代码如下：
	scoreMap := make(map[string]int, 8)
	scoreMap["张三"] = 90
	scoreMap["小明"] = 100
	fmt.Println(scoreMap)
	fmt.Println(scoreMap["小明"])
	fmt.Printf("type of a:%T\n", scoreMap)

	//map在声明的时候填充元素，例如：
	userInfo := map[string]string{
		"username": "Meta39",
		"password": "123456",
	}
	fmt.Println("map在声明的时候填充元素", userInfo)

	//判断某个键是否存在。Go语言中有个判断map中键是否存在的特殊写法，格式：value, ok := map[key]
	// 如果key存在ok为true,v为对应的值；不存在ok为false,v为值类型的零值
	fmt.Println("判断某个键是否存在。如：判断scoreMap是否包含张三这个key？")
	v, ok := scoreMap["张三"]
	if ok {
		fmt.Printf("包含，值为%v\n", v)
	} else {
		fmt.Println("不包含")
	}

	/*
		delete()函数删除键值对。
		使用delete()内建函数从map中删除一组键值对，delete()函数的格式：delete(map, key)
		map:表示要删除键值对的map
		key:表示要删除的键值对的键
	*/
	delete(scoreMap, "小明") ////将小明:100从map中删除

	//遍历map获取key和value(遍历map时的元素顺序与添加键值对的顺序无关。)
	forMap(scoreMap)
	//遍历map只获取key(遍历map时的元素顺序与添加键值对的顺序无关。)
	forMapKey(scoreMap)
	//按key顺序进行遍历map
	forSortMap()

	//元素为map类型的切片，简单理解：类型为slice切片，切片里的元素为map
	sliceMap()
	// 值为切片类型的map，简单理解：变量是map，key为KeyType(即：任意类型都可以)，value为切片。
	mapSlice()

	fmt.Println("============ map ============")
}

// 遍历map获取key和value
func forMap(scoreMap map[string]int) {
	fmt.Println("map的遍历，获取key和value")
	for k, v := range scoreMap {
		fmt.Printf("key = %v, value = %v\n", k, v)
	}
}

// 遍历map只获取key
func forMapKey(scoreMap map[string]int) {
	fmt.Println("map的遍历，只获取key，不获取value")
	for k := range scoreMap {
		fmt.Printf("key = %v\n", k)
	}
}

// 按顺序遍历map（一般情况下不需要）
func forSortMap() {
	fmt.Println("按顺序遍历map")
	rand.New(rand.NewSource(time.Now().UnixNano())) //初始化随机数种子

	var scoreMap = make(map[string]int, 20)

	for i := 0; i < 10; i++ {
		key := fmt.Sprintf("stu%02d", i) //生成stu开头的字符串
		value := rand.Intn(100)          //生成0~99的随机整数
		scoreMap[key] = value
	}
	//取出map中的所有key存入切片keys
	var keys = make([]string, 0, len(scoreMap))
	for key := range scoreMap {
		keys = append(keys, key)
	}
	//对切片进行排序
	sort.Strings(keys)
	//按照排序后的key遍历map
	for _, key := range keys {
		fmt.Println(key, scoreMap[key])
	}
}

// 元素为map类型的切片，简单理解：类型为slice切片，切片里的元素为map
func sliceMap() {
	fmt.Println("元素为map类型的切片，简单理解：类型为slice切片，切片里的元素为map")
	var sliceMap = make([]map[string]string, 2) //2表示的是 sliceMap 切片的初始长度(len)，不是map的容量(cap)。如果写0，那么就是nil，sliceMap[0]会数组下标越界报错。
	fmt.Printf("sliceMap len():%v cap():%v\n", len(sliceMap), cap(sliceMap))
	for index, value := range sliceMap {
		fmt.Printf("index:%d value:%v\n", index, value)
	}
	fmt.Println("after init")
	// 对切片中的map元素进行初始化和赋值
	//sliceMap = append(sliceMap, make(map[string]string, 10))。如果var sliceMap = make([]map[string]string, 0)，则需要用这种方式去添加切片元素后才能使用sliceMap[0]
	sliceMap[0] = make(map[string]string, 10) //10是map的容量(cap)，而不是长度(len)
	tempMap := sliceMap[0]
	tempMap["name"] = "Meta39"
	tempMap["password"] = "123456"
	tempMap["address"] = "广州"
	tempMap["age"] = "18"
	for index, value := range sliceMap {
		fmt.Printf("index:%d value:%v\n", index, value)
	}
}

// 值为切片类型的map，简单理解：变量是map，key为KeyType(即：任意类型都可以)，value为切片。
func mapSlice() {
	fmt.Println("值为切片类型的map，简单理解：变量是map，key为KeyType(即：任意类型都可以)，value为切片。")
	var mapSlice = make(map[string][]string, 3) //3是map的容量（cap）
	fmt.Println(mapSlice)
	fmt.Println("after init")
	key := "中国"                //是map的key
	value, ok := mapSlice[key] //判断是否存在对应的key，如果不存在，则初始化
	if !ok {
		//初始化切片
		value = make([]string, 0, 2)
	}
	value = append(value, "北京", "上海") //往map对应的key中的value添加切片元素
	mapSlice[key] = value             //把切片赋值给map对应的key
	fmt.Println(mapSlice)

}
