package oswrapper

//package main

import (
	"fmt"
	"os"
)

func GetEgid() int {
	return os.Getegid()
}

func main() {
	fmt.Printf("%d\n", GetEgid())
}
