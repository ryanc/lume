package main

import (
	"fmt"
	"git.kill0.net/chill9/lifx"
	"os"
	"time"
)

func main() {
	apiToken := os.Getenv("LIFX_API_TOKEN")
	if apiToken == "" {
		fmt.Println("LIFX_API_TOKEN is undefined")
		os.Exit(1)
	}
	s := &lifx.State{Power: "on", Color: "white"}
	c := lifx.NewSession(apiToken)
	c.SetState("group:Office", s)
	time.Sleep(10)
	c.SetState("all", &lifx.State{Power: "on", Color: "green"})
	time.Sleep(10)
	c.SetState("all", &lifx.State{Power: "off"})
	fmt.Println(lifx.EndpointState(lifx.Endpoint))
}
