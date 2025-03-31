package concurrent

import (
	"fmt"
	"golang.org/x/sync/errgroup"
	"net/http"
	"sync"
	"time"
)

var wg sync.WaitGroup // 声明全局等待组变量

/*
Goroutine 是 Go 语言支持并发的核心，在一个Go程序中同时创建成百上千个goroutine是非常普遍的，一个goroutine会以一个很小的栈开始其生命周期，一般只需要2KB。
区别于操作系统线程由系统内核进行调度， goroutine 是由Go运行时（runtime）负责调度。
例如Go运行时会智能地将 m个goroutine 合理地分配给n个操作系统线程，实现类似m:n的调度机制，不再需要Go开发者自行在代码层面维护一个线程池。
Goroutine 是 Go 程序中最基本的并发执行单元。每一个 Go 程序都至少包含一个 goroutine——main goroutine，当 Go 程序启动时它会自动创建。
在Go语言编程中你不需要去自己写进程、线程、协程；
你的技能包里只有一个技能——goroutine，当你需要让某个任务并发执行的时候，你只需要把这个任务包装成一个函数，开启一个 goroutine 去执行这个函数就可以了，就是这么简单粗暴。
*/
func Goroutine() {
	/*
		go关键字
		Go语言中使用 goroutine 非常简单，只需要在函数或方法调用前加上go关键字就可以创建一个 goroutine ，从而让该函数或方法在新创建的 goroutine 中执行。

			go f()  // 创建一个新的 goroutine 运行函数f

		匿名函数也支持使用go关键字创建 goroutine 去执行。

			go func(){
			  // ...
			}()

		一个 goroutine 必定对应一个函数/方法，可以创建多个 goroutine 去执行相同的函数/方法。
	*/
	fmt.Println("使用sleep函数阻塞主线程，防止goroutineSayHello没执行完毕，主线程就关闭了。")
	go goroutineSayHello() //执行后并不会输出hello，因为main函数已经执行完毕。main函数关闭后，其它goroutine 也会关闭，无法正常执行。

	/*
		为什么上面的go goroutineSayHello()并没有打印 hello。这是为什么呢？
		其实在 Go 程序启动时，Go 程序就会为 main 函数创建一个默认的 goroutine 。
		在上面的代码中我们在 main 函数中使用 go 关键字创建了另外一个 goroutine 去执行 hello 函数，而此时 main goroutine 还在继续往下执行，我们的程序中此时存在两个并发执行的 goroutine。
		当 main 函数结束时整个程序也就结束了，同时 main goroutine 也结束了，所有由 main goroutine 创建的 goroutine 也会一同退出。
		也就是说我们的 main 函数退出太快，另外一个 goroutine 中的函数还未执行完程序就退出了，导致未打印出“hello”。
		main goroutine 就像是《权利的游戏》中的夜王，其他的 goroutine 都是夜王转化出的异鬼，夜王一死它转化的那些异鬼也就全部GG了。
		所以我们要想办法让 main 函数“等一等”将在另一个 goroutine 中运行的 hello 函数。
		其中最简单粗暴的方式就是在 main 函数中“time.Sleep”1秒钟了（这里的1秒钟是我们根据经验而设置的一个值，在这个示例中1秒钟足够创建新的goroutine执行完hello函数了）。
	*/
	time.Sleep(time.Second) //阻塞main函数1秒，让另一个goroutine正常执行完 goroutineSayHello 函数打印hello，跟java里的sleep方法是一样的效果，生产环境千万不要这样干！！！

	/*
		在上面的程序中使用time.Sleep让 main goroutine 等待 goroutineSayHello goroutine执行结束是不优雅的，当然也是不准确的。
		Go 语言中通过sync包为我们提供了一些常用的并发原语，我们会在后面的小节单独介绍sync包中的内容。在这一小节，我们会先介绍一下 sync 包中的WaitGroup。
		当你并不关心并发操作的结果或者有其它方式收集并发操作的结果时，WaitGroup是实现等待一组并发操作完成的好方法。
		下面的示例代码中我们在 main goroutine 中使用sync.WaitGroup来等待 goroutineSayHello goroutine 完成后再退出。
		将代码编译后再执行，得到的输出结果和之前一致，但是这一次程序不再会有多余的停顿，hello goroutine 执行完毕后程序直接退出。
		PS：sync.WaitGroup 类似 java 多线程的 join() 或【wait() / notify()】，但Go的协程更轻量，java的线程比较重。
	*/
	fmt.Println("sync.WaitGroup Wait 阻塞 main 主线程的执行，等 goroutineSayHello2 Done 执行完才继续执行 main 函数")
	wg.Add(1)               // 登记1个goroutine
	go goroutineSayHello2() //里面要有wg.Done() // 告知当前goroutine完成
	wg.Wait()               // 阻塞等待登记的goroutine完成

	/*
		启动多个goroutine
		在 Go 语言中实现并发就是这样简单，我们还可以启动多个 goroutine 。
		同样使用了sync.WaitGroup来实现 goroutine 的同步。
		sync.WaitGroup Add添加 => defer Done 函数执行完，再通知执行完毕 => Wait等待全部执行完毕
	*/
	fmt.Println("启动多个goroutine。多次执行下面的代码会发现每次终端上打印数字的顺序都不一致。这是因为10个 goroutine 是并发执行的，而 goroutine 的调度是随机的。")
	for i := 0; i < 10; i++ {
		wg.Add(1) // 启动一个goroutine就登记+1
		go goroutineSayHello3(i)
	}
	wg.Wait() // 等待所有登记的goroutine都结束

	/*
		动态栈
		操作系统的线程一般都有固定的栈内存（通常为2MB）,而 Go 语言中的 goroutine 非常轻量级，一个 goroutine 的初始栈空间很小（一般为2KB）；
		所以在 Go 语言中一次创建数万个 goroutine 也是可能的。
		并且 goroutine 的栈不是固定的，可以根据需要动态地增大或缩小， Go 的 runtime 会自动为 goroutine 分配合适的栈空间。
	*/

	/*
		goroutine调度
		操作系统内核在调度时会挂起当前正在执行的线程并将寄存器中的内容保存到内存中，然后选出接下来要执行的线程并从内存中恢复该线程的寄存器信息，然后恢复执行该线程的现场并开始执行线程。从一个线程切换到另一个线程需要完整的上下文切换。
		因为可能需要多次内存访问，索引这个切换上下文的操作开销较大，会增加运行的cpu周期。
		区别于操作系统内核调度操作系统线程，goroutine 的调度是Go语言运行时（runtime）层面的实现，是完全由 Go 语言本身实现的一套调度系统——go scheduler。
		它的作用是按照一定的规则将所有的 goroutine 调度到操作系统线程上执行。
		在经历数个版本的迭代之后，目前 Go 语言的调度器采用的是 GPM 调度模型。
		单从线程调度讲，Go语言相比起其他语言的优势在于OS线程是由OS内核来调度的， goroutine 则是由Go运行时（runtime）自己的调度器调度的，完全是在用户态下完成的;
		不涉及内核态与用户态之间的频繁切换，包括内存的分配与释放，都是在用户态维护着一块大的内存池， 不直接调用系统的malloc函数（除非内存池需要改变），成本比调度OS线程低很多。
		另一方面充分利用了多核的硬件资源，近似的把若干goroutine均分在物理线程上， 再加上本身 goroutine 的超轻量级，以上种种特性保证了 goroutine 调度方面的性能。
	*/

	/*
		GOMAXPROCS
		Go运行时的调度器使用GOMAXPROCS参数来确定需要使用多少个 OS 线程来同时执行 Go 代码。
		默认值是机器上的 CPU 核心数。例如在一个 8 核心的机器上，GOMAXPROCS 默认为 8。
		Go语言中可以通过runtime.GOMAXPROCS函数设置当前程序并发时占用的 CPU逻辑核心数。
		（Go1.5版本之前，默认使用的是单核心执行。Go1.5 版本之后，默认使用全部的CPU 逻辑核心数。）
	*/

	/*
		在启用 goroutine 去执行任务的场景下，如果想要 recover goroutine 中可能出现的 panic 就需要在 goroutine 中使用 recover。
		如下所示：
	*/
	fmt.Println("goroutine 中处理 panic 就需要在 goroutine 中使用 recover")
	goroutineHandlePanic()

	/*
		errgroup 它能为处理公共任务的子任务而开启的一组 goroutine 提供同步、error 传播和基于context 的取消功能。
		注意：errgroup 属于 golang.org/x/sync/errgroup，而不是标准库。因此需要用以下命令进行导入：
		go get golang.org/x/sync/errgroup

		errgroup 包中定义了一个 Group 类型，它包含了若干个不可导出的字段。

			type Group struct {
				cancel func()

				wg sync.WaitGroup

				errOnce sync.Once
				err     error
			}

		errgroup.Group 提供了Go和Wait两个方法。

			func (g *Group) Go(f func() error)
			func (g *Group) Wait() error
			func WithContext(ctx context.Context) (*Group, context.Context)

		Go 函数会在新的 goroutine 中调用传入的函数f；第一个返回非零错误的调用将取消该Group；下面的Wait方法会返回该错误。
		Wait 会阻塞直至由上述 Go 方法调用的所有函数都返回，然后从它们返回第一个非nil的错误（如果有）。
		WithContext 创建带有 cancel 方法的 errgroup.Group
		下面的示例代码演示了如何使用 errgroup 包来处理多个子任务 goroutine 中可能返回的 error。
	*/
	fmt.Println("errgroup")
	err := errgroupHandleError()
	if err != nil {
		fmt.Println(err)
		return
	}

}

