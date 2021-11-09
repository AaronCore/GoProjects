package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func main() {
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       "root:123123@tcp(127.0.0.1:3306)/go_sample?charset=utf8&parseTime=True&loc=Local", // DSN data source name
		DefaultStringSize:         256,                                                                               // string 类型字段的默认长度
		DisableDatetimePrecision:  true,                                                                              // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,                                                                              // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,                                                                              // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false,                                                                             // 根据当前 MySQL 版本自动配置
	}), &gorm.Config{
		SkipDefaultTransaction:                   false,
		DisableForeignKeyConstraintWhenMigrating: true,
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "t_",
			SingularTable: false,
		},
	})
	if err != nil {
		fmt.Println("open mysql failed,", err)
	}

	sqlDB, err := db.DB()
	sqlDB.SetConnMaxIdleTime(10)
	sqlDB.SetMaxOpenConns(200)

	//err = db.AutoMigrate(&models.User{})
	//if err != nil {
	//	fmt.Println("db.AutoMigrate failed,", err)
	//}

	//m := db.Migrator()
	//if m.HasTable(&models.User{}) {
	//	m.DropTable(&models.User{})
	//} else {
	//	m.CreateTable(&models.User{})
	//}

	fmt.Println(db, err)
}
