package _2_register

import "fmt"

var reqMap = make(map[string]func())

func RegisterHandler(path string, h func()) {
	if _, ok := reqMap[path]; ok {
		panic("path aleary exited")
	}
	reqMap[path] = h
}

func RunServer() {
	for k, v := range reqMap {
		fmt.Println("Run " + k)
		v()
	}
}