func goroutineSayHello() {
	fmt.Println("hello")
}

func goroutineSayHello2() {
	fmt.Println("hello2")
	wg.Done() // 告知当前goroutine完成
}

func goroutineSayHello3(i int) {
	defer wg.Done() // goroutine结束就登记-1
	fmt.Println("hello", i)
}

/*
goroutine中处罚的panic，只能在goroutine中用defer 匿名函数中进行recover处理
*/
func goroutineHandlePanic() {
	// 开启一个goroutine执行任务
	go func() {
		//goroutine中处罚的panic，只能在goroutine中用defer 匿名函数中进行recover处理
		defer func() {
			if r := recover(); r != nil {
				fmt.Printf("recover inner panic:%v\n", r)
			}
		}()
		fmt.Println("in goroutine....")
		// 只能触发当前goroutine中的defer
		panic("panic in goroutine")
	}()

	time.Sleep(time.Second)
	fmt.Println("exit")
}

/*
使用 errgroup 包来处理多个子任务 goroutine 中可能返回的 error。
注意：errgroup 属于 golang.org/x/sync/errgroup，而不是标准库。因此需要用以下命令进行导入：
go get golang.org/x/sync/errgroup
*/
func errgroupHandleError() error {
	g := new(errgroup.Group) // 创建等待组（类似sync.WaitGroup）
	var urls = []string{
		"https://golang.google.cn",
		"https://www.baidu.com",
		"https://www.unkonw.com", //不存在的网页
	}
	for _, url := range urls {
		url := url // 注意此处声明新的变量
		// 启动一个goroutine去获取url内容
		g.Go(func() error {
			resp, err := http.Get(url)
			if err == nil {
				fmt.Printf("%s获取网页成功\n", url)
				_ = resp.Body.Close() //这里的错误不是重点，因此略过
			}
			return err // 返回错误
		})
	}

	//获取返回错误
	if err := g.Wait(); err != nil {
		// 处理可能出现的错误
		fmt.Println("获取网页失败", err)
		return err
	}
	fmt.Println("所有goroutine均成功")
	return nil
}
