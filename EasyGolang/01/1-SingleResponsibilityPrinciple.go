package main

import "fmt"

// https://www.yuque.com/aceld/lfhu8y/ckdht9
// 4：20 todo

// 穿衣服的方式
type Clothes struct {
	// 逛街穿衣

}

// 工作穿衣
func (c *Clothes) OnWork() {
	fmt.Println("工作装扮")
}

// 逛街穿衣
func (c *Clothes) OnShopping() {
	fmt.Println("逛街装扮")
}

func main() {
	c := Clothes{}

	// 工作的业务
	fmt.Println("在工作......")
	c.OnWork()

	// 逛街的业务
	fmt.Println("在逛街......")
	c.OnShopping()

}
