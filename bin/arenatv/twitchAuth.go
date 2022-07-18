package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type twitchAuth struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
	TokenType   string `json:"token_type"`
}

func getTwitchToken() <-chan twitchAuth {
	url := "https://id.twitch.tv/oauth2/token"
	method := "POST"
	payload := strings.NewReader("client_id=3f7ouqiwbe2x7mpqsl3qvipx4kjkxm&client_secret=3uc4xzm9vncape8va8eq7wkwcmddug&grant_type=client_credentials")
	client := &http.Client{}
	println("Called getTwitchToken...")
	tt := make(chan twitchAuth)
	go func() {
		req, err := http.NewRequest(method, url, payload)
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		if err != nil {
			app.Log(err)
			return
		}
		res, err := client.Do(req)
		if err != nil {
			fmt.Println(err)
			return
		}

		defer res.Body.Close()

		b, err := ioutil.ReadAll(res.Body)
		if err != nil {
			app.Log(err)
			return
		}
		var twitchToken twitchAuth
		err = json.Unmarshal(b, &twitchToken)
		if err != nil {
			panic(err)
		}
		tt <- twitchAuth(twitchToken)
	}()
	return tt
}
