package main

//func main() {
//	router := gin.Default()
//	// curl "http://localhost:8080/abc/zhangsan"
//	router.GET("/abc/:name", func(c *gin.Context) {
//		c.JSON(http.StatusOK, gin.H{
//			"message": c.Param("name"),
//		})
//	})
//	// curl "http://localhost:8080/efg/hij/hij.jpg"
//	router.GET("/efg/*path", func(c *gin.Context) {
//		c.JSON(http.StatusOK, gin.H{
//			"message": c.Param("path"),
//		})
//	})
//	// curl "http://localhost:8080/hij?age=18"
//	router.GET("/hij", func(c *gin.Context) {
//		c.JSON(http.StatusOK, gin.H{
//			"message": c.DefaultQuery("age", "0"),
//		})
//	})
//	// 监听并在 0.0.0.0:8080 上启动服务
//	log.Panic(router.Run())
//}

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	go handle(ctx, 2*time.Second)
	select {
	case <-ctx.Done():
		fmt.Println("main: ", ctx.Err())
	}
}

func handle(ctx context.Context, duration time.Duration) {
	select {
	case <-ctx.Done():
		fmt.Println("handle", ctx.Err())
	default:
		time.Sleep(duration)
	}
}
