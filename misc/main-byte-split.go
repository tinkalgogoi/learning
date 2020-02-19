package main

import (
	"bytes"
	"fmt"
)

func main(){
	myByte := []byte("172					/infoblox")
	fmt.Println(myByte)
	splits := bytes.SplitAfter(myByte, []byte("\t"))
	fmt.Println(string(splits[0]))
}
