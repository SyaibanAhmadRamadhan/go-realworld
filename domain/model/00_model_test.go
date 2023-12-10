package model

import (
	"testing"

	"github.com/SyaibanAhmadRamadhan/gocatch/ginfra/gdb"
)

func TestGenerator(t *testing.T) {
	gen := []gdb.GeneratorModelForStructParam{
		articleModel(),
	}

	gdb.GeneratorModelFromStruct(gen...)
}

func articleModel() gdb.GeneratorModelForStructParam {
	return gdb.GeneratorModelForStructParam{
		Src: &Article{},
		SpecifiationTable: gdb.SpecifiationTable{
			TableName: "article",
		},
		Tag:      "bson",
		FileName: "Article",
	}
}
