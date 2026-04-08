package csv

import (
	"encoding/csv"
	"os"
)

type Record struct {
	Role            string
	RootRoleDesc    string
	DerivedRole     string
	DerivedRoleDesc string
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

		if len(record) < 4 {
			continue
		}

		parsedRecords = append(parsedRecords, Record{
			Role:            record[0],
			RootRoleDesc:    record[1],
			DerivedRole:     record[2],
			DerivedRoleDesc: record[3],
		})
	}

	return parsedRecords, nil
}
