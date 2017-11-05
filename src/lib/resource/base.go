package resource

type RsConf struct {
	Host      string   `json:"host"`
	Port      string   `json:"port"`
	Username  string   `json:"username"`
	Password  string   `json:"password"`
	Database  string   `json:"database"`
	Tablename string   `json:"tablename"`
	Columns   []string `json:"columns"`
}

type Resource interface {
	GetData(dql string) ([]map[string]string, error)
	WriteData(targetName string, cols []string, data []map[string]string) error
}
