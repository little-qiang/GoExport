package resource

import (
	"bytes"
	"fmt"
	"strings"

	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type MysqlRs struct {
	Db *sql.DB
}

func NewMysqlRs(conf RsConf) (*MysqlRs, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		conf.Username,
		conf.Password,
		conf.Host,
		conf.Port,
		conf.Database)
	db, _ := sql.Open("mysql", dsn)
	return &MysqlRs{db}, nil
}

func (rs MysqlRs) GetData(strSql string) ([]map[string]string, error) {
	rows, err := rs.Db.Query(strSql)
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

func (rs MysqlRs) WriteData(targetName string, cols []string, records []map[string]string) error {
	var (
		strSql, strVals bytes.Buffer
	)
	strCols := strings.Join(cols, ",")
	l := len(records)

	for i, row := range records {
		strVals.WriteString("( ")
		for _, col := range cols {
			strVals.WriteString("'")
			strVals.WriteString(row[col])
			strVals.WriteString("', ")
		}
		strVals.Truncate(strVals.Len() - 2)
		strVals.WriteString(" ), ")
		if i%2 == 1 || i == l {
			strVals.Truncate(strVals.Len() - 2)
			strSql.WriteString("INSERT INTO ")
			strSql.WriteString(targetName)
			strSql.WriteString(" (")
			strSql.WriteString(strCols)
			strSql.WriteString(" ) VALUES ")
			strSql.WriteString(strVals.String())
			strSql.WriteString(";")

			if rt, err := rs.Db.Exec(strSql.String()); err != nil {
				return err
			} else {
				affectedRows, _ := rt.RowsAffected()
				lastInsertId, _ := rt.LastInsertId()
				fmt.Sprintf("affected_rows: %d; last_insert_id: %d", affectedRows, lastInsertId)
				strSql.Reset()
				strVals.Reset()
			}
		}
	}
	return nil
}
