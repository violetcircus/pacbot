package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"io"
	"log"
	"net/http"
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

func getGateway() string {
	resp, err := http.Get("https://discord.com/api/gateway")
	if err != nil {
		log.Fatal("error getting gateway URL", err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	result := struct {
		Url string `json:"url"`
	}{}
	err = json.Unmarshal(body, &result)
	return fmt.Sprintf("%s/%s", result.Url, gatewayOptions)
}

func newDiscordGateway() (*DiscordGateway, error) {
	url := getGateway()

	conn, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to gateway: %w", err)
	}

	return &DiscordGateway{
		conn: conn,
	}, nil

}
