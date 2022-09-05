// テストデータ書き込み用パッケージ
// ファイルフォーマットを拡張する際は、IWriterを継承して、Writeを実装する
package writer

import (
	"fmt"
	"os"

	"github.com/capybara-alt/datagen/config"
	"github.com/capybara-alt/datagen/data"
)

const OUTPUT_DIR = "./out"

// 書き込み用インターフェース
type IWriter interface {
	Write(idata data.IData) error
}

type AWriter struct {
	config.Config
	IWriter
}

func (a *AWriter) createFile(suffixNumber *int) (*os.File, error) {
	filename := ""
	if suffixNumber == nil {
		filename = fmt.Sprintf("%s/%s/%s.%s", OUTPUT_DIR, a.Config.Format, a.Config.Name, a.Config.Format)
	} else {
		filename = fmt.Sprintf("%s/%s/%s_%d.%s", OUTPUT_DIR, a.Config.Format, a.Config.Name, *suffixNumber, a.Config.Format)
	}
	file, err := os.Create(filename)
	if err != nil && !os.IsExist(err) {
		return nil, err
	}

	return file, nil
}
