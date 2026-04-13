package csv

import (
	"encoding/csv"
	"os"
)

type Record struct {
	Username       string // 用户名
	FullName       string // 全名
	Role           string // 角色
	Type           string // 类型
	AssignmentType string // 分配类型
	Assignment     string // 分配
	StartDate      string // 开始日期
	EndDate        string // 结束日期
	ShortRoleDesc  string // 简短角色描述
}

func ParseCSV(filePath string) ([]Record, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)

	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	var parsedRecords []Record

	for i, record := range records {
		if i == 0 {
			continue
		}

		if len(record) < 10 {
			continue
		}

		// 跳过空白行（所有字段都为空）
		isEmpty := true
		for _, field := range record {
			if field != "" {
				isEmpty = false
				break
			}
		}
		if isEmpty {
			continue
		}

		parsedRecords = append(parsedRecords, Record{
			Username:       record[1],
			FullName:       record[2],
			Role:           record[3],
			Type:           record[4],
			AssignmentType: record[5],
			Assignment:     record[6],
			StartDate:      record[7],
			EndDate:        record[8],
			ShortRoleDesc:  record[9],
		})
	}

	return parsedRecords, nil
}
