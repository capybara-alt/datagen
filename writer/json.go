package writer

import (
	"bufio"

	"github.com/capybara-alt/datagen/config"
	"github.com/capybara-alt/datagen/data"
)

type JsonWriter struct {
	*AWriter
}

func NewJsonWriter(conf config.Config) *JsonWriter {
	w := new(JsonWriter)
	w.AWriter = new(AWriter)
	w.AWriter.Config = conf

	return w
}

func (w *JsonWriter) Write(idata data.IData) error {
	for i := 0; i < *w.Config.Size; i++ {
		b, err := idata.GetValue()
		if err != nil {
			return err
		}

		suffixNumber := i + 1
		file, err := w.AWriter.createFile(&suffixNumber)
		if err != nil {
			return err
		}
		defer file.Close()

		writer := bufio.NewWriter(file)
		if _, err := writer.Write(b); err != nil {
			return err
		}
		writer.Flush()
	}

	return nil
}
