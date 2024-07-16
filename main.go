package main

import (
	"AvaImageServer/pkg/logging"
	"AvaImageServer/routers"
	"AvaImageServer/setting"
	"fmt"
	"log"
	"net/http"
)

func main() {
	setting.Setup()
	logging.Setup()
	router := routers.InitRouter()
	fmt.Println("启动端口：", setting.AppSetting.HttpPort)
	s := &http.Server{
		Addr:    fmt.Sprintf(":%d", setting.AppSetting.HttpPort),
		Handler: router,
	}
	if err := s.ListenAndServe(); err != nil {
		log.Printf("Listen: %s\n", err)
	}
}
