package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
)

type messageParams struct {
	apiVersion int64
	parameters map[string]string
	token      string
	channelID  string
	message    string
}

const contentType string = "application/json"
const userAgent string = "pacbot https://github.com/violetcircus/pacbot"

func main() {
	envs := loadEnv()
	// get user input from terminal for message + channel id
	channelID, message := getInput()
	msg := messageParams{
		apiVersion: 10,
		token:      envs["TOKEN"],
		channelID:  channelID,
		message:    message,
	}
	sendMessage(msg)
}

// load envs. doing it this way is dumb: use normal file reading and just string manip the lines into a struct lol
func loadEnv() map[string]string {
	f, err := os.Open("./.env")
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

func getInput() (string, string) {
	s := bufio.NewScanner(os.Stdin)
	var buf string

	// get channel id
	var channelID string
	fmt.Println("enter channel ID:")
	s.Scan()
	buf = s.Text()
	if len(buf) != 0 {
		channelID = buf
	}

	// get message
	var userMessage string
	fmt.Println("enter message:")
	s.Scan()
	buf = s.Text()
	if len(buf) != 0 {
		userMessage = buf
	}
	return channelID, userMessage
}

func sendMessage(msg messageParams) {
	targetUrl := fmt.Sprintf("https://discordapp.com/api/channels/%s/messages", strconv.FormatInt(msg.apiVersion, 10))

	// assemble params into a suitable format for post request
	postBody := url.Values{} // creates empty map (std. lib thing for handling URL-encoded form data)
	for a, b := range msg.parameters {
		path, err := url.PathUnescape(b) // unescapes value in parameters key:value pair for use in form
		if err != nil {
			log.Fatal(err)
		}
		postBody.Set(a, path) // adds the key, value pair to the post body with the value now unescaped for use in a form
	}

	req, error := http.NewRequest("POST", targetUrl, strings.NewReader(postBody.Encode()))
	if error != nil {
		log.Fatal("request failed lol")
	}
	req.Header.Add("Authorization", msg.token)
	req.Header.Add("User-Agent", userAgent)
	req.Header.Add("Content-Type", contentType)
}
