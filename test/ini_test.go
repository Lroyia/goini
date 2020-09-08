package test

import (
	"fmt"
	"goini"
	"testing"
)

func Test(t *testing.T) {
	conf, _ := goini.Read("test.ini")
	fmt.Println(conf)
}
