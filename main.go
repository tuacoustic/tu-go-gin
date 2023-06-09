package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/TUACOUSTIC/go-gin-example/pkg/setting"
	"github.com/TUACOUSTIC/go-gin-example/routers"
)

func init() {
	setting.Setup()
}
func main() {
	gin.SetMode(setting.ServerSetting.RunMode)

	routersInit := routers.InitRouter()
	readTimeout := setting.ServerSetting.ReadTimeout
	writeTimeout := setting.ServerSetting.WriteTimeout
	endPoint := fmt.Sprintf(":%d", setting.ServerSetting.HttpPort)
	maxHeaderBytes := 1 << 20

	server := &http.Server{
		Addr:           endPoint,
		Handler:        routersInit,
		ReadTimeout:    readTimeout,
		WriteTimeout:   writeTimeout,
		MaxHeaderBytes: maxHeaderBytes,
	}
	log.Printf("[info] start http server listening %s", endPoint)

	server.ListenAndServe()
	// // Khởi chạy máy chủ
	// err := server.ListenAndServeTLS("./ssl/cert.pem", "./ssl/key.pem")
	// if err != nil {
	// 	panic(err)
	// }
}
