package controller

import (
	"dataPlatform/myredis"
	. "dataPlatform/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func Bootstraphtml(c *gin.Context) {
	c.HTML(http.StatusOK, "bootstrap.html", gin.H{
		"title": "GIN: Bootstrap布局页面",
	})
}

func SendKafka1(c *gin.Context) {
	//SendMsg(Client, "11111")

	c.JSON(200, gin.H{
		"message": "success",
	})
}
func RedisTest(c *gin.Context) {
	myredis.RedisClient.Set("aa", "redistest", 0)
	r,err:=myredis.RedisClient.Get("aa").Result()
	fmt.Println(r,"--",err)
}
func SendKafka(c *gin.Context) {
	data, err := c.GetRawData()
	if err != nil {
		fmt.Println(err.Error())
	}

	msgs := strings.Split(string(data), "\n")
	for _, msg := range msgs {
		SendMsg(Client, msg)
	}

	c.JSON(200, gin.H{
		"message": "success",
	})
}
