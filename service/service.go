package service

import (
	"log"
	"os"

	"mini-zone-service/server"
	"mini-zone-service/server/context"

	"github.com/joho/godotenv"
)

func parseParams() (
	context.Metadata,
	string,
	context.ZoneList,
) {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	name := os.Getenv("SERVICE_NAME")
	if len(name) == 0 {
		log.Fatal("env param SERVICE_NAME is empty")
	}
	ver := os.Getenv("VERSION")
	if len(ver) == 0 {
		log.Fatal("env param VERSION is empty")
	}
	at := os.Getenv("SERVER_LISTEN_AT")
	if len(at) == 0 {
		log.Fatal("env param SERVER_LISTEN_AT is empty")
	}

	addr := os.Getenv("ZONE_ADDRESS")
	if len(addr) == 0 {
		log.Fatal("env param ZONE_ADDRESS is empty")
	}

	home := os.Getenv("HOME_PORT")
	if len(home) == 0 {
		log.Fatal("env param HOME_URL is empty")
	}

	zone1 := os.Getenv("ZONE1_PORT")
	if len(zone1) == 0 {
		log.Fatal("env param ZONE1_URL is empty")
	}

	zone2 := os.Getenv("ZONE2_PORT")
	if len(zone2) == 0 {
		log.Fatal("env param ZONE2_URL is empty")
	}

	return context.Metadata{
			Name:    name,
			Version: ver,
		},
		at,
		context.ZoneList{
			Address:   addr,
			HomePort:  home,
			Zone1Port: zone1,
			Zone2Port: zone2,
		}
}

func Run() {
	server.Run(parseParams())
}
