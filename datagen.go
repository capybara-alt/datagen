package main

import (
	"log"

	"github.com/capybara-alt/datagen/config"
	"github.com/capybara-alt/datagen/data"
	"github.com/capybara-alt/datagen/writer"
)

func main() {
	if err := config.Init(); err != nil {
		log.Fatal(err)
	}

	appConfigs := config.GetAppConfgs()
	for _, value := range appConfigs {
		var w writer.IWriter
		var idata data.IData
		switch value.Format {
			case "json":
				w = writer.NewJsonWriter(*value)
				idata = data.NewJsonData(value)
			case "csv":
				w = writer.NewCsvWriter(*value)
				idata = data.NewCsvData(value)
			case "xml":
				w = writer.NewXmlWriter(*value)
				idata = data.NewXmlData(value)
		}
		if err := w.Write(idata); err != nil {
			log.Fatal(err)
		}
	}
}
