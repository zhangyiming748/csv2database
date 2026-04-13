package csv

import (
	"testing"
)

// go test -v -timeout 1m -run TestParseCSV
func TestParseCSV(t *testing.T) {
	fp := "C:\\Users\\zhang\\Github\\csv2database\\csv\\导出权限_20260413074548-昆仑集团含万方不含原生态.csv"
	records, err := ParseCSV(fp)
	if err != nil {
		t.Errorf("Error parsing CSV: %v", err)
	}
	t.Logf("Parsed %+v records", records)
}
