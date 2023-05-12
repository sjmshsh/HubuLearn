package dao

import (
	"HubuLearn/conf"
	"github.com/jinzhu/gorm"

	"strings"
	"time"
)

var _db *gorm.DB

func MySQLInit() {
	conn := strings.Join([]string{conf.DbUser, ":", conf.DbPassword, "@tcp(", conf.DbHost, ":", conf.DbPort, ")/", conf.DbName, "?charset=utf8&parseTime=true"}, "")
	db, err := gorm.Open(conn)
	if err != nil {
		panic(err)
	}
	sqlDB := db.DB()
	sqlDB.SetMaxIdleConns(20)  // 设置连接池，空闲
	sqlDB.SetMaxOpenConns(100) // 打开
	sqlDB.SetConnMaxLifetime(time.Second * 30)
	_db = db
	migration()
}
