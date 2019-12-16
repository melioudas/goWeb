package router

import (
	"dataPlatform/auth"
	. "dataPlatform/controller"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	//init Auth2
	auth.InitAuth(router)
	//auth.OAuth2(router)
	//拦截校验
	router.Use(auth.Authorize())

	router.GET("/home/bootstrap", Bootstraphtml)

	router.GET("/send", SendKafka1)
	router.POST("/send", SendKafka1)
	router.POST("/redis-set", RedisTest)
	///api/v1/collector/{productName}/{platformName}/{logType}
	router.POST("/api/v1/collector/:productName/:platformName/:logType", SendKafka)

	return router
}
