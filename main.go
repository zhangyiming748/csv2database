package main

import (
	"fmt"
	"log"

	"csv2database/csv"
	"csv2database/sqlite"
)

func main() {
	// 1. 创建数据库连接
	sqlite.SetSqlite("erp.db")

	// 自动迁移表结构
	db := sqlite.GetSqlite()
	err := db.AutoMigrate(&sqlite.PermissionBefore20260413{})
	if err != nil {
		log.Fatalf("数据库表迁移失败: %v", err)
	}
	fmt.Println("数据库表结构初始化完成")

	// 2. 读取并解析指定csv文件
	csvFile := "C:\\Users\\zhang\\Github\\csv2database\\csv\\导出权限_20260413074548.csv"
	records, err := csv.ParseCSV(csvFile)
	if err != nil {
		log.Fatalf("解析CSV文件失败: %v", err)
	}
	fmt.Printf("成功解析CSV文件，共 %d 条记录\n", len(records))

	// 3. 批量插入记录（每批1000条）
	batchSize := 1000
	totalRecords := len(records)
	successCount := 0
	failCount := 0

	fmt.Printf("开始批量导入数据，共 %d 条记录...\n", totalRecords)

	for i := 0; i < totalRecords; i += batchSize {
		end := i + batchSize
		if end > totalRecords {
			end = totalRecords
		}

		// 准备当前批次的数据
		batch := make([]sqlite.PermissionBefore20260413, 0, end-i)
		for j := i; j < end; j++ {
			record := records[j]
			permission := sqlite.PermissionBefore20260413{
				Username:       record.Username,
				FullName:       record.FullName,
				Role:           record.Role,
				Type:           record.Type,
				AssignmentType: record.AssignmentType,
				Assignment:     record.Assignment,
				StartDate:      record.StartDate,
				EndDate:        record.EndDate,
				ShortRoleDesc:  record.ShortRoleDesc,
			}
			batch = append(batch, permission)
		}

		// 批量插入
		err := sqlite.InsertPermissions(batch)
		if err != nil {
			log.Printf("插入第 %d-%d 条记录失败: %v", i+1, end, err)
			failCount += len(batch)
		} else {
			successCount += len(batch)
			fmt.Printf("已导入 %d/%d 条记录\n", successCount, totalRecords)
		}
	}

	fmt.Printf("数据导入完成！成功: %d 条，失败: %d 条\n", successCount, failCount)
}
