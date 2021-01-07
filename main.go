package main

import (
	"fmt"
	"net/http"

	"go-gin-example/models"
	"go-gin-example/pkg/setting"
	"go-gin-example/routers"
)

func main() {
	models.MysqlInit()
	router := routers.InitRouter()

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.HTTPPort),
		Handler:        router,
		ReadTimeout:    setting.ReadTimeout,
		WriteTimeout:   setting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe() //.Error()
}
