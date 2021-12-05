package main

import (
	"fmt"
	"github.com/fvbock/endless"
	"log"
	"syscall"

	"go-gin-example2/src/gin-blog/pkg/setting"
	"go-gin-example2/src/gin-blog/routers"
)

func main() {

	endless.DefaultReadTimeOut = setting.ReadTimeout
	endless.DefaultWriteTimeOut = setting.WriteTimeout
	endless.DefaultMaxHeaderBytes = 1 << 20
	endPoint := fmt.Sprintf(":%d", setting.HTTPPort)

	server := endless.NewServer(endPoint, routers.InitRouter())
	server.BeforeBegin = func(add string) {
		log.Printf("Actual pid is %d", syscall.Getpid())
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Printf("server err: %v", err)
	}

	// router := gin.Default()
	// router.GET("/test", func(c *gin.Context) {
	// 	c.JSONP(200, gin.H{
	// 		"message": "test",
	// 	})
	// })

	//router := routers.InitRouter()
	//
	//s := &http.Server{
	//	Addr: fmt.Sprintf(":%d", setting.HTTPPort),
	//	Handler: router,
	//	ReadTimeout: setting.ReadTimeout,
	//	WriteTimeout: setting.WriteTimeout,
	//	MaxHeaderBytes: 1 << 20,
	//}
	//
	//s.ListenAndServe()

}
