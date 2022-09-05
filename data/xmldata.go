package data

import (
	"fmt"
	"strings"

	"github.com/capybara-alt/datagen/config"
	"github.com/capybara-alt/datagen/generator"
)

type XmlData struct {
	config.Config
}

type XmlEntity struct {
	NodeName string
	Attrs    map[string]string
	Value    interface{}
	Children []XmlEntity
	Column   config.Column
}

func NewXmlData(conf *config.Config) *XmlData {
	xmlData := new(XmlData)
	xmlData.Config = *conf

	return xmlData
}

func NewXmlEntity() *XmlEntity {
	xmlEntity := new(XmlEntity)

	return xmlEntity
}

func (x *XmlData) GetValue() ([]byte, error) {
	entities := columnsToXml(x.Config.Columns[0])
	xmlStr := ""
	for _, entity := range entities {
		xmlStr += entity.ToString()
	}

	return []byte(xmlStr), nil
}

func columnsToXml(column config.Column) []XmlEntity {
	var entities []XmlEntity
	if column.Rows != nil {
		entities = make([]XmlEntity, column.Rows.Size)
	} else {
		entities = make([]XmlEntity, 1)
	}

	for i := 0; i < len(entities); i++ {
		xobj := NewXmlEntity()
		xobj.NodeName = column.Name
		if column.Attrs != nil {
			xobj.Attrs = make(map[string]string)
			for key, value := range column.Attrs {
				xobj.Attrs[key] = generator.Generate(value)
			}
		}

		if len(column.Columns) > 0 {
			for _, childCol := range column.Columns {
				xobj.Children = append(xobj.Children, columnsToXml(childCol)...)
			}
		} else if column.Rows != nil && column.Rows.Format == "" {
			for _, childCol := range column.Rows.Columns {
				xobj.Children = append(xobj.Children, columnsToXml(childCol)...)
			}
		} else if column.Rows != nil && column.Rows.Format != "" {
			xobj.Value = generator.Generate(column.Rows.Format)
		} else {
			xobj.Value = generator.Generate(column.Format)
		}

		entities[i] = *xobj
	}

	return entities
}

func (x *XmlEntity) ToString() string {
	xmlStr := ""
	startTag := fmt.Sprintf("<%s>", x.NodeName)
	endTag := fmt.Sprintf("</%s>", x.NodeName)
	if x.Children != nil {
		childXmlStr := ""
		for _, child := range x.Children {
			childXmlStr += child.ToString()
		}
		xmlStr = fmt.Sprintf("%s%s%s", startTag, childXmlStr, endTag)
	} else if x.Value != nil {
		xmlStr = fmt.Sprintf("%s%s%s", startTag, x.Value, endTag)
	}

	if x.Attrs != nil {
		innerTag := ""
		for key, value := range x.Attrs {
			innerTag += fmt.Sprintf(" %s=\"%s\"", key, value)
		}
		xmlStr = strings.Replace(xmlStr, startTag, fmt.Sprintf("<%s%s>", x.NodeName, innerTag), 1)
	}

	return xmlStr
}
