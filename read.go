package main

import (
	"flag"
	"fmt"
	"net/url"
	"strconv"
	"time"

	"github.com/ChimeraCoder/anaconda"
	"github.com/kayleethemech/twuser/fileutil"
	"github.com/spf13/viper"
)

func main() {
	appConfig := TwuserConfig{}
	appConfig.readConfig()
	// --mode can be block/unblock/mute/unmute
	var mode string
	flag.StringVar(&mode, "mode", "notgiven", "a string")
	// Only one of the following can be chosen
	// --id can be given manually
	idPtr := flag.String("id", "", "enter an id to process")
	// --file as a file
	filePtr := flag.String("file", "", "a file containing ids")
	// --dir as a directory with files
	dirPtr := flag.String("dir", "", "a directory containing files with ids")
	flag.Parse()
	//--------------------------
	// sanitizing the input

	// mode

	switch mode {
	case "block", "unblock", "mute", "unmute":
		//continue
	case "":
		fmt.Println("Godess, instruct us with a mode setting please.")
	default:
		fmt.Println("mode options are block, unblock, mute, or unmute, was", mode)
		return
	}
	// Source of Ids
	var ids []int64
	if *idPtr != "" {
		// single id has been specified
		intValue, _ := strconv.ParseInt(*idPtr, 10, 64)
		ids = append(ids, intValue)
	} else if *filePtr != "" {
		// file was specificied
		ids = fileutil.ReadTwitterIds(*filePtr)
	} else if *dirPtr != "" {
		// directory was specified
		ids = fileutil.ReadTwitterIdsFromFilesInPath(*dirPtr)
	} else {
		fmt.Println("A source of twitter ids has to be given")
		fmt.Println("Possibilities:")
		fmt.Println("--dir PATH: where PATH directs to a folder containing files with twitter ids")
		fmt.Println("--file FILE: where FILE directs to a file containing twitter ids")
		fmt.Println("--id ID: where ID is a twitter id")
		return
	}

	//--------------------------//
	api := anaconda.NewTwitterApiWithCredentials(appConfig.accessToken, appConfig.accessSecret, appConfig.appApiKey, appConfig.appApiSecret)

	fmt.Println("Welcome Misses, we're humbly at your service.")
	v := url.Values{}

	switch mode {
	case "block":
		fmt.Println("We're blocking for you..")
		doIterate("Blocking user:", ids, api.BlockUserId, v)
	case "unblock":
		fmt.Println("We're unblocking users..")
		doIterate("Unblocking user:", ids, api.UnblockUserId, v)
	case "mute":
		fmt.Println("We're muting users..")
		doIterate("Muting user:", ids, api.MuteUserId, v)
	case "unmute":
		fmt.Println("We're unmuting users..")
		doIterate("Unblocking user:", ids, api.UnmuteUserId, v)
	default:
		fmt.Println("mode options are block, unblock, mute, or unmute")
	}
	fmt.Println("Done.")

}

type apiAction func(int64, url.Values) (anaconda.User, error)

func doIterate(output string, ids []int64, action apiAction, v url.Values) {
	for _, id := range ids {
		fmt.Println(output, id)
		action(id, v)
		time.Sleep(time.Second)
	}
}

type TwuserConfig struct {
	appApiKey    string
	appApiSecret string
	accessToken  string
	accessSecret string
}

func (config *TwuserConfig) readConfig() {
	viper.SetConfigType("yaml")
	viper.SetConfigName("twitterapi")
	viper.AddConfigPath("$HOME/.twuser/")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Error reading config: %s \n", err))
	}
	config.appApiKey = viper.GetString("ApiKey")
	config.appApiSecret = viper.GetString("ApiSecret")
	config.accessToken = viper.GetString("Token")
	config.accessSecret = viper.GetString("Secret")
}
