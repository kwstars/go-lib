package main

import (
	"fmt"
	"gorm.io/gorm/clause"
	"gorm/common"
)

func main() {
	common.Conn()

	//批量插入
	var users = []common.UserDemo{{Name: "jinzhu1"}, {Name: "jinzhu2"}, {Name: "jinzhu3"}}
	common.DB.Create(&users)
	for _, user := range users {
		fmt.Println(user.ID) // 1,2,3  返回主键id
	}

	//根据 Map 创建
	common.DB.Model(&common.UserDemo{}).Create(map[string]interface{}{
		"Name": "jinzhu", "Age": 18,
	})

	//根据 `[]map[string]interface{}{}` 批量插入
	common.DB.Model(&common.UserDemo{}).Create([]map[string]interface{}{
		{"Name": "jinzhu_1", "Age": 18},
		{"Name": "jinzhu_2", "Age": 20},
	})

	//upinsert
	u := common.UserDemo{
		ID:   1,
		Name: "111111111",
		Age:  111111,
	}
	if err := common.DB.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "id"}},
		DoUpdates: clause.AssignmentColumns([]string{"name", "age"}),
	}).Create(&u).Error; err != nil {
		fmt.Println(err)
	}

}
