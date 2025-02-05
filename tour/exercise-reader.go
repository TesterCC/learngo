package main

/*

http://127.0.0.1:3999/methods/22
https://tour.go-zh.org/methods/22

实现一个 Reader 类型，它产生一个 ASCII 字符 'A' 的无限流。

tour solution:
https://github.com/golang/tour/blob/master/solutions/readers.go
*/


import "golang.org/x/tour/reader"

type MyReader struct{}

// TODO: 给 MyReader 添加一个 Read([]byte) (int, error) 方法
func (r MyReader) Read(b []byte)(int, error){
	for i:= range b {
		b[i] = 'A'
	}
	return len(b),nil
}

func main() {
	reader.Validate(MyReader{})
}


