package config

import (
	"TestDevGolang/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"math/rand"
)

func InitialDB() *gorm.DB {
	dsn := "root:password@tcp(127.0.0.1:3306)/testdev?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&model.Power{})

	powerData := model.Power{}
	powerEmpty := model.Power{}
	//Find data power for check empty
	db.Find(&powerData)

	//If power is empty will generate mock data
	if powerData == powerEmpty {
		mockData(db)
	}

	return db
}

func mockData(db *gorm.DB) error {
	var powers []model.Power

	for i := 0; i < 1000; i++ {
		powers = append(powers, model.Power{
			ActivePower: random(1, 1000),
			PowerInput:  random(1, 1000),
		})
	}

	return db.Create(&powers).Error
}

func random(min int, max int) int {
	return rand.Intn(max-min) + min
}
