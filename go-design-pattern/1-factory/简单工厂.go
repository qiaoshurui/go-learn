package main

import "fmt"

// Fruit 抽象层
type Fruit interface {
	Show()
}

type Apple struct {
}

func (a *Apple) Show() {
	fmt.Println("我是苹果")
}

type Banana struct {
}

func (b *Banana) Show() {
	fmt.Println("我是香蕉")
}

type Pear struct {
}

func (p *Pear) Show() {
	fmt.Println("我是梨")
}

type Factory struct {
}

// CreatFruit 工厂的生产器返回的是抽象的方法
func (f *Factory) CreatFruit(name string) Fruit {
	var fruit Fruit
	if name == "apple" {
		fruit = new(Apple)
	} else if name == "banana" {
		fruit = new(Banana)
	} else if name == "pear" {
		fruit = new(Pear)
	}
	return fruit
}
func main() {
	factory := new(Factory)
	apple := factory.CreatFruit("apple")
	apple.Show()
	banana := factory.CreatFruit("banana")
	banana.Show()
	pear := factory.CreatFruit("pear")
	pear.Show()
}
