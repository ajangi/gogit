package common

import (
	"fmt"
	consts "github.com/ajangi/gogit/src/consts"
)
const helpText = `
usage: gogit [--version] [--help] <command> [-u <user>] [-c <count>]
These are common gogit commands used for various sources :

to get latest events use : 
	gogit evenets -c 10                get latest public events
	gogit events -u ajangi -c 10       get latest 10 public events for user "ajangi"
`+"\n"
func Help(){
	fmt.Printf(consts.InfoColor,helpText)
}

func Version(){
	fmt.Printf(consts.InfoColor,"gogot version "+consts.VERSION+"\n")
}