package main

import (
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/willoong9559/Gotunnel/tunnel"
)

var (
	listenAddr string
	serverAddr string
	obfsType   string
	obfsParam  string
	psk        string
)

func init() {
	flag.StringVar(&listenAddr, "listen", ":1080", "listen address")
	flag.StringVar(&serverAddr, "server", "", "server address")
	flag.StringVar(&obfsType, "obfs", "", "obfs type")
	flag.StringVar(&obfsParam, "obfs-param", "", "obfs param")
	flag.StringVar(&psk, "psk", "", "psk")
	flag.Parse()

	if serverAddr == "" {
		log.Fatalf("server address is empty")
	}

	if obfsType == "" && obfsParam != "" {
		log.Fatalf("obfs type is empty")
	}

	if obfsType == "none" || obfsType == "off" {
		obfsType = "" // disable obfs
	}
}

func main() {
	s, err := tunnel.newClient(listenAddr, serverAddr, obfsType, obfsParam, psk)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	<-sig

	s.Close()
}
