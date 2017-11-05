package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	"lib/resource"
)

func main() {
	strJson, err := ioutil.ReadFile("../../conf/conf.json")
	if err != nil {
		fmt.Println(err)
		return
	}

	var mapJson = make(map[string]map[string]resource.RsConf)
	if err = json.Unmarshal(strJson, &mapJson); err != nil {
		fmt.Println(err)
		return
	}

	rsSource, err := resource.NewMysqlRs(mapJson["source"]["mysql"])
	if err != nil {
		fmt.Println(err)
		return
	}

	sql := fmt.Sprintf("select %s from %s",
		strings.Join(mapJson["source"]["mysql"].Columns, ","),
		mapJson["source"]["mysql"].Tablename)

	records, err := rsSource.GetData(sql)
	if err != nil {
		fmt.Println(err)
		return
	}

	rsTarget, err := resource.NewMysqlRs(mapJson["target"]["mysql"])
	err = rsTarget.WriteData(mapJson["target"]["mysql"].Tablename, mapJson["target"]["mysql"].Columns, records)
	if err != nil {
		fmt.Println(err)
		return
	}

}
