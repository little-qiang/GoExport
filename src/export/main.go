package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	"lib/common"
	"lib/source"
	"lib/target"
)

func main() {
	strJson, err := ioutil.ReadFile("./conf/conf.json")
	if err != nil {
		fmt.Println(err)
		return
	}

	var conf common.ConfRoot
	if err = json.Unmarshal(strJson, &conf); err != nil {
		fmt.Println(err)
		return
	}

	sourceDb, err := common.GetDb(conf.Source.Mysql)
	if err != nil {
		fmt.Println(err)
		return
	}

	sql := fmt.Sprintf("select %s from %s",
		strings.Join(conf.Source.Mysql.Columns, ","),
		conf.Source.Mysql.Tablename)

	records, err := source.GetData(sourceDb, sql)
	if err != nil {
		fmt.Println(err)
		return
	}

	targetDb, err := common.GetDb(conf.Target.Mysql)
	if err != nil {
		fmt.Println(err)
		return
	}

	target.WriteData(targetDb, conf.Target.Mysql.Tablename, conf.Target.Mysql.Columns, records)

}
