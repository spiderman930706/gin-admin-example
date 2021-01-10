package main

import (
	"fmt"
	"go-gin-example/models"
	"net/http"

	"github.com/spiderman930706/gin_admin"
	"github.com/spiderman930706/gin_admin/config"
	"go-gin-example/pkg/setting"
	"go-gin-example/routers"
)

func main() {
	con := config.Config{
		Mysql: config.Mysql{
			DbName:   "blog",
			User:     "root",
			Password: "930706",
			Host:     "127.0.0.1",
		},
		JWT: config.JWT{
			SigningKey:   "example-key",
			ExpireSecond: 7 * 24 * 3600,
		},
	}
	router := routers.InitRouter()
	group := router.Group("admin")
	gin_admin.RegisterConfigAndRouter(con, group)
	gin_admin.RegisterTables(
		true,
		&models.User{},
		&models.Tag{},
		&models.Article{},
	)

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.HTTPPort),
		Handler:        router,
		ReadTimeout:    setting.ReadTimeout,
		WriteTimeout:   setting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe() //.Error()
}
