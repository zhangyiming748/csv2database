package csv

import (
	"testing"
)

// go test -v -timeout 1m -run TestParseCSV
func TestParseCSV(t *testing.T) {
	fp := "C:\\Users\\zen\\Github\\csv2database\\其它服务采购角色V2_20250507.csv"
	records, err := ParseCSV(fp)
	if err != nil {
		t.Errorf("Error parsing CSV: %v", err)
	}
	t.Logf("Parsed %+v records", records)
}
