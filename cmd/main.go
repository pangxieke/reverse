package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"reverse/config"
	"reverse/log"
	"reverse/middlewares"
	"reverse/models"
	"syscall"
)

func init() {
	if err := config.Init(); err != nil {
		panic(err)
	}
	if err := log.Init(); err != nil {
		panic(err)
	}
	if err := models.Init(); err != nil {
		panic(err)
	}
}

func main() {
	errc := make(chan error)
	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errc <- fmt.Errorf("%s", <-c)
	}()

	//开始监听
	go func() {
		port := ":8080"
		if config.Server.Port != 0 {
			port = fmt.Sprintf(":%d", config.Server.Port)
		}

		fmt.Println("server start... port", port)
		http.ListenAndServe(port, middlewares.Handler(models.NewReverseProxy()))
	}()

	// 开始运行，等待结束
	fmt.Println(" exit:", <-errc)
}
