package config

import (
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"
)

func OpenConnection() *gorm.DB {
	config := viper.New()
	config.SetConfigFile("test_.env")
	config.AddConfigPath(".")
	err := config.ReadInConfig()
	if err != nil {
		panic(err)
	}
	ENV := config.GetString("MYSQL_URL")

	dialect := mysql.Open(ENV)
	db, err := gorm.Open(dialect, &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		panic(err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}

	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetConnMaxLifetime(30 * time.Minute)
	sqlDB.SetConnMaxIdleTime(5 * time.Minute)

	return db

}
