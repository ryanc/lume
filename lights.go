package lifx

import (
	"bytes"
	//"crypto/tls"
	"encoding/json"
)

const (
	OK       Status = "ok"
	TimedOut Status = "timed_out"
	Offline  Status = "offline"
)

type (
	Status string

	State struct {
		Power      string  `json:"power,omitempty"`
		Color      Color   `json:"color,omitempty"`
		Brightness float64 `json:"brightness,omitempty"`
		Duration   float64 `json:"duration,omitempty"`
		Infrared   float64 `json:"infrared,omitempty"`
		Fast       bool    `json:"fast,omitempty"`
	}

	StateWithSelector struct {
		State
		Selector string `json:"selector"`
	}

	States struct {
		States   []StateWithSelector `json:"states",omitempty`
		Defaults State               `json:"defaults",omitempty`
	}

	Toggle struct {
		Duration float64 `json:"duration,omitempty"`
	}
)

func (s Status) Success() bool {
	return s == OK
}

func (c *Client) SetState(selector string, state State) ([]Result, error) {
	j, err := json.Marshal(state)
	if err != nil {
		return nil, err
	}

	resp, err := c.Request("PUT", EndpointState(selector), bytes.NewBuffer(j))
	if err != nil {
		return nil, err
	}

	if state.Fast {
		return nil, nil
	}

	s := &Results{}
	err = c.UnmarshalResponse(resp, s)
	if err != nil {
		return nil, err
	}

	return s.Results, nil
}

func (c *Client) FastSetState(selector string, state State) ([]Result, error) {
	state.Fast = true
	return c.SetState(selector, state)
}

func (c *Client) SetStates(states States) ([]Result, error) {
	j, err := json.Marshal(states)
	if err != nil {
		return nil, err
	}

	resp, err := c.Request("PUT", EndpointStates(), bytes.NewBuffer(j))
	if err != nil {
		return nil, err
	}

	s := &Results{}
	err = c.UnmarshalResponse(resp, s)
	if err != nil {
		return nil, err
	}

	return s.Results, nil
}

func (c *Client) Toggle(selector string, duration float64) ([]Result, error) {
	j, err := json.Marshal(&Toggle{Duration: duration})
	if err != nil {
		return nil, err
	}

	resp, err := c.Request("POST", EndpointToggle(selector), bytes.NewBuffer(j))
	if err != nil {
		return nil, err
	}

	s := &Results{}
	err = c.UnmarshalResponse(resp, s)
	if err != nil {
		return nil, err
	}

	return s.Results, nil
}

func (c *Client) PowerOff(selector string) ([]Result, error) {
	return c.SetState(selector, State{Power: "off"})
}

func (c *Client) FastPowerOff(selector string) {
	c.SetState(selector, State{Power: "off", Fast: true})
}

func (c *Client) PowerOn(selector string) ([]Result, error) {
	return c.SetState(selector, State{Power: "on"})
}

func (c *Client) FastPowerOn(selector string) {
	c.SetState(selector, State{Power: "on", Fast: true})
}