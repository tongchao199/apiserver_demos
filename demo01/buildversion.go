package main

import (
	"fmt"
	"os"
	"strings"
)

func init() {
	if len(os.Args) > 1 {
		if strings.ToLower(os.Args[1]) == "buildversion" {
			fmt.Println("Branch:", "master master")
			fmt.Println("Rev:", "aff9ed1")
			fmt.Println("BuildTime:", "2018-07-12 11:33:22 CST")
			fmt.Println("CompilerVersion:", "go1.10.1 darwin/amd64")
			os.Exit(0)
		}
	}
}
