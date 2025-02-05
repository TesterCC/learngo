package type_test

import "testing"

/*

## 类型转化
与其他主要编程语言的差异
- Go语言不允许隐式类型转换
- 别名和原有类型也不能进行隐式类型转换

*/

type MyInt int64   //别名


func TestImplict(t *testing.T) {
	var a int = 1
	var b int64
	//b = a   //cannot use a (type int) as type int64 in assignment
	b = int64(a)
	var c MyInt
	//c = b   //cannot use b (type int64) as type MyInt in assignment
	c = MyInt(b)
	t.Log(a, b, c)
}

func TestPoint(t *testing.T){
	a:=1
	aPtr := &a  //get a point
	//aPtr = aPtr+1   //不支持指针运算 //invalid operation: aPtr + 1 (mismatched types *int and int)
	t.Log(a, aPtr, &a)
	t.Logf("%T %T",a, aPtr)   // get var type
}

//string是值类型，其默认的初始化值为空字符串，而不是nil
func TestString(t *testing.T){
	var s string
	t.Log("*"+s+"*")
	t.Log(len(s))   // 可以判断 s==""
}
