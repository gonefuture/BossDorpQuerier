package main

import (
	"encoding/json"
	"fmt"
	"github.com/Luxurioust/excelize"
	"log"
	"strconv"
)

type ItemReward struct {
	Amount int `json:"amount"`
	Code int `json:"code"`
	//Type int  `json:"type"`
}

type ItemResource struct {
	Key int `int:"key"`
	Name string `string:"name"`
}


func main() {

	var data = `{"amount":1, "code":1, "type":"ITEM"}`

	var itemReward ItemReward
	if err := json.Unmarshal([]byte(data), &itemReward); err == nil {
		fmt.Println("================json str 转struct==")
		fmt.Println(itemReward)
		fmt.Println(itemReward.Code)
	}


	EXCEL_FILE_NAME := "./ItemResource.xlsx"

	f, err := excelize.OpenFile(EXCEL_FILE_NAME)
	if err != nil {
		log.Fatal(err)
	}

	// 读取某个表单的所有数据
	sheetMap := f.GetSheetMap()

	itemResourceStorage := make(map[int] ItemResource)

	for _, sheetName := range sheetMap {
       rows,err := f.GetRows(sheetName)
       if err != nil {
           log.Fatal(err)
       }
       for index, row := range rows {
       		// 过滤标题栏
		   if index <= 4 {
			   continue
		   }
		   itemResource := ItemResource{}
		   itemResource.Key, _ = strconv.Atoi(row[1])
		   itemResource.Name = row[2]
		   itemResourceStorage[itemResource.Key] = itemResource
       }
    }
	for k,v := range itemResourceStorage{
		println(k)
		println(v.Name)
	}

	println(itemResourceStorage[itemReward.Code].Name)

}




