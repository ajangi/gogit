package commands

import (
	"os"
	//"flag"
	//"fmt"
	common "github.com/ajangi/gogit/src/common"
)

func HandleFlags (){
	GetFlags()
}

func GetFlags () bool{
	args := os.Args[1:]
	if len(args) == 0 {
		common.Help()
	}
	for _, v := range args {
		if v == "help" || v == "-help" || v =="--help" || v == "-h"{
			common.Help()
			return true
		}
		if v == "version" || v == "-version" || v =="--version" || v == "-v"{
			common.Version()
			return true
		}
	}
	return true
	//helpFlag := flag.String("help","","");
	//flag.Parse();
	//fmt.Println("----------------");
	//fmt.Println(*helpFlag+" asdsadj");
	//fmt.Println("----------------");
	//if *helpFlag == "" {
    //    common.Help();
    //}
}