package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

const (
	contentType    string = "application/json"
	userAgent      string = "pacbot https://github.com/violetcircus/pacbot"
	gatewayOptions string = "?v=10&encoding=json"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type DiscordGateway struct {
	conn      *websocket.Conn
	sequence  int
	sessionID string
	ready     bool
}

func main() {
	// cliMessage()
}

func loadEnv() map[string]string {
	f, err := os.Open(".env")
	if err != nil {
		log.Fatal("error reading envs file", err)
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	s.Split(bufio.ScanLines)

	envs := make(map[string]string)
	for s.Scan() {
		a := strings.Split(s.Text(), "=")
		envs[a[0]] = a[1]
	}

	return envs
}

func getGateway() string {
	type gatewayResponse struct {
		Url string `json:"url"`
	}

	resp, err := http.Get("https://discord.com/api/gateway")
	if err != nil {
		log.Fatal("error getting gateway URL", err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	var result gatewayResponse
	err = json.Unmarshal(body, &result)
	return fmt.Sprintf("%s/%s", result.Url, gatewayOptions)
}
