package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"gopkg.in/go-playground/validator.v8"
	"reflect"
	"time"
)

// 2-13 gin基础：验证请求参数 - 自定义验证规则

/*
验证请求参数：
1.结构体验证
2.自定义验证
3.升级验证 - 支持多语言错误信息  (根据请求的语言进行错误信息转化)


Test:
curl -X GET "http://127.0.0.1:8888/testing?age=18&name=Lily&address=beijing"

curl -X GET "http://127.0.0.1:9999/bookable?check_in=2020-04-30&check_out=2020-06-25"

gin_scaffold课程代码：
gin入门实战 - 基础精髓: https://github.com/e421083458/hello_gin
后端代码：https://github.com/e421083458/gin_scaffold
前端代码：https://github.com/e421083458/vue-admin


validator.v8支持自定义结构体，期望定义一个validated标签，然后针对标签定义验证方法，把这个方法注册到validate验证器里面，完成测试。
Gin具体默认封装规则：https://godoc.org/gopkg.in/go-playground/validator.v8


*/


// 这里需要自定义 bookabledate
type Booking struct {
	CheckIn  time.Time `form:"check_in" binding:"required, bookabledate" time_format:"2006-01-02"`
	CheckOut time.Time `form:"check_out" binding:"required, gtfield=CheckIn" time_format:"2006-01-02"`
}



// 预约时间要比今天大，需要自定义验证
func bookableDate(v *validator.Validate, topStruct reflect.Value, currentStructOrField reflect.Value, field reflect.Value, fieldType reflect.Type, fieldKind reflect.Kind, param string) bool {

	if date, ok := field.Interface().(time.Time); ok {
		today := time.Now()
		if date.Unix() > today.Unix() {
			return true
		}

	}
	return false
}


func main() {
	r := gin.Default()

	// 注册bookableDate
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("bookabledate", bookableDate)
	}

	r.GET("/bookable", func(c *gin.Context) {
		var b Booking
		if err:=c.ShouldBind(&b);err!=nil {    // FIXME always error
			c.JSON(500,gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, gin.H{"message": "ok!", "booking": b})

	})

	bindAddres := "127.0.0.1:9999"
	r.Run(bindAddres)
}
