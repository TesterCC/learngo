package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"html/template"
	"log"
	"net/http"
	"time"
)

// go get github.com/go-sql-driver/mysql
// 实例 5-5 服务端基于MySQL的存储方式实践 P118-119
// Go语言没有内置的驱动支持任何的数据库，但是Go语言定义了database/sql接口，用户可以基于驱动接口开发相应数据库的驱动。

type User struct {
	Id          int
	Name        string
	Habits      string
	CreatedTime string
}

var tpl = `<html>
<head>
<title></title>
</head>
<body>
<form action="/info" method="post">
	用户名:<input type="text" name="username">
	兴趣爱好:<input type="text" name="habits">
	<input type="submit" value="提交">
</form>
</body>
</html>`

var db *sql.DB

var err error

func init() {
	db, err = sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/testdb?charset=utf8")
	checkErr(err)
}

// 按照name查询记录
func queryByName(name string) User {
	user := User{}
	stmt, err := db.Prepare("select  * from tuser where name=?")
	checkErr(err)

	rows, _ := stmt.Query(name)

	fmt.Println("\nafter deleting records: ")
	// 遍历结果
	for rows.Next() {
		var id int
		var name string
		var habits string
		var createdTime string
		err = rows.Scan(&id, &name, &habits, &createdTime)
		checkErr(err)
		fmt.Printf("[%d,%s,%s,%s]\n", id, name, habits, createdTime)
		user = User{id, name, habits, createdTime}
		break // 只取第一条记录
	}
	return user // 返回查询到的记录

}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func submitForm(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method: ", r.Method) // 获取请求方法
	var t *template.Template
	t = template.New("Products") // 创建一个模板
	t, _ = t.Parse(tpl)
	log.Println(t.Execute(w, nil))
}

func store(user User) {
	// insert data
	stmt, err := db.Prepare("INSERT INTO tuser SET name=?,habits=?,created_time=?")
	t := time.Now().UTC().Format("2006-01-02")
	res, err := stmt.Exec(user.Name, user.Habits, t)
	checkErr(err)

	id, err := res.LastInsertId()
	checkErr(err)

	fmt.Printf("last insert id is: %d\n", id)

}


func userInfo(w http.ResponseWriter, r *http.Request) {
	//请求的是登录数据，那么执行登录的逻辑判断
	_ = r.ParseForm()
	if r.Method == "POST" {
		user1 := User{Name: r.Form.Get("username"), Habits: r.Form.Get("habits")}
		store(user1)
		//fmt.Fprintf(w, " %v", queryByName("abc")) // 查询指定name, 这个写入到w的是输出到客户端的
		fmt.Fprintf(w, " %v", queryByName(r.Form.Get("username"))) // 查询插入的信息, 这个写入到w的是输出到客户端的
	}
}

func main() {
	http.HandleFunc("/form", submitForm)     //设置访问的路由
	http.HandleFunc("/info", userInfo)       //设置访问的路由
	err := http.ListenAndServe(":8080", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

// http://127.0.0.1:8080/form
// http://127.0.0.1:8080/info?username=abc
