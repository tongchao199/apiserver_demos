target := $(notdir $(PWD))

ifndef APPNAME
	APPNAME = ${target}_app
endif

.PHONY: all

all: build

build: generateVer main

.PHONY: main
main:
	#env GOOS=linux GOARCH=amd64 go build -o $(EXENAME) main.go
	go build -o $(APPNAME) main.go

.PHONY: clean
clean:
	rm -rf buildversion.go ${target}*

BRANCH			:= $(shell git branch | sed 's/* \(.*\)/\1/p')
GITREV 			:= $(shell git rev-parse --short HEAD)
BUILDTIME 		:= $(shell date '+%F %T %Z')
COMPILERVERSION	:= $(subst go version ,,$(shell go version))

define GENERATE_VERSION_CODE
cat << EOF | gofmt > buildversion.go
package main

import (
	"fmt"
	"os"
	"strings"
)

func init() {
	if len(os.Args) > 1 {
		if strings.ToLower(os.Args[1]) == "buildversion" {
			fmt.Println("Branch:", "$(BRANCH)")
			fmt.Println("Rev:", "$(GITREV)")
			fmt.Println("BuildTime:", "$(BUILDTIME)")
			fmt.Println("CompilerVersion:", "$(COMPILERVERSION)")
			os.Exit(0)
		}
	}
}
EOF
endef
export GENERATE_VERSION_CODE

generateVer:
	@echo "$$GENERATE_VERSION_CODE" | bash
