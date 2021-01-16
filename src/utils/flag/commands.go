package commands

import (
	"os"
	//"fmt"
	common "github.com/ajangi/gogit/src/common"
)

func HandleFlags (){
	GetFlags();
}

func GetFlags (){
	args := os.Args;
	
	if len(args) == 1 {
		common.Help()
	}
}