package models

import (
	beego "github.com/beego/beego/v2/server/web"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// DB 是 gorm.DB 类型，向外暴露
var DB *gorm.DB
var err error
var AdminPath string

func init() {
	mysqlAddr, _ := beego.AppConfig.String("mysql_addr")
	mysqlPort, _ := beego.AppConfig.String("mysql_port")
	mysqlAdmin, _ := beego.AppConfig.String("mysql_admin")
	mysqlPwd, _ := beego.AppConfig.String("mysql_pwd")
	mysqlDb, _ := beego.AppConfig.String("mysql_db")

	dsn := mysqlAdmin + ":" + mysqlPwd + "@tcp(" + mysqlAddr + ":" + mysqlPort + ")/" + mysqlDb + "?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	// ...
	AdminPath, _ = beego.AppConfig.String("admin_path")
}
