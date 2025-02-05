package main

import (
	"fmt"
	"testing"
)

//写测试用例测试product.go
func TestProduct_Create(t *testing.T) {
	// 通过普通方式创建
	product1 := &Product1{}
	product1.SetName("p1")
	fmt.Println(product1.GetName())

	product2 := &Product2{}
	product2.SetName("p2")
	fmt.Println(product2.GetName())

}


func TestProductFactory_Create(t *testing.T) {
	productFactory := productFactory{}
	product1 := productFactory.Create(p1)   // 传入productType
	product1.SetName("p1")
	fmt.Println(product1.GetName())

	product2 := productFactory.Create(p2)   // 传入productType
	product2.SetName("p2")
	fmt.Println(product2.GetName())
}