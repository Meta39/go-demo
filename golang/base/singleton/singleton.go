package singleton

import (
	"sync"
	"sync/atomic"
)

// Singleton 单例模式
type Singleton struct{}

var instance *Singleton
var onceSingleton sync.Once

// GetInstance 安全(推荐使用)
func GetInstance() *Singleton {
	onceSingleton.Do(func() {
		// 在Do包裹的匿名函数里执行安全的初始化
		instance = &Singleton{}
	})
	return instance
}

// GetInstanceUnsafe 不建议使用此方式实现单例模式
func GetInstanceUnsafe() *Singleton {
	if instance == nil {
		instance = &Singleton{} // 不是并发安全的
	}
	return instance
}

// GetInstanceSafe 虽然安全，但不推荐
var muSingleton sync.Mutex

func GetInstanceSafe() *Singleton {
	muSingleton.Lock() // 如果实例存在没有必要加锁
	defer muSingleton.Unlock()

	if instance == nil {
		instance = &Singleton{}
	}
	return instance
}

// GetInstanceSafe2 虽然安全，但不推荐。
var initialized uint32

func GetInstanceSafe2() *Singleton {

	if atomic.LoadUint32(&initialized) == 1 { // 原子操作
		return instance
	}

	muSingleton.Lock()
	defer muSingleton.Unlock()

	if initialized == 0 {
		instance = &Singleton{}
		atomic.StoreUint32(&initialized, 1)
	}

	return instance
}
