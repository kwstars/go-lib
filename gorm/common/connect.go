package common

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

type UserDemo struct {
	ID   uint64 `gorm:"column:id;primarykey;comment:'主键'"`
	Name string
	Age  uint32
	gorm.DeletedAt
}

func Conn() {
	dsn := "root:mysql@tcp(127.0.0.1:3306)/mydemo?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}

	DB = db

	db.AutoMigrate(new(UserDemo))

	u := []*UserDemo{
		{
			ID:   1,
			Name: "2222",
			Age:  2222,
		},
		{
			ID:   4,
			Name: "4444",
			Age:  4444,
		},
		{
			ID:   5,
			Name: "5555",
			Age:  5555,
		},
	}

	if err := db.Create(u).Error; err != nil {
		fmt.Println(err)
	}
}

func Close() {
	DB.Exec("DROP TABLE user_demos")
}
