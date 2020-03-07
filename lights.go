package lifx

import (
	//"crypto/tls"
	"encoding/json"
	"net/http"
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

func (c *Client) SetState(selector string, state State) (*Response, error) {
	var (
		err  error
		s    *Response
		resp *http.Response
	)

	if resp, err = c.setStateRequest(selector, state); err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if state.Fast && resp.StatusCode == http.StatusAccepted {
		return nil, nil
	}

	if err = json.NewDecoder(resp.Body).Decode(&s); err != nil {
		return nil, err
	}

	return s, nil
}

func (c *Client) FastSetState(selector string, state State) (*Response, error) {
	state.Fast = true
	return c.SetState(selector, state)
}

func (c *Client) SetStates(selector string, states States) (*Response, error) {
	var (
		err  error
		s    *Response
		resp *http.Response
	)

	if resp, err = c.setStatesRequest(selector, states); err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if err = json.NewDecoder(resp.Body).Decode(&s); err != nil {
		return nil, err
	}

	return s, nil
}

func (c *Client) Toggle(selector string, duration float64) (*Response, error) {
	var (
		err  error
		s    *Response
		resp *http.Response
	)

	if resp, err = c.toggleRequest(selector, duration); err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if err = json.NewDecoder(resp.Body).Decode(&s); err != nil {
		return nil, err
	}

	return s, nil
}

func (c *Client) PowerOff(selector string) (*Response, error) {
	return c.SetState(selector, State{Power: "off"})
}

func (c *Client) FastPowerOff(selector string) {
	c.SetState(selector, State{Power: "off", Fast: true})
}

func (c *Client) PowerOn(selector string) (*Response, error) {
	return c.SetState(selector, State{Power: "on"})
}

func (c *Client) FastPowerOn(selector string) {
	c.SetState(selector, State{Power: "on", Fast: true})
}
