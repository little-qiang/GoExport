package source

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func GetData(db *sql.DB, strSql string) ([]map[string]string, error) {
	rows, err := db.Query(strSql)
	if err != nil {
		return nil, err
	}
	cols, err := rows.Columns()
	if err != nil {
		return nil, err
	}

	scans := make([]interface{}, len(cols))
	values := make([]sql.RawBytes, len(cols))
	records := make([]map[string]string, 0)

	for i := range scans {
		scans[i] = &values[i]
	}
	for rows.Next() {
		if err := rows.Scan(scans...); err != nil {
			return nil, err
		}
		record := make(map[string]string)
		for i, v := range values {
			record[cols[i]] = string(v)
		}
		records = append(records, record)
	}
	return records, nil
}
