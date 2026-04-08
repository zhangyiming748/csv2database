package main

import (
	"fmt"
	"log"

	"csv2database/csv"
	"csv2database/sqlite"
)

func main() {
	// 1. 创建数据库连接
	sqlite.SetSqlite("roles.db")

	// 自动迁移表结构
	db := sqlite.GetSqlite()
	err := db.AutoMigrate(&sqlite.Role{})
	if err != nil {
		log.Fatalf("数据库表迁移失败: %v", err)
	}
	fmt.Println("数据库表结构初始化完成")

	// 2. 读取并解析指定csv文件
	csvFile := "其它服务采购角色V2_20250507.csv"
	records, err := csv.ParseCSV(csvFile)
	if err != nil {
		log.Fatalf("解析CSV文件失败: %v", err)
	}
	fmt.Printf("成功解析CSV文件，共 %d 条记录\n", len(records))

	// 3. 按行使用单独插入一条的方法用for循环遍历插入记录
	successCount := 0
	failCount := 0
	for i, record := range records {
		role := sqlite.Role{
			Role:            record.Role,
			RootRoleDesc:    record.RootRoleDesc,
			DerivedRole:     record.DerivedRole,
			DerivedRoleDesc: record.DerivedRoleDesc,
		}

		err := role.InsertOne()
		if err != nil {
			log.Printf("插入第 %d 条记录失败: %v", i+1, err)
			failCount++
		} else {
			successCount++
		}
	}

	fmt.Printf("数据导入完成！成功: %d 条，失败: %d 条\n", successCount, failCount)
}
