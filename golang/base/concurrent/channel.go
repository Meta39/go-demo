package concurrent

import (
	"errors"
	"fmt"
	channel2 "golang/base/concurrent/channel"
	"sync"
	"time"
)

/*
Channel
单纯地将函数并发执行是没有意义的。函数与函数间需要交换数据才能体现并发执行函数的意义。
虽然可以使用共享内存进行数据交换，但是共享内存在不同的 goroutine 中容易发生竞态问题。
为了保证数据交换的正确性，很多并发模型中必须使用互斥量对内存进行加锁，这种做法势必造成性能问题。
Go语言采用的并发模型是CSP（Communicating Sequential Processes），提倡通过通信共享内存而不是通过共享内存而实现通信。
如果说 goroutine 是Go程序并发的执行体，channel就是它们之间的连接。channel是可以让一个 goroutine 发送特定值到另一个 goroutine 的通信机制。
Go 语言中的通道（channel）是一种特殊的类型。通道像一个传送带或者队列，总是遵循先入先出（First In First Out）的规则，保证收发数据的顺序。
每一个通道都是一个具体类型的导管，也就是声明channel的时候需要为其指定元素类型。

channel类型
channel是 Go 语言中一种特有的类型。声明通道类型变量的格式如下：

	var 变量名称 chan 元素类型

其中：
chan：是关键字
元素类型：是指通道中传递元素的类型

var ch1 chan int   // 声明一个传递整型的通道
var ch2 chan bool  // 声明一个传递布尔型的通道
var ch3 chan []int // 声明一个传递int切片的通道

channel零值未初始化的通道类型变量其默认零值是nil。

	var ch chan int
	fmt.Println(ch) // <nil>

初始化channel
声明的通道类型变量需要使用内置的make函数初始化之后才能使用。具体格式如下：

	make(chan 元素类型, [缓冲大小])

其中：
channel的缓冲大小是可选的。

channel通道共有发送（send）、接收(receive）和关闭（close）三种操作。而发送和接收操作都使用<-符号。
一个通道值是可以被垃圾回收掉的。通道通常由发送方执行关闭操作，并且只有在接收方明确等待通道关闭的信号时才需要执行关闭操作。
它和关闭文件不一样，通常在结束操作之后关闭文件是必须要做的，但关闭通道不是必须的。
关闭后的通道有以下特点：
1、对一个关闭的通道再发送值就会导致 panic。
2、对一个关闭的通道进行接收会一直获取值直到通道为空。
3、对一个关闭的并且没有值的通道执行接收操作会得到对应类型的零值。
4、关闭一个已经关闭的通道会导致 panic。
*/
func Channel() {
	fmt.Println("channel通道")
	ch := make(chan int, 1) //创建缓冲大小为1的通道。如果不设置缓冲通道，会报错：all goroutines are asleep - deadlock!
	ch <- 10                // 把10发送到ch中。发送不会阻塞，因为缓冲区有空位
	x := <-ch               // 从ch中接收值并赋值给变量x
	//<-ch      // 从ch中接收值，忽略结果
	fmt.Println(x)
	close(ch) //关闭通道（非必须）

	/*
		无缓冲的通道：无缓冲的通道又称为阻塞的通道。

		func main() {
			ch := make(chan int)
			ch <- 10
			fmt.Println("发送成功")
		}

		上面这段代码能够通过编译，但是执行的时候会出现以下错误：
			fatal error: all goroutines are asleep - deadlock!

			goroutine 1 [chan send]:
			main.main()
					.../main.go:8 +0x54

		deadlock表示我们程序中的 goroutine 都被挂起导致程序死锁了。为什么会出现deadlock错误呢？
		因为我们使用ch := make(chan int)创建的是无缓冲的通道，无缓冲的通道只有在有接收方能够接收值的时候才能发送成功，否则会一直处于等待发送的阶段。
		同理，如果对一个无缓冲通道执行接收操作时，没有任何向通道中发送值的操作那么也会导致接收操作阻塞。
		就像田径比赛中的4x100接力赛，想要完成交棒必须有一个能够接棒的运动员，否则只能等待。简单来说就是无缓冲的通道必须有至少一个接收方才能发送成功。
		上面的代码会阻塞在ch <- 10这一行代码形成死锁，那如何解决这个问题呢？
		其中一种可行的方法是创建一个 goroutine 去接收值，例如：
	*/
	fmt.Println("无缓冲的通道：无缓冲的通道又称为阻塞的通道。")
	ch2 := make(chan int) //无缓冲通道
	go func() {
		ch2 <- 10 // 在另一个goroutine中发送
		fmt.Println("发送成功")
	}()
	x2 := <-ch2             // 在主goroutine中接收
	fmt.Println("接收成功", x2) //使用无缓冲通道进行通信将导致发送和接收的 goroutine 同步化。因此，无缓冲通道也被称为同步通道。

	/*
		有缓冲的通道（类似java的阻塞队列，但是go没法选择队列满的情况该怎么办，通道满了的时候，直接报错。java是可以选择满了该如何操作的，比如：队列满了，交给主线程执行、抛出异常等。）
		只要通道的容量大于零，那么该通道就属于有缓冲的通道，通道的容量表示通道中最大能存放的元素数量。
		当通道内已有元素数达到最大容量后，再向通道执行发送操作就会阻塞，除非有从通道执行接收操作。
		就像小区的快递柜只有那么个多格子，格子满了就装不下了，就阻塞了，等到别人取走一个快递员就能往里面放一个。
		我们可以使用内置的len函数获取通道内元素的数量，使用cap函数获取通道的容量，虽然我们很少会这么做。

		还有另外一种解决上面死锁问题的方法，那就是使用有缓冲区的通道。我们可以在使用 make 函数初始化通道时，可以为其指定通道的容量，例如：
	*/
	fmt.Println("有缓冲的通道")
	ch3 := make(chan int, 1) // 创建一个容量为1的有缓冲区通道
	ch3 <- 10
	//ch3 <- 11 //因为没有接收者。如果通道满了，就会引发all goroutines are asleep - deadlock!
	fmt.Printf("有缓冲的通道发送成功.len:%v,cap:%v\n", len(ch3), cap(ch3))

	//典型生产-消费模式
	fmt.Println("有缓冲的通道（典型生产-消费模式）")
	channelProductionConsumptionMode()

	/*
		多返回值模式
		当向通道中发送完数据时，我们可以通过close函数来关闭通道。当一个通道被关闭后，再往该通道发送值会引发panic，从该通道取值的操作会先取完通道中的值。
		通道内的值被接收完后再对通道执行接收操作得到的值会一直都是对应元素类型的零值。那我们如何判断一个通道是否被关闭了呢？
		对一个通道执行接收操作时支持使用如下多返回值模式。

			value, ok := <- ch

		其中：
		value：从通道中取出的值，如果通道被关闭则返回对应类型的零值。
		ok：通道ch关闭时返回 false，否则返回 true。
		下面代码片段中的f2函数会循环从通道ch中接收所有值，直到通道被关闭后退出。
	*/
	ch4 := make(chan int, 2)
	ch4 <- 1
	ch4 <- 2
	close(ch4)
	fmt.Println("通道多返回值模式")
	channelMultipleReturnValues(ch4)

	/*
		for range接收值
		通常我们会选择使用for range循环从通道中接收值，当通道被关闭后，会在通道内的所有值被接收完毕后会自动退出循环。
		上面那个示例 channelMultipleReturnValues 我们使用for range改写后会很简洁。
		注意：目前Go语言中并没有提供一个不对通道进行读取操作就能判断通道是否被关闭的方法。不能简单的通过len(ch)操作来判断通道是否被关闭。
	*/
	fmt.Println("使用for range接收值简化通道多返回值模式")
	ch5 := make(chan int, 2)
	ch5 <- 1
	ch5 <- 2
	close(ch5)
	for v := range ch5 {
		fmt.Println(v)
	}

	/*
		单向通道
		在某些场景下我们可能会将通道作为参数在多个任务函数间进行传递，通常我们会选择在不同的任务函数中对通道的使用进行限制，比如限制通道在某个函数中只能执行发送或只能执行接收操作。
		想象一下，我们现在有Producer和Consumer两个函数，其中Producer函数会返回一个通道，并且会持续将符合条件的数据发送至该通道，并在发送完成后将该通道关闭。
		而Consumer函数的任务是从通道中接收值进行计算，这两个函数之间通过Processes函数返回的通道进行通信。
		比如：消息队列，一个负责发送，一个负责接收。
	*/
	fmt.Println("单向通道(不强制)")
	ch6 := channel2.Producer()     //生产者，只负责发送
	res6 := channel2.Consumer(ch6) //消费者，只负责接收
	fmt.Println(res6)              // 25

	/*
		从上面的示例代码中可以看出正常情况下Consumer函数中只会对通道进行接收操作，但是这不代表不可以在Consumer函数中对通道进行发送操作。
		作为Producer函数的提供者，我们在返回通道的时候可能只希望调用方拿到返回的通道后只能对其进行接收操作。
		但是我们没有办法阻止在Consumer函数中对通道进行发送操作。
		Go语言中提供了单向通道来处理这种需要限制通道只能进行某种操作的情况。

			<- chan int // 只接收通道，只能接收不能发送
			chan <- int // 只发送通道，只能发送不能接收

		其中，箭头<-和关键字chan的相对位置表明了当前通道允许的操作，这种限制将在编译阶段进行检测。
		另外对一个只接收通道执行close也是不允许的，因为默认通道的关闭操作应该由发送方来完成。
		我们使用单向通道将上面的示例代码进行如下改造。
	*/
	fmt.Println("单向通道(强制)")
	ch7 := channel2.Producer2()
	res7 := channel2.Consumer2(ch7)
	fmt.Println(res7) // 25
	/*
		这一次，Producer函数返回的是一个只接收通道，这就从代码层面限制了该函数返回的通道只能进行接收操作，保证了数据安全。
		返回限制操作的单向通道也会让代码语义更清晰、更易读。在函数传参及任何赋值操作中全向通道（正常通道）可以转换为单向通道，但是无法反向转换。
	*/

	var ch8 = make(chan int, 1)
	ch8 <- 10
	close(ch8)
	channel2.Consumer2(ch8) // 函数传参时将ch8转为单向通道
	//close(ch8)//对已经关闭的通道再执行 close 也会引发 panic。

	var ch9 = make(chan int, 1)
	ch9 <- 10
	var ch10 <-chan int // 声明一个只接收通道ch10
	ch10 = ch9          // 变量赋值时将ch9转为单向通道
	<-ch10

	/*
		select多路复用
		Go 语言内置了select关键字，使用它可以同时响应多个通道的操作。
		Select 的使用方式类似于之前学到的 switch 语句，它也有一系列 case 分支和一个默认的分支。
		每个 case 分支会对应一个通道的通信（接收或发送）过程。
		select 会一直等待，直到其中的某个 case 的通信操作完成时，就会执行该 case 分支对应的语句。具体格式如下：

		select {
		case <-ch1:
			//...
		case data := <-ch2:
			//...
		case ch3 <- 10:
			//...
		default:
			//默认操作
		}

		Select 语句具有以下特点。
		一、可处理一个或多个 channel 的发送/接收操作。
		二、如果多个 case 同时满足，select 会随机选择一个执行。
		三、对于没有 case 的 select 会一直阻塞，可用于阻塞 main 函数，防止退出。

		下面的示例代码能够在终端打印出10以内的奇数，我们借助这个代码片段来看一下 select 的具体使用。
	*/
	fmt.Println("select多路复用")
	channelSelect()

	//通道误用示例
	fmt.Println("通道误用示例及修复")
	//demo1()//主要问题是用了for一直循环，没有break。
	demo1Bugfix()
	//demo2()//主要问题是队列一直处于阻塞，通道无法接收
	_ = demo2Bugfix()
}

