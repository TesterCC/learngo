package main

import (
	"github.com/gin-gonic/gin"
	en2 "github.com/go-playground/locales/en"
	zh2 "github.com/go-playground/locales/zh"
	"github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
)

// 2-14 gin基础：验证请求参数 - 多语言翻译验证
/*
curl -X GET "http://127.0.0.1:8080/testing"
[Age为必填字段 Name为必填字段 Address为必填字段]%

curl -X GET "http://127.0.0.1:8080/testing?name=Lily"
[Age为必填字段 Address为必填字段]

curl -X GET "http://127.0.0.1:8080/testing?name=Lily&age=9&address=beijing"
[Age必须大于10]%

curl -X GET "http://127.0.0.1:8080/testing?name=Lily&age=29&address=beijing"
{29 Lily beijing}%

curl -X GET "http://127.0.0.1:8080/testing?name=Lily&age=9&address=beijing&locale=en"
[Age must be greater than 10]
 */

type Person struct {
	Age     int    `form:"age" validate:"required,gt=10"` // v9 binding换成validate
	Name    string `form:"name" validate:"required"`
	Address string `form:"address" validate:"required"`
}

var (
	Uni      *ut.UniversalTranslator
	Validate *validator.Validate
)

// 验证信息多语言化
func main() {
	Validate = validator.New() // 创建验证器

	zh := zh2.New()
	en := en2.New()
	Uni = ut.New(zh, en) // 创建翻译器

	r := gin.Default()
	r.GET("/testing", func(c *gin.Context) {
		locale := c.DefaultQuery("locale", "zh")
		trans, _ := Uni.GetTranslator(locale)
		// 根据当前语言参数获取翻译器
		switch locale {
		case "zh":
			// 把翻译器注册到验证起中
			zh_translations.RegisterDefaultTranslations(Validate, trans)
		case "en":
			en_translations.RegisterDefaultTranslations(Validate, trans)
		default:
			zh_translations.RegisterDefaultTranslations(Validate, trans)
		}

		var person Person

		if err := c.ShouldBind(&person); err != nil {
			c.String(500, "%v", err)
			c.Abort()
			return
		}

		if err := Validate.Struct(person); err != nil {
			errs := err.(validator.ValidationErrors)
			sliceErrs := [] string{}
			for _, e := range errs {
				sliceErrs = append(sliceErrs, e.Translate(trans))
			}
			c.String(500, "%v", sliceErrs)
			c.Abort()
			return
		}
		c.String(200, "%v", person)

	})

	r.Run()
}
