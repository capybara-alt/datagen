package writer

import (
	"bufio"
	"encoding/csv"
	"encoding/json"

	"github.com/capybara-alt/datagen/config"
	"github.com/capybara-alt/datagen/data"
)

type CsvWriter struct {
	*AWriter
}

func NewCsvWriter(conf config.Config) *CsvWriter {
	w := new(CsvWriter)
	w.AWriter = new(AWriter)
	w.AWriter.Config = conf

	return w
}

func (w *CsvWriter) Write(idata data.IData) error {
	rows := [][]string{}
	if w.Config.WithHeader {
		rows = append(rows, data.GetColumnHeader(w.Config.Columns, []string{}))
	}
	for i := 0; i < *w.Config.Size; i++ {
		b, err := idata.GetValue()
		if err != nil {
			return err
		}
		records := [][]string{}
		json.Unmarshal(b, &records)
		rows = append(rows, records...)
	}
	file, err := w.createFile(nil)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(bufio.NewWriter(file))
	if err := writer.WriteAll(rows); err != nil {
		return err
	}
	writer.Flush()

	return nil
}
