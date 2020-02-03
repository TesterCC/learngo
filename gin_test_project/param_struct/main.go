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
*/

type Person struct {
	Name     string    `form:"name"`
	Address  string    `form:"address"`
	Birthday time.Time `form:"birthday" time_format:"2006-01-02" time_utc:"1"`    // 注意是"2006-01-02", 不设置time_utc则默认是CST，原因参考：https://www.jianshu.com/p/c7f7fbb16932 和 $GOROOT/src/time/format.go
}

// 参数绑定到结构体的演示   GET,POST都指向一个回调函数
func main() {
	r := gin.Default()
	r.GET("/testing", testing)
	r.POST("/testing", testing)
    r.Run()
}

func testing(c *gin.Context) {
	var person Person
	// 这里是根据请求的content-type来做不同的binding操作
	// 如果要获取error信息，需要一个变量接收

	//如果想在函数里面改变结构体数据内容，需要传入指针
	if err := c.ShouldBind(&person); err == nil {
		c.String(200, "%v", person)
	} else {
		c.String(500, "person bind error: %v", err)
	}

}
