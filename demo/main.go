package main

import (
	"demo/dao"
	"demo/model"
	"demo/routers"
	"fmt"
)

func initAll() {
	// 模型更新
	if err := dao.DB.AutoMigrate(&model.User{}); err != nil {
		fmt.Printf("err:%v\n", err)
		return
	}
}
func main() {
	dao.InitDB()
	initAll()
	routers.InitRouters()
	fmt.Println("helloWorld!!")
}
