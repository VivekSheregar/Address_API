package utils

import (
	"context"
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetConnection(ctx context.Context, username, password, host, port string) *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=address port=%s sslmode=disable TimeZone=Asia/Kolkata", host, username, password, port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	if err == nil {
		fmt.Println("connnected")
	}
	return db
}
