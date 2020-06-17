package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

/*
Gin实现商品详情页
1.搭建静态HTML服务
2.实现AJAX服务器部分
3.实现AJAX HTML部分

命令行运行，否则会找不到static文件夹
cd ~/gin_test_project/detail_page
go build -o router_static && ./router_static

访问：http://127.0.0.1:8080/html/
*/

// 伪装访问数据库（实际应该通过ORM去数据库取数据）
type photo struct {
	Id        string
	ImagePath string
}

// 返回3项数据，photo，title，price
func detailGet(c *gin.Context) {
	// 假装处理商品id
	id := c.Param("id")
	fmt.Printf("id is: %s\n", id)

	// 假装返回数据
	// 初始化photo
	buffer := []photo{
		photo{
			Id:        "pic1",
			ImagePath: "img01.png",
		},
		photo{
			Id:        "pic2",
			ImagePath: "img02.png",
		},
		photo{
			Id:        "pic3",
			ImagePath: "img03.png",
		},
		photo{
			Id:        "pic4",
			ImagePath: "img04.png",
		},
		photo{
			Id:        "pic5",
			ImagePath: "img05.png",
		},
		photo{
			Id:        "pic6",
			ImagePath: "img06.png",
		},
	}

	// 初始化price
	price := 139900
	intPrice := price / 100
	decPrice := price % 100

	priceString := fmt.Sprintf("%d.%02d", intPrice, decPrice)


	// 返回JSON
	c.JSON(http.StatusOK, gin.H{
		"photos": buffer,
		"title": "测试Gin详情页实现",
		"price": priceString,
	})
}

func main() {
	router := gin.Default()
	// 设定html文件夹路由，用于存放index.html以及css/js等文件
	router.Static("/html", "./html")

	// 路由，设定 /detail/:id ，返回json数据
	router.GET("/detail/:id", detailGet)    // http://127.0.0.1:8999/detail/3

	// 服务启动
	router.Run(":8999")
}
