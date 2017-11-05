package target

import (
	"bytes"
	"fmt"
	"strings"

	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func WriteData(db *sql.DB, targetName string, cols []string, records []map[string]string) error {
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

			if rt, err := db.Exec(strSql.String()); err != nil {
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
