package main

import (
	common "github.com/ajangi/gogit/src/common"
	commands "github.com/ajangi/gogit/src/utils/flag"
)

func main(){
	common.Intro();
	commands.HandleFlags();
	/* helpFlag := flag.String("help","","The help flag.")
	flag.Parse()
	if *helpFlag == "" {
        flag.PrintDefaults()
        os.Exit(1)
    }
	fmt.Printf("help: %s \n", *helpFlag, ) */
}