package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

const (
	contentType    string = "application/json"
	userAgent      string = "pacbot https://github.com/violetcircus/pacbot"
	gatewayOptions string = "?v=10&encoding=json"
)

func main() {
	cliMessage()
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
