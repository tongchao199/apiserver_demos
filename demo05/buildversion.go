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
			fmt.Println("Rev:", "dffa0b8")
			fmt.Println("BuildTime:", "2018-07-16 12:38:45 CST")
			fmt.Println("CompilerVersion:", "go1.10.1 darwin/amd64")
			os.Exit(0)
		}
	}
}
