// @Desc: \\todo
// @Author: MaXiaoTian 2020/11/24 4:37 下午
// @Update: xxx 2020/11/24 4:37 下午

package main

import "github.com/gin-gonic/gin"

func getting(c *gin.Context){}
func posting(c *gin.Context){}
func putting(c *gin.Context){}
func deleting(c *gin.Context){}
func patching(c *gin.Context){}
func head(c *gin.Context){}
func options(c *gin.Context){}

func main() {
	// Creates a gin router with default middleware:
	// logger and recovery (crash-free) middleware
	router := gin.Default()

	router.GET("/someGet", getting)
	router.POST("/somePost", posting)
	router.PUT("/somePut", putting)
	router.DELETE("/someDelete", deleting)
	router.PATCH("/somePatch", patching)
	router.HEAD("/someHead", head)
	router.OPTIONS("/someOptions", options)

	// By default it serves on :8080 unless a
	// PORT environment variable was defined.
	router.Run()
	// router.Run(":3000") for a hard coded port
}