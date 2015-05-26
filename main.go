package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
)

var (
	token    = flag.String("token", "", "token")
	username = flag.String("username", "bot", "username displayed")
	args     []string
)

const (
	SLACK = "https://slack.com"
)

func get(path string, qs map[string]string, f func(map[string]interface{}) string) string {
	Url, _ := url.Parse(SLACK)
	Url.Path += "/api"
	Url.Path += path

	params := url.Values{}
	params.Add("token", *token)
	for k, v := range qs {
		params.Add(k, v)
	}
	Url.RawQuery = params.Encode()

	resp, err := http.Get(Url.String())
	if err != nil {
		os.Exit(1)
	}
	body, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	var res map[string]interface{}
	if err := json.Unmarshal(body, &res); err != nil {
		os.Exit(1)
	}

	if f != nil {
		return f(res)
	}
	return ""
}

func handle_auth_test(j map[string]interface{}) string {
	v := j["user_id"]
	return v.(string)
}

func handle_im_open(j map[string]interface{}) string {
	channel := j["channel"]
	id := channel.(map[string]interface{})["id"]
	return id.(string)
}

func main() {
	flag.Parse()
	args = flag.Args()
	msg := strings.Join(args, " ")
	// 1 get user id
	user_id := get("/auth.test", map[string]string{}, handle_auth_test)
	// 2 open channel
	channel_id := get("/im.open", map[string]string{"user": user_id}, handle_im_open)
	// 3 send message
	get("/chat.postMessage", map[string]string{"channel": channel_id, "text": msg, "username": *username}, nil)

}
