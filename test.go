package zhourui_test

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func GetStr() string {
	dsn := "root:1astWeekend@tcp(127.0.0.1)/test?charset=utf8mb4&parseTime=True&loc=Local"
	mysqlConfig := mysql.Config{
		DSN:               dsn, // DSN data source name
		DefaultStringSize: 191, // string 类型字段的默认长度
	}

	db, err := gorm.Open(mysql.New(mysqlConfig))
	if err != nil {
		fmt.Printf("err:%+v \n", err)
	}
	fmt.Println(db)
	return "aaa"
}
