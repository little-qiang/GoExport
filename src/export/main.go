package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"strings"

	"lib/resource"
)

func main() {
	//解析命令行参数
	configPath := flag.String("c", "./conf/conf.json", "conf path")
	source := flag.String("i", "", "data source")
	target := flag.String("o", "", "data target")
	flag.Parse()
	//读取配置文件字符串
	strJson, err := ioutil.ReadFile(*configPath)
	if err != nil {
		fmt.Println(err)
		return
	}
	var (
		mapJson            map[string]map[string]resource.RsConf
		rsSource, rsTarget resource.Resource
		dql, tableName     string
		columns            []string
	)
	//解析Json
	if err = json.Unmarshal(strJson, &mapJson); err != nil {
		fmt.Println(err)
		return
	}
	//获得读取数据类
	switch *source {
	case "mysql":
		rsSource = resource.NewMysqlRs(mapJson["source"]["mysql"])
		dql = fmt.Sprintf("select %s from %s limit 2",
			strings.Join(mapJson["source"]["mysql"].Columns, ","),
			mapJson["source"]["mysql"].Tablename)
	case "xlsx":
		rsSource = new(resource.XlsxRs)
		dql = mapJson["source"]["xlsx"].Tablename
	default:
		fmt.Println("error source")
		return
	}
	//获得写入数据类
	switch *target {
	case "mysql":
		rsTarget = resource.NewMysqlRs(mapJson["target"]["mysql"])
	case "xlsx":
		rsTarget = new(resource.XlsxRs)
	default:
		fmt.Println("error target")
		return
	}
	//读取数据
	records, err := rsSource.GetData(dql)
	if err != nil {
		fmt.Println(err)
		return
	}

	tableName = mapJson["target"][*target].Tablename
	columns = mapJson["target"][*target].Columns
	//写入数据
	err = rsTarget.WriteData(tableName, columns, records)
	if err != nil {
		fmt.Println(err)
		return
	}

}
