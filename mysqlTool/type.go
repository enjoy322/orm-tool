package mysqlTool

// mysql字段类型和golang对应关系
var mysqlToGo = map[string]string{
	"int":     "int",
	"integer": "int",

	"tinyint": "int8",

	"json": "[]byte",

	"smallint":           "int16",
	"mediumint":          "int",
	"bigint":             "int",
	"int unsigned":       "uint32",
	"integer unsigned":   "int",
	"tinyint unsigned":   "int",
	"smallint unsigned":  "int",
	"mediumint unsigned": "int",
	"bigint unsigned":    "int",
	"bit":                "int",
	"bool":               "bool",
	"enum":               "string",
	"set":                "string",
	"varchar":            "string",
	"char":               "string",
	"tinytext":           "string",
	"mediumtext":         "string",
	"text":               "string",
	"longtext":           "string",
	"blob":               "string",
	"tinyblob":           "string",
	"mediumblob":         "string",
	"longblob":           "string",
	"date":               "string",
	"datetime":           "string",
	"timestamp":          "string",
	"time":               "string",
	"float":              "float64",
	"double":             "float64",
	"decimal":            "float64",
	"binary":             "string",
	"varbinary":          "string",
}
