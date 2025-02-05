package main

// 在 遍历数组中 过滤没有用的数据

import "fmt"

func main(){
	var list = make([]int, 10)
	for i := range list{
		list[i] = i
	}

	fmt.Printf("Generate origin list: %v \n", list)

	var le, hi = 0, len(list)-1
	for le <= hi{
		if list[le] % 2 == 0{
			list[le] = list[hi]
			hi--
		}else{
			le++
		}
		//fmt.Println(list)
	}
	list = list[:hi+1]
	fmt.Println(list)

}