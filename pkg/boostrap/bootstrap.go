package boostrap

import (
	"fmt"
	"os"
	"v/internal/domain"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func DBConnection() (*gorm.DB, error) {
	_ = godotenv.Load()

	user := os.Getenv("DATABASE_USER")
	password := os.Getenv("DATABASE_PASSWORD")
	host := os.Getenv("DATABASE_HOST")
	port := os.Getenv("DATABASE_PORT")
	name := os.Getenv("DATABASE_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		user,
		password,
		host,
		port,
		name)

	fmt.Println(dsn)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	if os.Getenv("DATABASE_DEBUG") == "true" {
		db = db.Debug()
	}
	if os.Getenv("DATABASE_MIGRATE") == "true" {
		if err = db.AutoMigrate(&domain.User{}); err != nil {
			return nil, err
		}
		if err = db.AutoMigrate(&domain.Course{}); err != nil {
			return nil, err
		}
		if err = db.AutoMigrate(&domain.Enrollment{}); err != nil {
			return nil, err
		}
	}
	return db, nil
}
