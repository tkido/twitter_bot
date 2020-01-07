package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"time"

	"github.com/ChimeraCoder/anaconda"
)

var (
	api   *anaconda.TwitterApi
	month int
	week  int
	date  int
)

func init() {
	// prepare twitter api
	raw, err := ioutil.ReadFile("config/account/account.json")
	if err != nil {
		panic(err)
	}
	var acc Account
	json.Unmarshal(raw, &acc)
	api = anaconda.NewTwitterApiWithCredentials(acc.AccessToken, acc.AccessTokenSecret, acc.ConsumerKey, acc.ConsumerSecret)
	// prepare DateTime info
	now := time.Now()
	origin := time.Date(1900, 1, 1, 0, 0, 0, 0, time.Local)
	duration := now.Sub(origin)
	date = int(math.Ceil(duration.Hours()/24) + 1)
	week = date / 7
	month = int(now.Month()) - 1
}

func main() {
	fmt.Println("twitter_bot_test")
	fmt.Println(month)
	fmt.Println(week)
	fmt.Println(date)

	fmt.Println(week % 13)
	fmt.Println(week % 12)
}

type Account struct {
	AccessToken       string `json:"accessToken"`
	AccessTokenSecret string `json:"accessTokenSecret"`
	ConsumerKey       string `json:"consumerKey"`
	ConsumerSecret    string `json:"consumerSecret"`
}

func tweet(text string) error {
	tweet, err := api.PostTweet("botのテスト", nil)
	if err != nil {
		return err
	}
	fmt.Println(tweet.Text)
	return nil
}
