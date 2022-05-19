package database

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Connector interface {
	Open() (*gorm.DB, error)
	Close() error
}

type db struct {
	DB *gorm.DB
}

func (d *db) Open() (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		viper.Get("db_user"),
		viper.Get("db_pass"),
		viper.Get("db_host"),
		viper.Get("db_port"),
		viper.Get("db_name"),
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	d.DB = db
	return db, nil
}

func (d *db) Close() error {
	sqlDB, err := d.DB.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}

func NewDBGORM() Connector {
	return &db{}
}
