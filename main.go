package main

import (
	_ "dataPlatform/auth"
	. "dataPlatform/model"
	"dataPlatform/myredis"
	. "dataPlatform/router"
	. "dataPlatform/service"
)

func main() {
	//1.加载配置文件
	InitConfig("conf/conf.conf")
	//2.init加载kafka
	InitKafka()
	//3.加载redis
	myredis.InitRedis()
	//4.加载router
	engine := InitRouter()

	engine.Run(":" + ConfigParam["port"])

}
