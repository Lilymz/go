package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
)

type ProductDao struct {
	db *gorm.DB
}

func (p *ProductDao) New(db *gorm.DB) *ProductDao {
	return &ProductDao{db: db}
}

// InsertBySpecific  插入指定字段的实体数据
func (p *ProductDao) InsertBySpecific(product *Product, any ...interface{}) {
	p.db.Select(any).Create(product)
}

// InsertByNotSpecific  插入指定字段的实体数据
func (p *ProductDao) InsertByNotSpecific(product *Product, columns ...string) {
	p.db.Omit(columns...).Create(product)
}

type Dao interface {
	InsertBySpecific(product *Product, any ...interface{})
	InsertByNotSpecific(product *Product, columns ...string)
}

func Connect(dsn string, config *gorm.Config) *gorm.DB {
	logger := logger.New(log.New(os.Stdout, "\r\n", log.LstdFlags), logger.Config{})
	db, err := gorm.Open(mysql.Open(dsn), config)
	if err != nil {
		logger.Error(nil, "failed to connect database\n")
		return nil
	}
	return db
}
