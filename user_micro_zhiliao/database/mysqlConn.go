package database

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"user_micro_zhiliao/models"
)

var Db *gorm.DB
var Err error

func init() {
	fmt.Println("mysql 的 init 函数执行")
	Db, Err := gorm.Open("mysql", "root:root@tcp(localhost:3306)/zhiliao?charset=utf8&parseTime=true&loc=Local")

	if Err != nil {
		panic(Err)
	}

	defer Db.Close()

	Db.AutoMigrate(&models.FrontUser{})
}
