package main

import (
	"./cmd"
	"./util"
)

func main() {
	util.PrintBanner()
  cmd.Execute()
}
