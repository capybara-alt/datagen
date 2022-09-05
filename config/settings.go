package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/capybara-alt/datagen/utils"
)

type Rows struct {
	Format  string   `json:"format"`
	Size    int      `json:"size"`
	Columns []Column `json:"columns"`
}

type Column struct {
	Name    string            `json:"name"`
	Format  string            `json:"format"`
	Attrs   map[string]string `json:"attrs"`
	Columns []Column          `json:"columns"`
	Rows    *Rows             `json:"rows"`
}

type Config struct {
	Name       string
	Format     string   `json:"format"`
	WithHeader bool     `json:"withHeader"`
	Size       *int     `json:"size"`
	Columns    []Column `json:"columns"`
	Rows       *Rows    `json:"rows"`
}

var appConfigs map[string]*Config

func Init() error {
	files, err := filepath.Glob("./config/settings/*.json")
	if err != nil {
		return err
	}

	appConfigs = make(map[string]*Config, len(files))
	for _, file := range files {
		f, err := os.Open(file)
		if err != nil {
			return err
		}
		defer f.Close()

		data, err := ioutil.ReadAll(f)
		if err != nil {
			return err
		}

		var appConfig *Config
		if err = json.Unmarshal(data, &appConfig); err != nil {
			return err
		}

		if appConfig.Size == nil {
			return errors.New("Column size cannot be empty.")
		}

		if !utils.ContainsString(OutputFormats, appConfig.Format) {
			return fmt.Errorf("Config format must be: %s", strings.Join(OutputFormats, ", "))
		}

		appConfig.Name = strings.ReplaceAll(filepath.Base(file), ".json", "")
		appConfigs[appConfig.Name] = appConfig
	}

	return nil
}

func GetAppConfg(key string) *Config {
	return appConfigs[key]
}

func GetAppConfgs() map[string]*Config {
	return appConfigs
}
