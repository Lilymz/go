package main

import (
	"gorm.io/gorm"
	"os"
)

var (
	product *Product
)

const (
	DSN = "root:hstest2014@tcp4(172.17.114.177:3306)/cluster_server_gj?charset=utf8mb4&parseTime=True&loc=Local"
)

func main() {
	// 连接数据库
	db := Connect(DSN, &gorm.Config{})
	product = &Product{
		Code:  "CD875",
		Price: 23,
	}
	err := db.AutoMigrate(&product)
	if err != nil {
		return
	}
	// Create
	db.Omit("CreatedAt", "UpdatedAt", "DeletedAt").Create(&product)
	//db.Select("CreatedAt","UpdatedAt","DeletedAt").Create(&product)
	os.Exit(200)
}

type Product struct {
	gorm.Model
	Code  string
	Price uint
}
