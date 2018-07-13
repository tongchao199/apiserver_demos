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
			fmt.Println("Rev:", "5da7ed2")
			fmt.Println("BuildTime:", "2018-07-13 11:20:19 CST")
			fmt.Println("CompilerVersion:", "go1.10.1 darwin/amd64")
			os.Exit(0)
		}
	}
}
