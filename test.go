package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

//func main() {
//	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
//	dsn := "root:930706@tcp(127.0.0.1:3306)/blog?charset=utf8mb4&parseTime=True&loc=Local"
//	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
//	fmt.Println(err)
//	fmt.Println(db)
//}

func main() {
	//r := gin.Default()
	//r.GET("/ping", func(c *gin.Context) {
	//	c.JSON(200, gin.H{
	//		"message": "pong",
	//	})
	//})
	//r.Run() // listen and serve on 0.0.0.0:8080
	c := make(map[string]interface{})
	c["a"] = 1
	c["b"] = "b"
	b := c["a"].(int)
	a := c["b"]
	fmt.Println(a, b)

	type Auth struct {
		ID       int    `gorm:"primary_key" json:"id"`
		Username string `json:"username"`
		Password string `json:"password"`
	}

	dsn := "root:930706@tcp(127.0.0.1:3306)/blog?charset=utf8mb4&parseTime=True&loc=Local"
	mysqlConfig := mysql.Config{
		DSN:                       dsn,   // DSN data source name
		DefaultStringSize:         191,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据版本自动配置
	}
	if db, err := gorm.Open(mysql.New(mysqlConfig), &gorm.Config{}); err != nil {
		log.Printf("MySQL启动异常 %s", err)
		os.Exit(0)
	} else {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(10)
		sqlDB.SetMaxOpenConns(100)
		var auth Auth
		db.Select("id").Find(&auth)
		fmt.Println(auth)
	}

}
