package utils

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)


func ConnectDB() *gorm.DB{
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", InitConfig().DBUser, InitConfig().DBPass, InitConfig().DBHost, InitConfig().DBPort, InitConfig().DBName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
        logrus.Fatal(err)
    }

	logrus.Info("Database : conected")
	return db
}