package main

import (
	"errors"
	"fmt"
	"go-gorm/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"time"
)

var GlobalDb *gorm.DB

// https://gorm.io/zh_CN/docs/
func main() {
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       "root:123123@tcp(127.0.0.1:3306)/go_sample?charset=utf8&parseTime=True&loc=Local", // DSN data source name
		DefaultStringSize:         256,                                                                               // string 类型字段的默认长度
		DisableDatetimePrecision:  true,                                                                              // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,                                                                              // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,                                                                              // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false,                                                                             // 根据当前 MySQL 版本自动配置
	}), &gorm.Config{
		SkipDefaultTransaction:                   true,
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

	GlobalDb = db

	//CreateUserTable()
	//Add()
	//Find()
	//Update()
	//Delete()
	Exec()
	//fmt.Println(db, err)
}

// CreateUserTable 创建表
func CreateUserTable() {
	err := GlobalDb.AutoMigrate(&models.User{})
	if err != nil {
		return
	}
}

// Add 添加
func Add() {
	//user := models.User{
	//	Model: models.Model{
	//		UUID: "102",
	//	},
	//	Name:     "A102",
	//	Age:      25,
	//	Birthday: time.Now(),
	//}
	users := &[]models.User{
		{
			Model: models.Model{
				UUID: "500",
			},
			Name:     "B500",
			Age:      28,
			Birthday: time.Now(),
		},
		{
			Model: models.Model{
				UUID: "501",
			},
			Name:     "B501",
			Age:      27,
			Birthday: time.Now(),
		},
	}
	dbRes := GlobalDb.Create(&users)
	fmt.Println(dbRes.Error, dbRes.RowsAffected)
	if dbRes != nil {
		fmt.Println("add failed...")
	}
}

// Find 查询
func Find() {
	var user models.User
	var users []models.User

	//dbRes := GlobalDb.First(&user)
	//dbRes := GlobalDb.Last(&user)
	//dbRes := GlobalDb.Take(&user)
	//dbRes := GlobalDb.First(&user,"102")
	//dbRes := GlobalDb.First(&user, map[string]interface{}{
	//	"name": "A101",
	//})
	//dbRes := GlobalDb.First(&user, models.User{
	//	Name: "A102",
	//})
	//dbRes := GlobalDb.Where("name = ?","B500").First(&user)

	//result := map[string]interface{}{}
	//GlobalDb.Model(&models.User{}).First(&result)

	//result := map[string]interface{}{}
	//GlobalDb.Table("t_users").Take(&result)

	dbRes := GlobalDb.Find(&users)
	//dbRes := GlobalDb.Select("name").Find(&users)

	fmt.Println(errors.Is(dbRes.Error, gorm.ErrRecordNotFound))
	fmt.Println("user：", user)
	fmt.Println("users：", users)
}

// Update 更新
func Update() {
	var user models.User
	//var users []models.User
	//dbRes := GlobalDb.Model(&models.User{}).Where("name LIKE ?", "%abc%").Update("name", "666")

	//dbRes := GlobalDb.Where("name = ?", "666").Find(&users)
	//for k := range users {
	//	users[k].Name = "555"
	//}
	//dbRes.Save(&users)

	dbRes := GlobalDb.First(&user).Updates(models.User{Name: "A1"})
	//dbRes := GlobalDb.First(&user).Updates(map[string]interface{}{"name": "", "age": 0})
	fmt.Println(errors.Is(dbRes.Error, gorm.ErrRecordNotFound))
}

// Delete 更新
func Delete() {
	var user models.User
	//var users []models.User
	dbRes := GlobalDb.Unscoped().Where("name =?", "").Delete(&user)
	fmt.Println(errors.Is(dbRes.Error, gorm.ErrRecordNotFound))
}

func Raw() {
	var user models.User
	dbRes := GlobalDb.Raw("SELECT * FROM t_users WHERE `uuid` = ?", "101").Scan(&user)
	fmt.Println(errors.Is(dbRes.Error, gorm.ErrRecordNotFound))
	fmt.Println(user)
}

func Exec() {
	dbRes := GlobalDb.Exec("UPDATE t_users SET email = ? WHERE `uuid` = ?", "100@qq.com", "101")
	fmt.Println(errors.Is(dbRes.Error, gorm.ErrRecordNotFound))
}
