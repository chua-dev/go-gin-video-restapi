package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/v1/topics", func(c *gin.Context) {
		if c.Query("username") == "" {
			c.String(http.StatusOK, "获取帖子列表")
		} else {
			c.String(http.StatusOK, "获取一个用户名%s的帖子列表", c.Query("username"))
		}
	})

	router.GET("/v1/topics/:topic_id", func(c *gin.Context) {
		c.String(http.StatusOK, "获取单个帖子ID%s", c.Param("topic_id"))
	})

	router.Run(":8080")
}
