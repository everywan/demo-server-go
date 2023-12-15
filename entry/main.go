package main

import (
	"github.com/everywan/demo-server-go/commons/utils"
	"github.com/everywan/demo-server-go/internal/cmd"
)

func main() {
	utils.PrintBuildInfo()
	cmd.Execute()
}
