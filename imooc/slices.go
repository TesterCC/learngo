package main

import "fmt"

/*
3-2 切片的概念
https://www.bilibili.com/video/BV18Q4y1M7NV?p=10

3-3 切片的操作
https://www.bilibili.com/video/BV18Q4y1M7NV?p=11
*/

func updateSlice(s []int) {
	s[0] = 100
}

func main() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	fmt.Println("arr[2:6] = ", arr[2:6]) // slice是array的视图view
	fmt.Println("arr[:6] = ", arr[:6])
	fmt.Println("arr[2:] = ", arr[2:])
	fmt.Println("arr[:] = ", arr[:])

	arr2 := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	s1 := arr2[2:]
	fmt.Println("new s1 arr2[2:] = ", s1)
	s2 := arr2[:]
	fmt.Println("new s2 arr2[:] = ", s2)

	fmt.Println("After updateSlice(s1)")
	updateSlice(s1)
	fmt.Println(s1)
	fmt.Println(arr2)

	fmt.Println("After updateSlice(s2)")
	updateSlice(s2)
	fmt.Println(s2)
	fmt.Println(arr2)

	// reSlice
	fmt.Println("ReSlice")
	s2 = s2[:5]
	fmt.Println(s2)
	s2 = s2[2:]
	fmt.Println(s2)

	fmt.Println("Extending slice")
	arr2[0], arr2[2] = 0, 2
	fmt.Println("arr=", arr2)
	s1 = arr2[2:6]
	s2 = s1[3:5] // slice可以向后扩展，不可以向前扩展
	fmt.Println("s1=", s1)
	fmt.Printf("len(s1)=%d, cap(s1)=%d\n", len(s1), cap(s1))
	fmt.Println("s2=", s2)
	fmt.Printf("len(s2)=%d, cap(s2)=%d\n", len(s2), cap(s2))

	s3 := append(s2, 10)
	s4 := append(s3, 11)
	s5 := append(s4, 12)
	fmt.Println("s3,s4,s5 = ", s3, s4, s5)   // s4 , s5 no longer the view of arr2.
	fmt.Println("arr =", arr2)
}
