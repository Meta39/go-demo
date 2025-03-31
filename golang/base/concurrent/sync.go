package concurrent

import (
	"fmt"
	"golang/base/singleton"
	"sync"
)

var safeMap = sync.Map{} //并发安全的map

func Sync() {
	fmt.Printf("sync相关方法")
	/*
		sync.WaitGroup
		在代码中生硬的使用time.Sleep肯定是不合适的，Go语言中可以使用sync.WaitGroup来实现并发任务的同步。

		sync.WaitGroup有以下几个方法：
		方法名										功能
		func (wg * WaitGroup) Add(delta int)		计数器+delta
		(wg *WaitGroup) Done()						计数器-1
		(wg *WaitGroup) Wait()						阻塞直到计数器变为0

		sync.WaitGroup内部维护着一个计数器，计数器的值可以增加和减少。例如当我们启动了 N 个并发任务时，就将计数器值增加N。
		每个任务完成时通过调用 Done 方法将计数器减1。
		通过调用 Wait 来等待并发任务执行完，当计数器值为 0 时，表示所有并发任务已经完成。
		注意：需要注意sync.WaitGroup是一个结构体，进行参数传递的时候要传递指针。
	*/
	fmt.Println("sync.WaitGroup代替Sleep实现并发任务的同步")
	wg.Add(1)
	go goroutineSayHello4() //启动另外一个goroutine去执行hello函数。wg.Done() 在 goroutineSayHello4 函数里面
	fmt.Println("main goroutine done!")
	wg.Wait()

	/*
		sync.Once（一般用来实现懒汉式单例模式，只在用的时候执行一次）
		在某些场景下我们需要确保某些操作即使在高并发的场景下也只会被执行一次，例如只加载一次配置文件等。
		Go语言中的sync包中提供了一个针对只执行一次场景的解决方案——sync.Once，sync.Once只有一个Do方法，其签名如下：

			func (o *Once) Do(f func())

		注意：如果要执行的函数f需要传递参数就需要搭配闭包来使用。

		sync.Once其实内部包含一个互斥锁和一个布尔值，互斥锁保证布尔值和数据的安全，而布尔值用来记录初始化是否完成。
		这样设计就能保证初始化操作的时候是并发安全的并且初始化操作也不会被执行多次。
	*/
	fmt.Println("借助sync.Once实现的并发安全的单例模式")
	singleton.GetInstance()

	/*
		Go 语言中内置的 map 不是并发安全的。
		高并发场景下就需要为 map 加锁来保证并发的安全性了，Go语言的sync包中提供了一个开箱即用的并发安全版 map——sync.Map。
		开箱即用表示其不用像内置的 map 一样使用 make 函数初始化就能直接使用。同时sync.Map内置了诸如Store、Load、LoadOrStore、Delete、Range等操作方法。

		方法名																					功能
		func (m *Map) Store(key, value interface{})												存储key-value数据
		func (m *Map) Load(key interface{}) (value interface{}, ok bool)						查询key对应的value
		func (m *Map) LoadOrStore(key, value interface{}) (actual interface{}, loaded bool)		查询或存储key对应的value
		func (m *Map) LoadAndDelete(key interface{}) (value interface{}, loaded bool)			查询并删除key
		func (m *Map) Delete(key interface{})													删除key
		func (m *Map) Range(f func(key, value interface{}) bool)								对map中的每个key-value依次调用f

	*/
	fmt.Println("sync.Map 并发安全版 map")
	safeMapOption()

}

// 使用sync.WaitGroup优化后的goroutineSayHello函数
func goroutineSayHello4() {
	defer wg.Done()
	fmt.Println("Hello Goroutine!")
}

// 并发安全的map
func safeMapOption() {
	wg := sync.WaitGroup{}
	// 对m执行20个并发的读写操作
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func(key int) {
			//第一个key是键，第二个key是值
			safeMap.Store(key, key)                 // 存储key-value
			value, _ := safeMap.Load(key)           // 根据key取值
			fmt.Printf("k=:%v,v:=%v\n", key, value) //虽然不是按顺序输出，但是存储和取值都是对的。
			wg.Done()
		}(i)
	}
	wg.Wait()
}
