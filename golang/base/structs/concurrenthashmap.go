package structs

import "sync"

/*
ConcurrentHashMap 模仿java ConcurrentHashMap并实现类似ComputeIfAbsent方法
如果是生产环境，推荐使用第三方库：concurrent-map
用法如下：
import cmap "github.com/orcaman/concurrent-map"

	cache := cmap.New()

	// 类似 computeIfAbsent 的操作
	val, ok := cache.Get("key1")
	if !ok {
		computedVal := "computed_value"
		cache.Set("key1", computedVal)
		val = computedVal
	}
	fmt.Println(val)
*/
type ConcurrentHashMap struct {
	sync.RWMutex
	items map[string]any
}

func NewConcurrentHashMap() *ConcurrentHashMap {
	return &ConcurrentHashMap{items: make(map[string]any)}
}

// Get 安全地获取值（使用读锁）
func (m *ConcurrentHashMap) Get(key string) (any, bool) {
	m.RLock()
	defer m.RUnlock()
	val, ok := m.items[key]
	return val, ok
}

// ComputeIfAbsent 实现原子性“不存在则计算”
func (m *ConcurrentHashMap) ComputeIfAbsent(key string, computeFunc func() any) any {
	// 先尝试读（不加锁）
	m.RLock()
	val, ok := m.items[key]
	m.RUnlock()
	if ok {
		return val
	}

	// 加写锁
	m.Lock()
	defer m.Unlock()

	// 双重检查（避免在加锁前其他协程已写入）
	if val, ok := m.items[key]; ok {
		return val
	}

	// 计算并存入
	newVal := computeFunc()
	m.items[key] = newVal
	return newVal
}
