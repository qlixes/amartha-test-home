package database

import (
	"fmt"
	"log"

	"github.com/qlixes/amartha/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewMysql(cfg *config.Db) *gorm.DB {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.User,
		cfg.Pass,
		cfg.Host,
		cfg.Port,
		cfg.Name,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("Database disconnected, %s \n", err.Error())
	}

	log.Println("Database connected.")

	return db
}
