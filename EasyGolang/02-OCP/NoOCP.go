package main

import "fmt"

type Banker struct {
}

// 平铺设计，没有OCP

// 存款业务
func (b *Banker) Save() {
	fmt.Println("进行了 存款业务...")
}

// 转账业务
func (b *Banker) Transfer() {
	fmt.Println("进行了 转账业务...")
}

// 支付业务
func (b *Banker) Pay() {
	fmt.Println("进行了 支付业务...")
}

func main() {
	// & 取址    * 指针
	banker := &Banker{}

	banker.Save()
	banker.Transfer()
	banker.Pay()

}
