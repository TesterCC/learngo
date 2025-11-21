package main

import "net/http"

/*
golang用函数还是用接口？
https://www.bilibili.com/video/BV1MSCDBPEGn

1.接口从理解和实现层面来说都更为复杂，为什么不直接使用函数？
函数 VS 接口

如果接口里只有一个函数时，确实可以用函数平替掉这个接口
如果接口里有多个函数时，就无法被函数平替，更容易扩展

*/

func Hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello, go world"))
}

func main() {
	// Hello是函数，第1种用HandleFunc定义路由，直接传一个函数即可
	http.HandleFunc("/1", Hello)
	// http.HandlerFunc是接口，本质上还是去执行Hello函数
	// 标准库通过type关键字，将http.HandlerFunc定义为一种类型，使得它可以拥有自己的成员方法ServeHTTP
	// 此时就实现了 http.Handle 中的 Handler这个接口，就可以作为Handle函数的第二个参数
	// 第2种用Handle定义路由，需要传一个接口。下面的代码实现是给函数hello外面套了一层 http.HandlerFunc(...),转为接口。
	http.Handle("/2", http.HandlerFunc(Hello))
	// default 0.0.0.0:5678
	http.ListenAndServe(":5678", nil)
}

/*
testcase:
curl http://127.0.0.1:5678/1
curl http://127.0.0.1:5678/2
*/
