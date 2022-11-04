package main

import "fmt"

// Singleton 保证一个类永远只能有一个对象，这个对象还能被系统的其他模块使用
type singleton struct{}

var instance *singleton = &singleton{}

func GetInstance() *singleton {
	return instance
}
func (s *singleton) SomeTing() {
	fmt.Println("单例中的某方法")
}
