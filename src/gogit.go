package main

import (
	"github.com/ajangi/gogit/src/common"
	"github.com/ajangi/gogit/src/utils/flag"
)

func main(){
	common.Intro()
	commands.HandleFlags()
}