// 通道典型生产-消费模式
func channelProductionConsumptionMode() {
	ch := make(chan int, 3) // 容量3

	// 生产者
	go func() {
		for i := 0; i < 5; i++ {
			ch <- i
			fmt.Printf("发送 %d\n", i)
		}
		close(ch)
	}()

	// 消费者
	for v := range ch {
		fmt.Printf("接收 %d\n", v)
		//time.Sleep(time.Second) // 模拟处理耗时
	}
}

// 通道多返回值
func channelMultipleReturnValues(ch chan int) {
	//for { ... } 是一个无限循环，它会一直执行循环体内的代码，直到遇到 break、return 或 panic 才会退出。
	for {
		v, ok := <-ch //多返回值。v:值，ok:通道ch关闭时返回 false，否则返回 true
		if !ok {
			fmt.Println("通道已关闭")
			break
		}
		fmt.Printf("v:%#v ok:%#v\n", v, ok)
	}
}

// select多路复用
func channelSelect() {
	/*
		迭代		通道状态		可能执行的 		case				结果
		i=1		空			ch <- 1（发送）	通道存入 				1
		i=2		有值 		1				x := <-ch（接收）		打印 1，通道变空
		i=3		空			ch <- 3（发送）	通道存入 				3
		i=4		有值 		3				x := <-ch（接收）		打印 3，通道变空
		...	...	...	...
	*/
	ch := make(chan int, 1)
	for i := 1; i <= 10; i++ {
		select {
		case x := <-ch: //i = 偶数条件才成立,因为偶数的时候ch里面才有值。
			fmt.Println(x)
		case ch <- i: //i = 奇数条件才成立，因为奇数的时候才会往ch写入i值
		}
	}
}

