package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/ChimeraCoder/anaconda"
)

func main() {
	fmt.Println("twitter_bot_test")
}

type Account struct {
	AccessToken       string `json:"accessToken"`
	AccessTokenSecret string `json:"accessTokenSecret"`
	ConsumerKey       string `json:"consumerKey"`
	ConsumerSecret    string `json:"consumerSecret"`
}

func tweet(text string) error {
	raw, err := ioutil.ReadFile("config/account/account.json")
	if err != nil {
		return err
	}
	var acc Account
	json.Unmarshal(raw, &acc)
	api := anaconda.NewTwitterApiWithCredentials(acc.AccessToken, acc.AccessTokenSecret, acc.ConsumerKey, acc.ConsumerSecret)
	tweet, err := api.PostTweet("botのテスト", nil)
	if err != nil {
		return err
	}
	fmt.Println(tweet.Text)
	return nil
}
