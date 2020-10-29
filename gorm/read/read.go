package main

import (
	"fmt"
	"gorm/common"
)

func main() {
	common.Conn()

	result := common.DB.Model(&common.UserDemo{}).Where("age = ?", 18)
	if result.Error != nil {
		fmt.Println(result.Error)
	}

	resp := make([]common.UserDemo, 0, result.RowsAffected)

	rows, err := result.Rows()
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()

	for rows.Next() {
		t := common.UserDemo{}
		if err := common.DB.ScanRows(rows, &t); err != nil {
			fmt.Println(err)
		}

		resp = append(resp, t)
	}

	fmt.Printf("%+v", resp)
}
