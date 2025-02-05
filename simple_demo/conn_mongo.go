package main

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

/*
go mod init simple_demo
go get gopkg.in/mgo.v2
*/

func main() {
    // 没调通，最要还是用官方的mongo驱动
	session, err := mgo.Dial("10.0.4.148:27017")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	c := session.DB("asset").C("asset")

    // query  todo ref: https://blog.csdn.net/impressionw/article/details/81638157
	var result interface{}
	err = c.Find(bson.M{"admin":"tt"}).One(&result)
}