package concurrent

import (
	"fmt"
	"sync"
	"time"
)

var (
	unlockNum      int64        //不加锁修改
	mutexLockNum   int64        //加互斥锁修改
	rwMutexLockNum int64        //加读写互斥锁修改
	mutexLock      sync.Mutex   //互斥锁
	rwMutexLock    sync.RWMutex //读写互斥锁
)

/*
Lock 并发安全和锁
*/
func Lock() {
	fmt.Println("lock锁")
	fmt.Println("先不加锁进行修改全局变量，进而引出问题。")
	/*
		在下面的示例代码片中，我们开启了5个 goroutine 分别执行 add 函数，这5个 goroutine 在访问和修改全局的unlockNum变量时就会存在数据竞争;
		某个 goroutine 中对全局变量unlockNum的修改可能会覆盖掉另一个 goroutine 中的操作，所以导致最后的结果与预期不符。
	*/
	wg.Add(5)
	go add()
	go add()
	go add()
	go add()
	go add()

	wg.Wait()
	fmt.Printf("不加锁，结果%v != 25000\n", unlockNum) //结果并不是25000

	/*
		互斥锁
		互斥锁是一种常用的控制共享资源访问的方法，它能够保证同一时间只有一个 goroutine 可以访问共享资源。Go 语言中使用sync包中提供的Mutex类型来实现互斥锁。
		使用互斥锁能够保证同一时间有且只有一个 goroutine 进入临界区，其他的 goroutine 则在等待锁；
		当互斥锁释放后，等待的 goroutine 才可以获取锁进入临界区，多个 goroutine 同时等待一个锁时，唤醒的策略是随机的。
		sync.Mutex提供了两个方法供我们使用。

		方法名						功能
		func (m *Mutex) Lock()		获取互斥锁
		func (m *Mutex) Unlock()	释放互斥锁

		在Java中，与Go语言的sync.Mutex最相似的锁是 java.redisutil.concurrent.locks.ReentrantLock（非公平模式）
		一、两者都是互斥锁（Mutual Exclusion），保证同一时间只有一个线程能访问临界区。
		二、均支持显式的加锁（Lock）和解锁（Unlock）操作，需要手动控制锁的生命周期。

		特性			Go sync.Mutex			Java ReentrantLock
		重入性		不支持					支持（可通过计数器实现）
		公平性		非公平（唤醒顺序随机）		可配置公平/非公平（默认非公平）
		条件变量		需配合sync.Cond使用		内置Condition机制（newCondition()）
		锁中断		不支持					支持lockInterruptibly()
		尝试加锁		不支持					支持tryLock()

		我们在下面的示例代码中使用互斥锁限制每次只有一个 goroutine 才能修改全局变量x，从而修复上面代码中的问题。
	*/
	wg.Add(5)
	go addByLock()
	go addByLock()
	go addByLock()
	go addByLock()
	go addByLock()

	wg.Wait()
	fmt.Printf("加锁（互斥锁sync.Mutex），结果%v == 25000\n", mutexLockNum) //结果是25000

	/*
		读写互斥锁
		互斥锁是完全互斥的，但是实际上有很多场景是读多写少的，当我们并发的去读取一个资源而不涉及资源修改的时候是没有必要加互斥锁的，这种场景下使用读写锁是更好的一种选择。
		读写锁在 Go 语言中使用sync包中的RWMutex类型。
		sync.RWMutex提供了以下5个方法。
		方法名									功能
		func (rw *RWMutex) Lock()				获取写锁
		func (rw *RWMutex) Unlock()				释放写锁
		func (rw *RWMutex) RLock()				获取读锁
		func (rw *RWMutex) RUnlock()			释放读锁
		func (rw *RWMutex) RLocker() Locker		返回一个实现Locker接口的读写锁

		读写锁分为两种：读锁和写锁。当一个 goroutine 获取到读锁之后，其他的 goroutine 如果是获取读锁会继续获得锁，如果是获取写锁就会等待；
		而当一个 goroutine 获取写锁之后，其他的 goroutine 无论是获取读锁还是写锁都会等待。
		下面我们使用代码构造一个读多写少的场景，然后分别使用互斥锁和读写锁查看它们的性能差异。
		从最终的执行结果可以看出，使用读写互斥锁在读多写少的场景下能够极大地提高程序的性能。
		不过需要注意的是如果一个程序中的读操作和写操作数量级差别不大，那么读写互斥锁的优势就发挥不出来。
		结论：绝大数情况下都应该用互斥锁，而不是读写互斥锁。遇锁不定，互斥锁！因为没有需要那么多数量级的数据需要处理！工作中大多数需要处理的数据数量级都很低。
	*/
	// 使用互斥锁，10并发写，1000并发读
	do("互斥锁", writeWithLock, readWithLock, 10, 1000) // x:10 cost:1.466500951s

	// 使用读写互斥锁，10并发写，1000并发读
	do("读写互斥锁", writeWithRWLock, readWithRWLock, 10, 1000) // x:10 cost:117.207592ms

}

// add 不加锁对全局变量unlockNum执行5000次加1操作
func add() {
	for i := 0; i < 5000; i++ {
		unlockNum = unlockNum + 1
	}
	wg.Done()
}

// addByLock 加锁对全局变量lockNum1执行5000次加1操作
func addByLock() {
	mutexLock.Lock() // 修改x前加锁
	for i := 0; i < 5000; i++ {
		mutexLockNum = mutexLockNum + 1
	}
	mutexLock.Unlock() // 改完解锁
	wg.Done()
}

// writeWithLock 使用互斥锁的写操作
func writeWithLock() {
	mutexLock.Lock() // 加互斥锁
	rwMutexLockNum = rwMutexLockNum + 1
	time.Sleep(10 * time.Millisecond) // 假设读操作耗时10毫秒
	mutexLock.Unlock()                // 解互斥锁
	wg.Done()
}

// readWithLock 使用互斥锁的读操作
func readWithLock() {
	mutexLock.Lock()             // 加互斥锁
	time.Sleep(time.Millisecond) // 假设读操作耗时1毫秒
	mutexLock.Unlock()           // 释放互斥锁
	wg.Done()
}

// writeWithLock 使用读写互斥锁的写操作
func writeWithRWLock() {
	rwMutexLock.Lock() // 加写锁
	rwMutexLockNum = rwMutexLockNum + 1
	time.Sleep(10 * time.Millisecond) // 假设读操作耗时10毫秒
	rwMutexLock.Unlock()              // 释放写锁
	wg.Done()
}

// readWithRWLock 使用读写互斥锁的读操作
func readWithRWLock() {
	rwMutexLock.RLock()          // 加读锁
	time.Sleep(time.Millisecond) // 假设读操作耗时1毫秒
	rwMutexLock.RUnlock()        // 释放读锁
	wg.Done()
}

/*
wf, rf指传满足无入参的函数名
wc表示循环的开始值
rc表示循环的结束值
*/
func do(n string, wf, rf func(), wc, rc int) {
	start := time.Now()
	// wc个并发写操作
	for i := 0; i < wc; i++ {
		wg.Add(1)
		go wf()
	}

	//  rc个并发读操作
	for i := 0; i < rc; i++ {
		wg.Add(1)
		go rf()
	}

	wg.Wait()
	cost := time.Since(start)
	fmt.Printf("%v x:%v cost:%vms\n", n, rwMutexLockNum, cost.Milliseconds())
}
