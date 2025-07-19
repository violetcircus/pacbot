package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"net/url"
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
	message := messageParams{
		apiVersion: 10,
	}
	sendMessage(message)
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
