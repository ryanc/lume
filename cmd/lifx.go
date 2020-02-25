package main

import (
	"fmt"
	"git.kill0.net/chill9/go-lifx"
	"os"
	"time"
)

func main() {
	apiToken := os.Getenv("LIFX_API_TOKEN")
	if apiToken == "" {
		fmt.Println("LIFX_API_TOKEN is undefined")
		os.Exit(1)
	}
	s := &lifx.State{Power: "on", Color: "blue"}
	c := lifx.NewSession(apiToken)
	c.SetState("group:Office", s)
	time.Sleep(10 * time.Second)
	s.Color = "white"
	res, _ := c.SetState("group:Office", s)
	fmt.Println(res)
	//c.SetState("all", &lifx.State{Power: "on", Color: "green"})
	time.Sleep(10)
	//c.PowerOff("all")
}
