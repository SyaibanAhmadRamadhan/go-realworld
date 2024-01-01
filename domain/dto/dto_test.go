package dto

import (
	"fmt"
	"testing"

	"github.com/SyaibanAhmadRamadhan/gocatch/gcommon"
	"github.com/SyaibanAhmadRamadhan/gocatch/gvalidation"
)

func TestValidation(t *testing.T) {
	validate := gvalidation.New()
	validate.SetIdTranslation()
	req := new(RequestCreateArticle)

	err := validate.StructM(req)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(gcommon.NewUlid())
}
