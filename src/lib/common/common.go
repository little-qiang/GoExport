package common

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type ConfRoot struct {
	Source ConfSon `json:"source"`
	Target ConfSon `json:"target"`
}

type ConfSon struct {
	Mysql MysqlConf `json:"mysql"`
}

type MysqlConf struct {
	Host      string   `json:"host"`
	Port      string   `json:"port"`
	Username  string   `json:"username"`
	Password  string   `json:"password"`
	Database  string   `json:"database"`
	Tablename string   `json:"tablename"`
	Columns   []string `json:"columns"`
}

func GetDb(conf MysqlConf) (*sql.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		conf.Username,
		conf.Password,
		conf.Host,
		conf.Port,
		conf.Database)
	return sql.Open("mysql", dsn)
}
