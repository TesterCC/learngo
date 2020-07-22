package main

/*

http://127.0.0.1:3999/methods/23
https://tour.go-zh.org/methods/23

练习：rot13Reader
有种常见的模式是一个 io.Reader 包装另一个 io.Reader，然后通过某种方式修改其数据流。

例如，gzip.NewReader 函数接受一个 io.Reader（已压缩的数据流）并返回一个同样实现了 io.Reader 的 *gzip.Reader（解压后的数据流）。

编写一个实现了 io.Reader 并从另一个 io.Reader 中读取数据的 rot13Reader，通过应用 rot13 代换密码对数据流进行修改。

rot13Reader 类型已经提供。实现 Read 方法以满足 io.Reader。

tour solution:
https://github.com/golang/tour/blob/master/solutions/rot13.go
*/

import (
	"io"
	"os"
	"strings"
)

// 实现 rot13 代换密码方法
func rot13(b byte) byte {
	var a, z byte
	switch {
	case 'a' <= b && b <= 'z':
		a, z = 'a', 'z'
	case 'A' <= b && b <= 'Z':
		a, z = 'A', 'Z'
	default:
		return b
	}
	return (b-a+13)%(z-a+1) + a
}

type rot13Reader struct {
	r io.Reader
}

// todo
func (r rot13Reader) Read(p []byte)(n int, err error){
	n,err = r.r.Read(p)
	for i:=0; i<n; i++ {
		p[i] = rot13(p[i])
	}
	return
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")   // answer: You cracked the code!
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
