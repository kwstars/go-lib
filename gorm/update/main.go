package main

import (
	"fmt"
	"gorm/common"
)

func main() {
	common.Conn()

	if err := common.DB.Model(&common.UserDemo{ID: 5}).Updates(map[string]interface{}{"name": "hello"}).Error; err != nil {
		fmt.Println(err)
	}

	if err := common.DB.Model(&common.UserDemo{ID: 5}).Updates(common.UserDemo{Name: "1111111"}).Error; err != nil {
		fmt.Println(err)
	}

	if err := common.DB.Model(&common.UserDemo{ID: 5}).Select("name").Updates(map[string]interface{}{"name": "hello", "age": 18}).Error; err != nil {
		fmt.Println(err)
	}
}
