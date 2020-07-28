package main

import "fmt"

/*
P53-P54
4.向切片添加元素
实例3-3 切片的动态扩容

Go语言提供append内建函数用于动态向切片Slice添加元素，它将返回新的切片Slice。

if slice size(len) <= cap:
Slice添加操作将在切片指向的原有数据上进行，会覆盖掉原有数组的值。

if slice size(len) > cap:
当前Slice不足以容纳更多元素，切片将会进行扩容。

具体扩容过程：
1.申请一个新的连续内存空间，空间大小一般是原有容量的2倍。
2.将原数组中的数据复制到新数组中，同时将切片slice中的指针指向新数组
3.将新元素添加到新数组中

如果原有数组还可以添加新的元素，但是切片自身的容量已经饱和，此时进行append操作，同样会扩容，申请新的内存空间。
*/

func main() {
	arr1 := [...]int{1, 2, 3, 4}
	arr2 := [...]int{1, 2, 3, 4}

	fmt.Printf("arr1 pointer is %p\n", &arr1)
	fmt.Printf("arr2 pointer is %p\n", &arr2)

	sli1 := arr1[0:2] // len 2, cap 4   [1 2]
//	sli0 := arr2[1:3] // len 3, cap 4, value is [1 2 3]
	sli2 := arr2[2:4] // len 2, cap 2   [3 4]

	fmt.Printf("sli1 pointer is %p, len is %v, cap is %v, value is %v\n", &sli1, len(sli1), cap(sli1), sli1)
//	fmt.Printf("sli0 pointer is %p, len is %v, cap is %v, value is %v\n", &sli0, len(sli0), cap(sli0), sli0)

	fmt.Printf("sli2 pointer is %p, len is %v, cap is %v, value is %v\n", &sli2, len(sli2), cap(sli2), sli2)

	newSli1 := append(sli1, 5)   //  [1 2 5]   cap足，直接覆盖原数组
	fmt.Printf("newSli1 pointer is %p, len is %v, cap is %v, value is %v\n", &newSli1, len(newSli1), cap(newSli1), newSli1)
	fmt.Printf("source arr1 become %v\n", arr1)    //  [1 2 5 4]

	newSli2 := append(sli2, 5)  // cap不足，进行了扩容
	fmt.Printf("newSli2 pointer is %p, len is %v, cap is %v, value is %v\n", &newSli2, len(newSli2), cap(newSli2), newSli2)
	fmt.Printf("source arr2 become %v\n", arr2)

	// 如果原有数组还可以添加新的元素，但是切片自身的容量已经饱和，此时进行append操作，同样会扩容，申请新的内存空间。
	arr3 := [...]int{1, 2, 3, 4}
	sli3 := arr3[0:2:2]   // [1 2] len 2, cap 2   begin,end,cap
	fmt.Printf("source arr3: %v\n", arr3)
	fmt.Printf("sli3 pointer is %p, len is %v, cap is %v, value is %v\n", &sli3, len(sli3), cap(sli3), sli3)

	newSli3 := append(sli3,5)
	fmt.Printf("newSli3 pointer is %p, len is %v, cap is %v, value is %v\n", &newSli3, len(newSli3), cap(newSli3), newSli3)
	fmt.Printf("source arr3 become %v\n", arr3)
}