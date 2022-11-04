package __singleton

import "sync"

var once sync.Once

type singleton struct{}

var instance *singleton

func GetInstance() *singleton {
	once.Do(func() {
		instance = &singleton{}
	})
	return instance
}
