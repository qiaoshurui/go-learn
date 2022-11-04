package __singleton

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// Singleton 保证一个类永远只能有一个对象，这个对象还能被系统的其他模块使用
type singleton struct{}

var instance *singleton

// 定义一个锁
var lock sync.Mutex

// 定义一个标记
var initialized uint32

func GetInstance() *singleton {
	if atomic.LoadUint32(&initialized) == 1 {
		return instance
	}
	lock.Lock()
	defer lock.Unlock()
	//只有首次GetInstance()方法被调用，才会生成这个单例的对象
	if instance == nil {
		instance = &singleton{}
		atomic.StoreUint32(&initialized, 1)
		return instance
	}
	return instance
}
func (s *singleton) SomeTing() {
	fmt.Println("单例中的某方法")
}
