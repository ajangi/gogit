package main

import (
	common "github.com/ajangi/gogit/src/common"
	commands "github.com/ajangi/gogit/src/utils/flag"
)

func main(){
	common.Intro()
	commands.HandleFlags()
}