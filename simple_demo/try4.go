package main

import "fmt"

func main() {

	e := []int32{1,2,3}
	fmt.Println("cap of e before:",cap(e))
	e = append(e,4)
	fmt.Println("cap of e after:",cap(e))

	f := []int{1,2,3}
	fmt.Println("cap of f before:",cap(f))
	f = append(f,4)
	fmt.Println("cap of f after:",cap(f))

	// 扩容后的slice容量，还和类型有关呢。

	//cap of e before: 3
	//cap of e after: 8
	//cap of f before: 3
	//cap of f after: 6

	/*
	append的时候发生扩容的动作

	append单个元素，或者append少量的多个元素，这里的少量指double之后的容量能容纳，这样就会走以下扩容流程，不足1024，双倍扩容，超过1024的，1.25倍扩容。

	若是append多个元素，且double后的容量不能容纳，直接使用预估的容量。

	敲重点！！！！
	此外，以上两个分支得到新容量后，均需要根据slice的类型size，算出新的容量所需的内存情况capmem，然后再进行capmem向上取整，得到新的所需内存，除上类型size，得到真正的最终容量,作为新的slice的容量。

	ref：https://www.jianshu.com/p/303daad705a3

	 */
	
}
