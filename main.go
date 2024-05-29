package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func GetRealIp(c *gin.Context) {
	reqIP := c.ClientIP()
	logrus.Info(reqIP)
	c.JSON(200, gin.H{"ip": reqIP})
}

func GetHeader(c *gin.Context) {
	logrus.Info(c.Request.Header)
	c.JSON(200, c.Request.Header)
}

func main() {
	r := gin.Default()
	r.GET("/test/ip", GetRealIp)
	r.GET("/test/header", GetHeader)
	r.Run("0.0.0.0:8080")
}
