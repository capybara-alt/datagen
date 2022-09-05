package generator

import (
	"log"

	"github.com/zach-klippenstein/goregen"
)

func Generate(format string) string {
	result, err := regen.Generate(format)
	if err != nil {
		log.Println(err)
		return ""
	}

	return result
}

