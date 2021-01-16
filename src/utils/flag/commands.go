package commands

import (
	"fmt"
	"github.com/ajangi/gogit/src/common"
	"github.com/ajangi/gogit/src/consts"
	"github.com/ajangi/gogit/src/utils/rest"
	"os"
	"strconv"
)

func HandleFlags (){
	GetFlags()
}

func GetFlags (){
	var Flag = rest.Flag{}
	args := os.Args[1:]
	if len(args) == 0 {
		common.Help()
	}
	for k, v := range args {
		if v == "help" || v == "-help" || v =="--help" || v == "-h"{
			common.Help()
			return
		}

		if v == "version" || v == "-version" || v =="--version" || v == "-v"{
			common.Version()
			return
		}

		if v == "-user" || v =="--user" || v == "-u"{
			if len(args) < 2 {
				fmt.Printf(consts.ErrorColor,"Error : username not entered after using \"-u\" or \"-user\" or \"--user\" \n")
				return
			}
			Flag.User = args[k+1]
		}

		if v == "-c" || v =="--count" || v == "-count"{
			if len(args) < 2 {
				fmt.Printf(consts.ErrorColor,"Error : count not entered after using \"-c\" or \"-count\" or \"--count\" \n")
				return
			}
			Flag.PerPage, _ = strconv.Atoi(args[k+1])
		}
	}
	
	return
	//helpFlag := flag.String("help","","");
	//flag.Parse();
	//fmt.Println("----------------");
	//fmt.Println(*helpFlag+" asdsadj");
	//fmt.Println("----------------");
	//if *helpFlag == "" {
    //    common.Help();
    //}
}