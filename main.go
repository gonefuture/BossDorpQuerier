package main

import (
	"fmt"
	"log"
    "github.com/Luxurioust/excelize"
)

func main() {
    excelName := "./ItemResource.xlsx"

	f, err := excelize.OpenFile(excelName)
    if err != nil {
        log.Fatal(err)
    }
 
    /*
      //读取某个单元格的值
        value, err := f.GetCellValue("成绩表", "D2")
        if err != nil {
            log.Fatal(err)
        }
        fmt.Println(value)
    */
 
    // 读取某个表单的所有数据
    sheetMap := f.GetSheetMap()

    for 



    fmt.Println(sheetMap)


    // if err != nil {
    //     log.Fatal(err)
    // }
    // for _, row := range rows {
    //     for _, value := range row {
    //         fmt.Printf("\t%s", value)
    //     }
    //     fmt.Println()
    // }
}