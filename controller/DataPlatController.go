package controller

import (
	"dataPlatform/myredis"
	. "dataPlatform/service"
	"encoding/json"
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
	r, err := myredis.RedisClient.Get("aa").Result()
	fmt.Println(r, "--", err)
}
func SendKafka(c *gin.Context) {

	ip := getIpAddr(c.Request)
	ip = strings.Split(ip, ":")[0]
	data, err := c.GetRawData()
	if err != nil {
		fmt.Println(err.Error())
	}

	msgs := strings.Split(string(data), "\n")
	for _, msg := range msgs {
		if (isJSON(msg)) {
			m := msg[0 : len(msg)-1]
			message := m + ",\"clientip\":" + "\"" + ip + "\"" + "}"
			//fmt.Println("输出：" + message)
			SendMsg(Client, message)
		} else {
			c.JSON(500, gin.H{
				"message": "it's not a json!",
			})
			return
		}
	}

	c.JSON(200, gin.H{
		"message": "success",
	})
}

func isJSON(s string) bool {
	var js map[string]interface{}
	return json.Unmarshal([]byte(s), &js) == nil
}

func getIpAddr(request *http.Request) string {

	ip := request.Header.Get("x-forward-for")

	if (ip == "" || strings.EqualFold("unknown", ip)) {
		ip = request.Header.Get("X-Forwarded-For")
	}
	if (ip == "" || strings.EqualFold("unknown", ip)) {
		ip = request.Header.Get("Proxy-Client-IP")
	}
	if (ip == "" || strings.EqualFold("unknown", ip)) {
		ip = request.Header.Get("WL-Proxy-Client-IP")
	}
	if (ip == "" || strings.EqualFold("unknown", ip)) {
		ip = request.RemoteAddr
	}
	return ip
}
