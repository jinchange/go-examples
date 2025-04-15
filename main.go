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
	"fmt"
	"net/http"
	"time"
)

func main() {
	// 创建一个监听 8000 端口的 HTTP 服务器
	http.ListenAndServe(":8000", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		// 这会打印到 STDOUT 以表明处理已经开始
		fmt.Println("processing request")
		// 我们使用 select 来处理事件，事件取决于哪个通道先接收到消息
		select {
		case <-time.After(2 * time.Second):
			// 如果我们在2秒后收到消息，表示请求已经被处理
			w.Write([]byte("request processed"))
		case <-ctx.Done():
			// 如果请求被取消，将其记录到STDERR
			fmt.Println("request cancelled")
		}
	}))
}
