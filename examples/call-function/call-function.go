package main

import (
	"fmt"
	"../.."
)

type Data struct {
	A, B int
}

func main() {
	l := luna.New(luna.AllLibs)
	l.LoadFile("call-function.lua")

	_, err := l.Call("noparams")
	if err != nil {
		fmt.Println("Error calling 'noparams':", err)
	}
	_, err = l.Call("basicTypes", 3, 4.2, "hello", false, nil)
	if err != nil {
		fmt.Println("Error calling 'basicTypes':", err)
	}
	_, err = l.Call("struct", Data{3, 2})
	if err != nil {
		fmt.Println("Error calling 'struct':", err)
	}
	ret, err := l.Call("ret")
	if err != nil {
		fmt.Println("Error calling 'ret':", err)
	} else {
		fmt.Println("Return values:", ret)
	}
}