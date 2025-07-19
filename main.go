package main

import (
	"bufio"
	"io"
	// "encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	// "strconv"
	"strings"
)

type messageParams struct {
	apiVersion int64
	token      string
	channelID  string
	content    string
}

const contentType string = "application/json"
const userAgent string = "pacbot https://github.com/violetcircus/pacbot"

func main() {
	envs := loadEnv()
	// get user input from terminal for message + channel id
	channelID := getInput("channel ID:")
	for {
		content := getInput("message content:")
		msg := messageParams{
			apiVersion: 10,
			token:      envs["DISCORD_TOKEN"],
			channelID:  channelID,
			content:    content,
		}
		// fmt.Printf("msg: %s \n", msg.content)
		sendMessage(msg)
	}
}

// load envs. doing it this way is dumb: use normal file reading and just string manip the lines into a struct lol
func loadEnv() map[string]string {
	f, err := os.Open(".env")
	if err != nil {
		log.Fatal("error reading envs file", err)
	}

	s := bufio.NewScanner(f)
	s.Split(bufio.ScanLines)

	envs := make(map[string]string)
	for s.Scan() {
		a := strings.Split(s.Text(), "=")
		envs[a[0]] = a[1]
	}

	return envs
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
	targetUrl := fmt.Sprintf("https://discordapp.com/api/channels/%s/messages", msg.channelID)

	payload := fmt.Sprintf("{ \"content\": \"%s\" }", msg.content)
	req, error := http.NewRequest("POST", targetUrl, strings.NewReader(payload))
	if error != nil {
		log.Fatal("request failed lol")
	}
	authHeader := fmt.Sprintf("Bot %s", msg.token)
	req.Header.Add("Authorization", authHeader)
	req.Header.Add("User-Agent", userAgent)
	req.Header.Add("Content-Type", contentType)

	// fmt.Println("headers:")
	// for a, b := range req.Header {
	// 	fmt.Printf(" %s: %s\n", a, b)
	// }

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
