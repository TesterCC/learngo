package main

import (
	"fmt"
	"reflect"
)

/*
24-golang反射解析结构体标签Tag   key:value 形式
ref: https://www.bilibili.com/video/BV1gf4y1r79E?p=24

*/

type resume struct {
	Name string `info:"name" doc:"我的名字"`
	Sex  string `info:"sex"`
}

func findTag(str interface{}) {
	t := reflect.TypeOf(str).Elem()

	for i := 0; i < t.NumField(); i++ {
		taginfo := t.Field(i).Tag.Get("info")
		tagdoc := t.Field(i).Tag.Get("doc")
		fmt.Println("info: ", taginfo, " doc: ", tagdoc)
	}

}

func main() {
	var re resume
	findTag(&re)
}
