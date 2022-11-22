package main

import "fmt"

/*
ref: https://www.yuque.com/aceld/lfhu8y/ckdht9
单一职责原则
(Single Responsibility Principle, 01-SRP)
定义：类的职责单一，对外只提供一种功能，而引起类变化的原因都应该只有一个。

每个类只有一个方法。当然实际项目中不太可能每个类一个方法，那样耦合太高，不符合高内聚低耦合的原则。
*/

// 穿衣服的方式
type ClothesShop struct {
}

func (cs *ClothesShop) Style() {
	fmt.Println("逛街装扮")
}

type ClothesWork struct {
}

func (cw *ClothesWork) Style() {
	fmt.Println("工作装扮")
}

func main() {
	// 工作的业务
	cw := ClothesWork{}
	cw.Style()

	// 逛街的业务
	cs := ClothesShop{}
	cs.Style()

}