/*
demo1 通道误用导致的bug
将上述代码编译执行后，匿名函数所在的 goroutine 并不会按照预期在通道被关闭后退出。
因为task := <- ch的接收操作在通道被关闭后会一直接收到零值，而不会退出。此处的接收操作应该使用task, ok := <- ch ，通过判断布尔值ok为假时退出；或者使用select 来处理通道。
*/
func demo1() {
	wg := sync.WaitGroup{}

	ch := make(chan int, 10)
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch) //问题1：在启动 goroutine 前就关闭了通道，导致获取ch时可能读到零值

	wg.Add(3)
	for j := 0; j < 3; j++ {
		go func() {
			//问题3：所有 goroutine 共享同一个 j 的引用！
			for { //问题2：死循环，应改为for task := range ch {  // 自动检测通道关闭
				//for task := range ch {  //问题2的修复。自动检测通道关闭
				task := <-ch
				// 这里假设对接收的数据执行某些操作
				fmt.Println(task)
			}
			wg.Done()
		}()

		/*
				问题3修复。
				go func(id int) {
					for task := range ch {
						task := <-ch
						fmt.Println(task)
					}
			    }(j)  // 传递当前 j 的值
		*/
	}

	//close(ch)//问题1的修复。确保所有 goroutine 启动后再关闭
	wg.Wait()
}

