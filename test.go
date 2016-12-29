package main

import (
	"fmt"
	"time"

	"gopkg.in/mgo.v2-unstable/bson"

	m "./logoperation"
)

func main() {
	logser, err1 := m.NewLogServer("mongodb://192.168.1.41:27017")

	content := m.M{
		"info": "3",
		"msg":  18,
	}

	logmodel := &m.Log{"xxx", content, "333", time.Now().String()}

	if err1 == nil {
		err := logser.Writelog("log", "test", logmodel)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("OK")
		}
	} else {
		fmt.Println(err1)
	}
	fmt.Println("--------------------")
	where := bson.M{
		"content.info": "3",
		"content.msg":  18,
	}
	log, err2 := logser.ReadlogALL("log", "test", where)
	if err2 != nil {
		fmt.Println(err2)
	} else {
		fmt.Println(log)
	}
}
