package commands

import (
	"encoding/json"
	"fmt"
	"github.com/ajangi/gogit/src/common"
	"github.com/ajangi/gogit/src/consts"
	"github.com/ajangi/gogit/src/utils/rest"
	"github.com/olekukonko/tablewriter"
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
	var commandsList = args
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
			commandsList = append(commandsList[:k], commandsList[k+2:]...)
		}

		if v == "-c" || v =="--count" || v == "-count"{
			if len(args) < 2 {
				fmt.Printf(consts.ErrorColor,"Error : count not entered after using \"-c\" or \"-count\" or \"--count\" \n")
				return
			}
			Flag.PerPage, _ = strconv.Atoi(args[k+1])
			commandsList = append(commandsList[:k], commandsList[k+2:]...)
		}
	}
	var verifiedCommands = commandsList
	var j = 0
	for _,command := range commandsList {
		for _, value := range consts.GetCommandsList() {
			if value == command {
				verifiedCommands[j] = command
				j++
				break
			}
		}
	}
	verifiedCommands = verifiedCommands[:j]
	if len(verifiedCommands) == 0 {
		fmt.Printf(consts.ErrorColor,"Error : no valid commands found! \n")
		return
	}
	for _,v := range verifiedCommands {
		err,response := rest.Fetch(v, Flag)
		if err != nil  {
			fmt.Printf(consts.ErrorColor, "Error : " + err.Error() +"\n")
		}
		var events rest.Events
		err = json.Unmarshal(response, &events)
		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"#","ID", "Type", "Username", "Repo", "RepoUrl"})
		for index,event := range events{
			table.Append([]string{strconv.Itoa(index),event.ID,event.Type,event.Actor.Username,event.Repo.Name, event.Repo.Url})
		}
		table.Render()
	}
	return
}