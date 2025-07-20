package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

type messageParams struct {
	apiVersion int64
	token      string
	channelID  string
	content    string
}

func cliMessage() {
	env := loadEnv()
	// get user input from terminal for message + channel id
	channelID := getInput("channel ID:")
	for {
		content := getInput("message content:")
		msg := messageParams{
			apiVersion: 10,
			token:      env["DISCORD_TOKEN"],
			channelID:  channelID,
			content:    content,
		}
		// fmt.Printf("msg: %s \n", msg.content)
		sendMessage(msg)
	}
}

func getInput(prompt string) string {
	s := bufio.NewScanner(os.Stdin)
	var buf string

	fmt.Println(prompt)
	s.Scan()
	buf = s.Text()
	var result string
	if len(buf) != 0 {
		result = buf
	} else {
		result = "fail"
	}
	return result
}

func sendMessage(msg messageParams) {
	// type payload struct {
	// 	Content string `json:"content"`
	// }
	targetUrl := fmt.Sprintf("https://discordapp.com/api/channels/%s/messages", msg.channelID)

	payload, err := json.Marshal(
		struct {
			Content string `json:"content"`
		}{msg.content})

	req, error := http.NewRequest("POST", targetUrl, strings.NewReader(string(payload)))
	if error != nil {
		log.Fatal("request failed lol")
	}
	authHeader := fmt.Sprintf("Bot %s", msg.token)
	req.Header.Add("Authorization", authHeader)
	req.Header.Add("User-Agent", userAgent)
	req.Header.Add("Content-Type", contentType)

	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("discord response: %s", body)
}
