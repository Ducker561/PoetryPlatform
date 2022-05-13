package model

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

var DB *gorm.DB

func Database(connString string) {
	fmt.Println("connstring", connString)
	db, err := gorm.Open("mysql", connString)
	if err != nil {
		panic(err)
	}
	db.LogMode(true)
	if gin.Mode() == "release" {
		db.LogMode(false)
	}
	db.SingularTable(true)                       //默认不加复数s
	db.DB().SetMaxIdleConns(20)                  //设置连接池，空闲
	db.DB().SetMaxOpenConns(100)                 //最大连接数
	db.DB().SetConnMaxLifetime(time.Second * 30) //连接的生命周期
	DB = db
	migration()
}
