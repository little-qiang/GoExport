package resource

import (
	"github.com/tealeg/xlsx"
)

type XlsxRs struct {
}

func (rs XlsxRs) GetData(targetName string) ([]map[string]string, error) {
	file, err := xlsx.OpenFile(targetName)
	if err != nil {
		return nil, err
	}
	var cols []string
	records := make([]map[string]string, 0)

	for _, sheet := range file.Sheets {
		for i, row := range sheet.Rows {
			//第一次循环，生成列名
			if i == 0 {
				cols = make([]string, len(row.Cells), len(row.Cells))
				for j, cell := range row.Cells {
					cols[j] = cell.String()
				}
				continue
			}
			record := make(map[string]string)
			for j, cell := range row.Cells {
				record[cols[j]] = cell.String()
			}
			records = append(records, record)
		}
	}
	return records, nil
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
