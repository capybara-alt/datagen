package data

import (
	"encoding/json"
	"sort"

	"github.com/capybara-alt/datagen/config"
	"github.com/capybara-alt/datagen/generator"
)

type FlattenRow struct {
	Format  string
	Depth   int
	Column  int
	RowSize int
	Value   string
}

type CsvData struct {
	config.Config
	WithHeader bool
}

func NewCsvData(conf *config.Config) *CsvData {
	csvData := new(CsvData)
	csvData.Config = *conf
	csvData.WithHeader = conf.WithHeader

	return csvData
}

func (c *CsvData) GetValue() ([]byte, error) {
	flattenRow := nestedConfigToFlattenRow(c.Config.Columns, []FlattenRow{}, 1, 0)
	values := c.convertTo2dArr(flattenRow)
	b, err := json.Marshal(values)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func GetColumnHeader(columns []config.Column, headers []string) []string {
	for _, column := range columns {
		if len(column.Columns) > 0 {
			headers = append(headers, GetColumnHeader(column.Columns, []string{})...)
		} else if column.Rows != nil {
			headers = append(headers, GetColumnHeader(column.Rows.Columns, []string{})...)
		} else {
			headers = append(headers, column.Name)
		}
	}

	return headers
}

func (c *CsvData) convertTo2dArr(flattenRow []FlattenRow) [][]string {
	sort.Slice(flattenRow[:], func(i, j int) bool {
		return flattenRow[i].Depth > flattenRow[j].Depth
	})

	flattenRows := make([][]FlattenRow, getRowSize(c.Config.Columns, 1))
	for i := 0; i < len(flattenRows); i++ {
		flattenRows[i] = flattenRow
	}

	for index, column := range flattenRow {
		for i := 0; i < column.RowSize; i++ {
			value := generator.Generate(column.Format)
			size := len(flattenRows) / column.RowSize
			for j := 0; j < size; j++ {
				flattenRows[i+(j*column.RowSize)][index].Value = value
			}
		}
	}

	values := make([][]string, len(flattenRows))
	for rowIndex, row := range flattenRows {
		columns := make([]string, len(flattenRow))
		sort.SliceStable(row[:], func(i, j int) bool {
			return row[i].Column < row[j].Column
		})
		for index, value := range row {
			columns[index] = value.Value
		}
		values[rowIndex] = columns
	}

	return values
}

func nestedConfigToFlattenRow(columns []config.Column, flattenRow []FlattenRow, rowSize, depth int) []FlattenRow {
	depth++
	for _, column := range columns {
		if column.Rows != nil {
			flattenRow = append(flattenRow, nestedConfigToFlattenRow(column.Rows.Columns, []FlattenRow{}, column.Rows.Size, depth)...)
		} else if len(column.Columns) > 0 {
			flattenRow = append(flattenRow, nestedConfigToFlattenRow(column.Columns, []FlattenRow{}, rowSize, depth)...)
		} else {
			flattenRow = append(flattenRow, FlattenRow{
				Format:  column.Format,
				Depth:   depth,
				RowSize: rowSize,
			})
		}
	}

	for i := 0; i < len(flattenRow); i++ {
		flattenRow[i].Column = i
	}

	return flattenRow
}

func getRowSize(columns []config.Column, rowSize int) int {
	for _, column := range columns {
		if len(column.Columns) > 0 {
			rowSize = getRowSize(column.Columns, rowSize)
		} else if column.Rows != nil {
			rowSize *= column.Rows.Size
			rowSize = getRowSize(column.Rows.Columns, rowSize)
		}
	}

	return rowSize
}
