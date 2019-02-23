package main

import (
	"flag"
	"fmt"
	"net/url"
	"strconv"

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
		return
	default:
		fmt.Println("mode options are block, unblock, mute, or unmute")
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
		for _, id := range ids {
			fmt.Println("Blocking user: ", id)
			api.BlockUserId(id, v)
		}
	case "unblock":
		fmt.Println("We're unblocking users..")
		for _, id := range ids {
			fmt.Println("Unblocking user: ", id)
			api.UnblockUserId(id, v)
		}
	case "mute":
		fmt.Println("We're muting users..")
		for _, id := range ids {
			fmt.Println("Muting user: ", id)
			api.MuteUserId(id, v)
		}
	case "unmute":
		fmt.Println("We're unmuting users..")
		for _, id := range ids {
			fmt.Println("Unblocking user: ", id)
			api.UnmuteUserId(id, v)
		}
	default:
		fmt.Println("mode options are block, unblock, mute, or unmute")
	}
	fmt.Println("Done.")

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
