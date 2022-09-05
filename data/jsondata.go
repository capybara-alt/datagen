package data

import (
	"encoding/json"
	"fmt"

	"github.com/capybara-alt/datagen/config"
	"github.com/capybara-alt/datagen/generator"
)

type JsonData struct {
	config.Config
}

func NewJsonData(conf *config.Config) *JsonData {
	testData := new(JsonData)
	testData.Config = *conf

	return testData
}

func (j *JsonData) GetValue() ([]byte, error) {
	var b []byte
	if len(j.Config.Columns) > 0 {
		b, _ = json.Marshal(j.ColumnsToObject(j.Config.Columns))
	} else if j.Config.Rows != nil {
		list := make([]map[string]interface{}, j.Config.Rows.Size)
		for i := 0; i < len(list); i++ {
			list[i] = j.ColumnsToObject(j.Config.Rows.Columns)
		}
		b, _ = json.Marshal(list)
	}
	return b, nil
}

func (j *JsonData) ColumnsToObject(columns []config.Column) map[string]interface{} {
	obj := make(map[string]interface{})
	for _, column := range columns {
		attrs := createAttrsObject(column)
		if len(column.Columns) < 1 && column.Rows == nil {
			obj[column.Name] = attrs
		}

		if len(column.Columns) > 0 {
			obj[column.Name] = j.ColumnsToObject(column.Columns)
			for key, value := range attrs {
				obj[column.Name].(map[string]interface{})[key] = value
			}
		} else if column.Rows != nil && column.Rows.Format == "" {
			list := make([]map[string]interface{}, column.Rows.Size)
			for i := 0; i < column.Rows.Size; i++ {
				list[i] = j.ColumnsToObject(column.Rows.Columns)
				for key, value := range attrs {
					list[i][key] = value
				}
			}
			obj[column.Name] = list
		} else if column.Rows != nil && column.Rows.Format != "" {
			list := make([]string, column.Rows.Size)
			for i := 0; i < column.Rows.Size; i++ {
				list[i] = generator.Generate(column.Rows.Format)
			}
			obj[column.Name] = list
			for key, value := range attrs {
				obj[key] = value
			}
		} else if attrs == nil {
			obj[column.Name] = generator.Generate(column.Format)
		}
	}

	return obj
}

func createAttrsObject(column config.Column) map[string]string {
	if column.Attrs == nil {
		return nil
	}

	obj := make(map[string]string)
	for key, value := range column.Attrs {
		obj[fmt.Sprintf("@%s", key)] = generator.Generate(value)
	}

	if column.Rows == nil && len(column.Columns) < 1 && column.Format != "" {
		obj["#text"] = generator.Generate(column.Format)
	}

	return obj
}
