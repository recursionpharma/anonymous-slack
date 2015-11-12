package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"
)

const (
	tokenConfig   = "INCOMING_SLACK_TOKEN"
	webhookConfig = "INCOMING_SLACK_WEBHOOK"
	// Incoming payload form will have the following keys:
	// (See: https://api.slack.com/slash-commands)
	keyToken       = "token"
	keyTeamID      = "team_id"
	keyChannelID   = "channel_id"
	keyChannelName = "channel_name"
	keyUserID      = "user_id"
	keyUserName    = "user_name"
	keyCommand     = "command"
	keyText        = "text"
)

type slackMsg struct {
	Text      string `json:"text"`
	Username  string `json:"username"` // Anonymous animal sender
	Channel   string `json:"channel"`  // Recipient
	AsUser    string `json:"as_user"`
	IconURL   string `json:"icon_url"`
	LinkNames string `json:"link_names"`
}

var (
	port       int
	token      string
	channel    string
	avatars    = []string{}
	avatarIcon = []string{}
	avatarText = []string{}
)

// readAnonymousMessage parses the username and re-routes
// the message to the user from an anonymous avatar
func readAnonymousMessage(r *http.Request) string {
	err := r.ParseForm()
	// TODO: Change HTTP status code
	if err != nil {
		return string(err.Error())
	}
	if len(r.Form[keyToken]) == 0 || r.Form[keyToken][0] != token {
		return "Config error."
	}
	if len(r.Form[keyText]) == 0 {
		return "Slack bug; inform the team."
	}
	msg := strings.TrimSpace(r.Form[keyText][0])

	err = sendAnonymousMessage(msg)
	if err != nil {
		return "Failed to send message."
	}
	return fmt.Sprintf("Anonymously sent [%s] ", msg)
}

// sendAnonymousMessage uses an incoming hook to Direct Message
// the given user the message, from a random animal.
func sendAnonymousMessage(message string) error {
	url := os.Getenv(webhookConfig)
	avatarID := rand.Intn(len(avatars))
	payload, err := json.Marshal(slackMsg{
		Text:      message,
		Channel:   channel,
		AsUser:    "False",
		IconURL:   avatarIcon[avatarID],
		LinkNames: "1",
		Username:  fmt.Sprintf("%s : %s", avatars[avatarID], avatarText[avatarID]),
	})
	if err != nil {
		return err
	}
	_, err = http.Post(url, "application/json", bytes.NewBuffer(payload))
	return err
}

func main() {
	rand.Seed(time.Now().UnixNano())
	channel = os.Getenv("SLACK_CHANNEL_ID")
	token = os.Getenv(tokenConfig)
	getAvatars()
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		result := readAnonymousMessage(r)
		fmt.Fprintf(w, result)
	})
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}

func init() {
	flag.IntVar(&port, "port", 5000, "HTTP server port")
	flag.Parse()
}

// Avatar data structure
type Avatar struct {
	Username    string `json:"username"`
	DefaultText string `json:"default_text"`
	IconURL     string `json:"icon_url"`
}

func getAvatars() []Avatar {
	raw, err := ioutil.ReadFile("./avatars.json")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	var c []Avatar
	json.Unmarshal(raw, &c)
	for _, p := range c {
		avatars = append(avatars, p.Username)
		avatarIcon = append(avatarIcon, p.IconURL)
		avatarText = append(avatarText, p.DefaultText)

	}
	return c
}
