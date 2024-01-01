package model

import (
	"fmt"
	"testing"

	"github.com/SyaibanAhmadRamadhan/gocatch/ginfra/gdb"
)

func TestGenerator(t *testing.T) {
	gen := []gdb.GeneratorModelForStructParam{
		articleModel(),
		tagModel(),
		userModel(),
		commentModel(),
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

func commentModel() gdb.GeneratorModelForStructParam {
	return gdb.GeneratorModelForStructParam{
		Src: &Comment{},
		SpecifiationTable: gdb.SpecifiationTable{
			TableName: "comment",
		},
		Tag:      "bson",
		FileName: "Comment",
	}
}

func TestName(t *testing.T) {
	user := NewUser()

	printUser(user, []string{
		user.SetId("asd"),
	})
}

func printUser(user *User, columns []string) {
	fmt.Println(user)
	fmt.Println(user.GetValuesByColums(columns...))
}
