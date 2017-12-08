package mysql

import (
	"github.com/astaxie/beego"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// var db *gorm.DB

// func init() {
// 	url := beego.AppConfig.String("datasource_url")
// 	db = connect(url)
// }

func Connect() *gorm.DB {
	url := beego.AppConfig.String("datasource_url")
	con, err := gorm.Open("mysql", url)
	if err != nil {
		panic(err)
	}
	return con
}
