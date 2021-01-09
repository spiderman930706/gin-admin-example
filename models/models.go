package models

import (
	"log"
	"os"

	"go-gin-example/pkg/setting"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func MysqlInit() {
	var (
		err                          error
		dbName, user, password, host string
	)

	sec, err := setting.Cfg.GetSection("database")
	if err != nil {
		log.Fatal(2, "Fail to get section 'database': %v", err)
	}

	dbName = sec.Key("NAME").String()
	user = sec.Key("USER").String()
	password = sec.Key("PASSWORD").String()
	host = sec.Key("HOST").String()
	//tablePrefix = sec.Key("TABLE_PREFIX").String()

	dsn := user + ":" + password + "@tcp(" + host + ")/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"
	mysqlConfig := mysql.Config{
		DSN:                       dsn,   // DSN data source name
		DefaultStringSize:         191,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据版本自动配置
	}
	if db, err = gorm.Open(mysql.New(mysqlConfig), &gorm.Config{}); err != nil {
		log.Printf("MySQL启动异常 %s", err)
		os.Exit(0)
	} else {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(10)
		sqlDB.SetMaxOpenConns(100)
	}
}

//func init1() {
//	var (
//		err error
//	)
//
//	db, err := gorm.Open(oracle.Open("TULIN/930706@192.168.0.112:1521/XE"), &gorm.Config{})
//	if err != nil {
//		log.Println(err)
//	}
//	sqlDB, _ := db.DB()
//	sqlDB.SetMaxIdleConns(10)
//	sqlDB.SetMaxOpenConns(100)
//
//	MysqlTables(db)
//}

//MysqlTables 注册数据库表专用
func MysqlTables(db *gorm.DB) {
	err := db.AutoMigrate(
		Article{},
		Tag{},
	)
	if err != nil {
		log.Println(err)
		os.Exit(0)
	}
	log.Println("register table success")
}
