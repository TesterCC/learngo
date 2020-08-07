package main
import "fmt"

type Student struct {
	StudentID int64
	Name string
	birth string  // 只能在包内访问
}

func main()  {
	s0 := Student{}   // 其它都是默认值
	// Key:Value
	s1 := Student{
		StudentID:1,
		Name:"s1",
		birth:"19900101",
	}
	// 字段赋值顺序与结构体字段定义顺序一致，可以省略key
	s2 := Student{
		2,
		"s2",
		"19900102",
	}
	// 获取指针
	s3 := &Student{
		StudentID:3,
		Name:"s3",
		birth:"19900103",
	}
	fmt.Println(s0, s1, s2, s3)
	fmt.Println(s0.Name, s1.Name, s2.Name, s3.Name)
}

