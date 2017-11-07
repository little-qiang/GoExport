package resource

import (
	"fmt"
	"github.com/tealeg/xlsx"
)

type XlsxRs struct {
}

func (rs XlsxRs) GetData(targetName string) ([]map[string]string, error) {
	file, err := xlsx.OpenFile(targetName)
	if err != nil {
		return nil, err
	}
	for _, sheet := range file.Sheets {
		for _, row := range sheet.Rows {
			for _, cell := range row.Cells {
				text := cell.String()
				fmt.Printf("%s\n", text)
			}
		}
	}
	return nil, nil
}

func (rs XlsxRs) WriteData(targetName string, cols []string, data []map[string]string) error {
	file := xlsx.NewFile()
	sheet, err := file.AddSheet("Sheet1")
	if err != nil {
		return err
	}
	row := sheet.AddRow()
	for _, col := range cols {
		cell := row.AddCell()
		cell.Value = col
	}

	for _, dataRow := range data {
		row := sheet.AddRow()
		for _, col := range cols {
			cell := row.AddCell()
			cell.Value = dataRow[col]
		}
	}
	return file.Save(targetName)
}
