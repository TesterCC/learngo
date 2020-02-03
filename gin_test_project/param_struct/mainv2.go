package main

import (
	"fmt"
	"log"
	"time"
	"github.com/gin-gonic/gin"
)

// curl -X POST "http://127.0.0.1:8888/testing" -d 'name=Lily&address=zhejiang&birthday=2020-02-04'
// ref: gin入门实战 - 基础精髓: https://github.com/e421083458/hello_gin

type Person2 struct {
	Name     string    `form:"name"`
	Address  string    `form:"address"`
	Birthday time.Time `form:"birthday" time_format:"2006-01-02" time_utc:"1"`
}

func main() {
	route := gin.Default()
	route.GET("/testing", startPage)
	route.POST("/testing", startPage)
	route.Run(":8888")
}

func startPage(c *gin.Context) {
	var person Person2
	if c.ShouldBind(&person) == nil {
		log.Println(person.Name)
		log.Println(person.Address)
		log.Println(person.Birthday)
	}
	c.String(200, fmt.Sprintf("%#v", person))
}
