package commands

import (
	"os"
	//"flag"
	//"fmt"
	common "github.com/ajangi/gogit/src/common"
)

func HandleFlags (){
	GetFlags();
}

func GetFlags (){
	args := os.Args[1:];
	if len(args) == 0 {
		common.Help()
	}
	for _, v := range args {
		if v == "help" || v == "-help" || v =="--help"{
			common.Help();
		}
	}
	//helpFlag := flag.String("help","","");
	//flag.Parse();
	//fmt.Println("----------------");
	//fmt.Println(*helpFlag+" asdsadj");
	//fmt.Println("----------------");
	//if *helpFlag == "" {
    //    common.Help();
    //}
}