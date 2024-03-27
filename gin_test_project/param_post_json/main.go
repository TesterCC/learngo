package main

import (
	"github.com/gin-gonic/gin"
	"time"
)

// 2-10 gin基础： 获取请求参数 - 获取bind参数

/*
Test GET Method:
curl -X GET "http://127.0.0.1:8080/testing?name=Lily"
curl -X GET "http://127.0.0.1:8080/testing?name=Lily&address=Beijing"
curl -X GET "http://127.0.0.1:8080/testing?name=Lily&address=Beijing&birthday=2020-02-02"

Test POST Method:
curl -X POST "http://127.0.0.1:8080/testing?name=Lily&address=Beijing&birthday=2020-02-03"
curl -X POST "http://127.0.0.1:8080/testing" -d 'name=Lily&address=zhejiang&birthday=2020-02-04'

测试json
curl -H "Content-Type:application/json" -X POST "http://127.0.0.1:8080/testing" -d '{"name":"Lily","address":"shenzhen"}'

ref:
https://www.runoob.com/go/go-structures.html
https://www.jianshu.com/p/c7f7fbb16932

2006-01-02 15:04:05  另一种常用的

https://blog.csdn.net/heart66_a/article/details/100796964

ref2:
Gin框架中，有bind函数可以非常方便的将url的查询参数query parameter、http的Header，body中提交上来的数据格式，如form，json，xml等，绑定到go中的结构体中去
https://blog.csdn.net/xiaojinran/article/details/118788032

*/

// 参数绑定到结构体的演示   GET,POST都指向一个回调函数

type Person struct {
	Name     string    `form:"name"`
	Address  string    `form:"address"`
	Birthday time.Time `form:"birthday" time_format:"2006-01-02" time_utc:"1"` // 注意是"2006-01-02", 不设置time_utc则默认是CST，原因参考：https://www.jianshu.com/p/c7f7fbb16932 和 $GOROOT/src/time/format.go
}

/*
gin的BindJSON和ShouldBindJSON，ShouldBindWith的区别
1.解析错误时，会在header中写一个400的状态码
//内部根据Content-Type去解析
c.Bind(obj interface{})
//内部替你传递了一个binding.JSON，对象去解析
c.BindJSON(obj interface{})
//解析哪一种绑定的类型，根据你的选择
c.BindWith(obj interface{}, b binding.Binding)

2.解析错误直接返回，至于要给客户端返回什么错误状态码由开发者决定
//内部根据Content-Type去解析
c.ShouldBind(obj interface{})
//内部替你传递了一个binding.JSON，对象去解析
c.ShouldBindJSON(obj interface{})
//解析哪一种绑定的类型，根据你的选择
c.ShouldBindWith(obj interface{}, b binding.Binding)

Shouldxxx和bindxxx区别就是bindxxx会在head中添加400的返回信息，而Shouldxxx不会

仅接受post json，可用 BindJSON
要同时支持GET和POST，则使用Bind
*/

func main() {
	r := gin.Default()
	r.GET("/info", info)
	r.POST("/info", info)
	r.Run(":9999")
}

func info(c *gin.Context) {
	var person Person
	// 这里是根据请求的content-type来做不同的binding操作
	// 如果要获取error信息，需要一个变量接收

	//如果想在函数里面改变结构体数据内容，需要传入指针
	//if err := c.BindJSON(&person); err == nil {
	if err := c.ShouldBind(&person); err == nil {
		c.String(200, "%v", person)
	} else {
		c.String(500, "person bind error: %v", err)
	}

}
