package main

import (
	"bufio"
	"flag"
	"log"
	"os"

	helper "github.com/clsung/line-oauth2-helper"
)

func main() {
	var (
		filePath  = flag.String("file", "", "file path to privatekey.json")
		channelID = flag.String("channel_id", os.Getenv("CHANNEL_ID"), "Channel ID on channel console")
	)
	flag.Parse()
	h := helper.New(*channelID)
	var jwt string
	if *filePath == "" {
		jwt, _ = h.GetLineJWT(bufio.NewReader(os.Stdin))
	} else {
		jwt, _ = h.GetLineJWTFromFile(*filePath)
	}
	log.Printf("JSON Web Token for LINE OAuth2 v2.1: %s", jwt)
}
