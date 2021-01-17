package rest

import (
	"fmt"
	"github.com/ajangi/gogit/src/consts"
	"io/ioutil"
	"net/http"
)

var mainUrl = "https://api.github.com"

func Fetch(command string, flag Flag) (error,[]byte){
	res, err := http.Get( GetUri(command,flag))
	if err != nil {
		fmt.Printf( consts.ErrorColor,"Error: " + err.Error() )
		return err,nil
	}
	defer res.Body.Close()
	//var items Items
	body, err := ioutil.ReadAll(res.Body)
	//_err := json.Unmarshal(body, &items)
	return err,body
}

func GetUri(command string, flag Flag)string{
	switch command {
	case "events":
		return GetEventsUrl(flag)
	}
	return mainUrl
}

func GetEventsUrl(flag Flag)string{
	var uri = mainUrl
	if flag.User != "" {
		uri = uri + "/users/" + flag.User
	}
	return uri + "/events"
}