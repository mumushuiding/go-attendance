package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/mumushuiding/go-attendance/model"

	"github.com/mumushuiding/go-attendance/config"

	"github.com/mumushuiding/go-attendance/controller"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/attendance/hello", controller.HelloWorld)
	//配置
	var config = *config.Config
	// 启动数据库连接
	model.Setup()
	// 启动服务
	server := &http.Server{
		Addr:           fmt.Sprintf(":%s", config.Port),
		Handler:        mux,
		ReadTimeout:    time.Duration(config.ReadTimeout * int(time.Second)),
		WriteTimeout:   time.Duration(config.WriteTimeout * int(time.Second)),
		MaxHeaderBytes: 1 << 20,
	}
	log.Printf("the application start up at port%s\n", server.Addr)
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
