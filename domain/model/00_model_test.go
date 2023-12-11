package model

import (
	"testing"

	"github.com/SyaibanAhmadRamadhan/gocatch/ginfra/gdb"
)

func TestGenerator(t *testing.T) {
	gen := []gdb.GeneratorModelForStructParam{
		articleModel(),
		tagModel(),
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

func tagModel() gdb.GeneratorModelForStructParam {
	return gdb.GeneratorModelForStructParam{
		Src: &Tag{},
		SpecifiationTable: gdb.SpecifiationTable{
			TableName: "tag",
		},
		Tag:      "bson",
		FileName: "Tag",
	}
}
