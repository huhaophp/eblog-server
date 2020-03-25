package main

import (
	"fmt"
	"github.com/fvbock/endless"
	"github.com/huhaophp/eblog/models"
	"github.com/huhaophp/eblog/pkg/setting"
	"github.com/huhaophp/eblog/routers"
	"log"
	"syscall"
)

func main() {
	setting.Setup()
	models.Setup()
	endless.DefaultReadTimeOut = setting.ServerSetting.ReadTimeout
	endless.DefaultWriteTimeOut = setting.ServerSetting.WriteTimeout
	endless.DefaultMaxHeaderBytes = 1 << 20
	endPoint := fmt.Sprintf(":%d", setting.ServerSetting.HttpPort)
	server := endless.NewServer(endPoint, routers.InitRouter())
	server.BeforeBegin = func(add string) {
		log.Printf("Actual pid is %d", syscall.Getpid())
	}
	err := server.ListenAndServe()
	if err != nil {
		log.Printf("Server err: %v", err)
	}
}
