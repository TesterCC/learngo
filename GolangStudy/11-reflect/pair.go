package _1_reflect

/*
22-变量的内置pair结构详细说明
ref:
https://www.bilibili.com/video/BV1gf4y1r79E?p=22

变量的pair包含：type 和 value
*/
func main() {
	var a string
	// pair<static type:string, value: "where">
	a = "where"
	var allType interface{}
	allType = a  // 5:05
}
