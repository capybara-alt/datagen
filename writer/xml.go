package writer

import (
	"bufio"

	"github.com/capybara-alt/datagen/config"
	"github.com/capybara-alt/datagen/data"
)

type XmlWriter struct {
	*AWriter
}

func NewXmlWriter(conf config.Config) *XmlWriter {
	w := new(XmlWriter)
	w.AWriter = new(AWriter)
	w.AWriter.Config = conf

	return w
}

func (w *XmlWriter) Write(idata data.IData) error {
	for i := 0; i < *w.Config.Size; i++ {
		b, err := idata.GetValue()
		if err != nil {
			return err
		}

		suffixNum := i + 1
		file, err := w.AWriter.createFile(&suffixNum)
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
