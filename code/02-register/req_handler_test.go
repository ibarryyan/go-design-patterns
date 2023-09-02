package _2_register

import (
	"fmt"
	"testing"
)

func TestRegister(t *testing.T) {
	RegisterHandler("/index", func() {
		fmt.Println("index...")
	})
	RegisterHandler("/add", func() {
		fmt.Println("add...")
	})

	RunServer()
}