// 修复demo1的bug
func demo1Bugfix() {
	wg := sync.WaitGroup{}
	ch := make(chan int, 10)

	// 生产数据
	for i := 0; i < 10; i++ {
		ch <- i
	}

	// 启动 3 个消费者 goroutine
	wg.Add(3)
	for j := 0; j < 3; j++ {
		go func(id int) {
			defer wg.Done() // 确保退出时调用 Done
			for task := range ch {
				fmt.Printf("worker %d: %d\n", id, task)
			}
		}(j)
	}

	close(ch) // 所有 goroutine 启动后再关闭通道
	wg.Wait() // 等待所有 goroutine 退出
}

/*
demo2 通道误用导致的bug
上述代码片段可能导致 goroutine 泄露（goroutine 并未按预期退出并销毁）。
由于 select 命中了超时逻辑，导致通道没有消费者（无接收操作），而其定义的通道为无缓冲通道，因此 goroutine 中的ch <- "job result"操作会一直阻塞，最终导致 goroutine 泄露。
*/
func demo2() {
	ch := make(chan string) //问题1：通道未关闭
	go func() {
		// 这里假设执行一些耗时的操作
		time.Sleep(3 * time.Second) //问题2：超时时间比较大time.Sleep(3 * time.Second) 大于 <-time.After(time.Second)
		ch <- "job result"
	}()

	select {
	case result := <-ch:
		fmt.Println(result)
	case <-time.After(time.Second): // 较小的超时时间
		return
	}
}

// 修复demo2的问bug
func demo2Bugfix() error {
	ch := make(chan string, 1) // 缓冲通道避免阻塞发送

	go func() {
		defer close(ch) //确保关闭通道
		time.Sleep(3 * time.Second)
		ch <- "job result"
	}()

	select {
	case result, ok := <-ch:
		if !ok {
			return errors.New("通道意外关闭")
		}
		fmt.Println("结果:", result)
		return nil
	case <-time.After(5 * time.Second): //调大超时时间，防止处理超时。
		return errors.New("处理超时")
	}
}
