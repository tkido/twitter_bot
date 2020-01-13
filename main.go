package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"regexp"
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
	// fmt.Println("twitter_bot_test")
	// fmt.Println(month)
	// fmt.Println(week)
	// fmt.Println(date)
	// fmt.Println(week % 13)
	// fmt.Println(week % 12)

	task("honki", month)
	task("franklin", week%13)
	task("survival", week%12)
}

type Account struct {
	AccessToken       string `json:"accessToken"`
	AccessTokenSecret string `json:"accessTokenSecret"`
	ConsumerKey       string `json:"consumerKey"`
	ConsumerSecret    string `json:"consumerSecret"`
}

func tweet(text string) error {
	tweet, err := api.PostTweet(text, nil)
	if err != nil {
		return err
	}
	fmt.Println(tweet.Text)
	return nil
}

func testTweet(text string) error {
	fmt.Println(text)
	return nil
}

var re = regexp.MustCompile(`\\n`)

func task(name string, index int) error {
	path := "config/text/" + name + ".txt"
	lines, err := readlines(path)
	if err != nil {
		return err
	}
	line := lines[index]
	line = re.ReplaceAllString(line, "\r\n")
	err = tweet(line)
	return err
}

func readlines(path string) (ss []string, err error) {
	f, err := os.Open(path)
	if err != nil {
		return
	}
	defer f.Close()
	s := bufio.NewScanner(f)
	for s.Scan() {
		ss = append(ss, s.Text())
	}
	if err = s.Err(); err != nil {
		return
	}
	return
}
