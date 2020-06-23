package main

import "fmt"

func printSlice(s []int) {
	fmt.Printf("%v,len=%d, cap=%d\n", s, len(s), cap(s))
}

func main() {
	fmt.Println("Creating Slice")
	// create slice directly
	var s []int // Zero value for slice is nil

	for i := 0; i < 100; i++ {
		printSlice(s) // 由cap可知，slice会重新分配空间
		s = append(s, 2*i+1)
	}
	fmt.Println(s)

	// other method to create slice
	s1 := []int{2, 4, 6, 8}
	printSlice(s1)

	s2 := make([]int, 16) // 指定len
	printSlice(s2)

	s3 := make([]int, 10, 32) // 指定len，cap
	printSlice(s3)

	fmt.Println("Copying Slice")
	copy(s2, s1)
	printSlice(s2)

	fmt.Println("Deleting elements from Slice")

	s2 = append(s2[:3], s2[4:]...)    // 不像python那样可以直接+，要用特殊语法
	printSlice(s2)

	//删除头尾元素
	fmt.Println("Popping from head")
	head := s2[0]
	s2 = s2[1:]
	fmt.Println(head)
	printSlice(s2)
	fmt.Println("Popping from tail")
	tail := s2[len(s2)-1]
	s2 = s2[:len(s2) -1]

	fmt.Println(tail)
	printSlice(s2)

}
