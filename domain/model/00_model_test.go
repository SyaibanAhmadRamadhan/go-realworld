package model

import (
	"testing"

	"github.com/SyaibanAhmadRamadhan/gocatch/ginfra/gdb"
)

func TestGenerator(t *testing.T) {
	gen := []gdb.GeneratorModelForStructParam{
		articleModel(),
		tagModel(),
		userModel(),
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
		FileName: "Tags",
	}
}

func userModel() gdb.GeneratorModelForStructParam {
	return gdb.GeneratorModelForStructParam{
		Src: &User{},
		SpecifiationTable: gdb.SpecifiationTable{
			TableName: "user",
		},
		Tag:      "bson",
		FileName: "User",
	}
}
