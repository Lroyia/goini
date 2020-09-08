package test

import (
	"fmt"
	"goini"
	"testing"
)

func Test(t *testing.T) {
	conf, _ := goini.Read("test.ini")
	fmt.Println(conf)
	fmt.Println(conf.GetValueBySection("ina", "mat"))
	fmt.Println(conf.GetValueByItem("mat"))
	fmt.Println(conf.GetAllItemInSection("ina"))
}
