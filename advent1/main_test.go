package main

import "testing"

func Test_expenseReportFixer(t *testing.T) {
	//Given
	expenseReportData := []int{1721, 979, 366, 299, 675, 1456}
	//When
	fixedReport, err := expenseReportFixer(expenseReportData)
	if err != nil {
		t.Fatal(err)
	}
	if fixedReport != 514579 {
		t.Fatal("Report is not fixed")
	}
}
