package main

import (
	"fmt"
	"git.kill0.net/chill9/go-lifx"
	"os"
	//"time"
)

func main() {
	accessToken := os.Getenv("LIFX_ACCESS_TOKEN")
	if accessToken == "" {
		fmt.Println("LIFX_ACCESS_TOKEN is undefined")
		os.Exit(1)
	}
	color1 := lifx.RGBColor{R: 235, G: 191, B: 255}
	color2 := lifx.NewHSBKColor()
	color2.H = 27
	color2.S = 1
	color2.B = 0.39
	fmt.Println(color1.Hex())
	fmt.Println(color1.ColorString())
	fmt.Println(color2.ColorString())
	s := lifx.State{Power: "on", Color: color2}
	c := lifx.NewClient(accessToken)
	selector := "group:Office"
	r, err := c.FastSetState(selector, s)
	/*
		time.Sleep(10 * time.Second)
		s.Color = "white"
		res, _ := c.SetState(selector, s)
		fmt.Println(res)
		//c.SetState("all", &lifx.State{Power: "on", Color: "green"})
		time.Sleep(10 * time.Second)
		c.FastPowerOff(selector)
		time.Sleep(10 * time.Second)
		c.FastPowerOn(selector)
		//c.PowerOff("all")
	*/
	fmt.Println(err)
	fmt.Println(r)
	r, err = c.Toggle(selector, 10.0)
	fmt.Println(err)
	fmt.Println(r)
}
