package main

import (
	"Gin_study/model"
	"Gin_study/router"
)

/*
待办事项：
1.增加
2.删除
3.修改
4.查看
*/

func main() {
	model.SetupDB()
	router.SetupRouters()

}
