package lifx

import (
	"bytes"
	//"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

var errorMap = map[int]error{
	http.StatusNotFound:            errors.New("Selector did not match any lights"),
	http.StatusUnauthorized:        errors.New("Bad access token"),
	http.StatusForbidden:           errors.New("Bad OAuth scope"),
	http.StatusUnprocessableEntity: errors.New("Missing or malformed parameters"),
	http.StatusUpgradeRequired:     errors.New("HTTP was used to make the request instead of HTTPS. Repeat the request using HTTPS instead"),
	http.StatusTooManyRequests:     errors.New("The request exceeded a rate limit"),
	http.StatusInternalServerError: errors.New("Something went wrong on LIFX's end"),
	http.StatusBadGateway:          errors.New("Something went wrong on LIFX's end"),
	http.StatusServiceUnavailable:  errors.New("Something went wrong on LIFX's end"),
	523:                            errors.New("Something went wrong on LIFX's end"),
}

func NewClient(accessToken string) *Client {
	tr := &http.Transport{
		//TLSNextProto: make(map[string]func(authority string, c *tls.Conn) http.RoundTripper),
	}
	return &Client{
		accessToken: accessToken,
		Client:      &http.Client{Transport: tr},
	}
}

func (c *Client) NewRequest(method, url string, body io.Reader) (req *http.Request, err error) {
	req, err = http.NewRequest(method, url, body)
	if err != nil {
		return
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.accessToken))
	return
}

func (c *Client) Request(method, url string, body io.Reader) ([]Result, error) {
	req, err := c.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}

	switch resp.StatusCode {
	case http.StatusAccepted:
		return nil, nil
	case http.StatusMultiStatus:
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}

		r := Results{}
		err = json.Unmarshal(body, &r)
		if err != nil {
			return nil, err
		}

		return r.Results, nil
	}

	err, ok := errorMap[resp.StatusCode]
	if ok {
		return nil, err
	}

	return nil, nil
}

func (c *Client) SetState(selector string, state State) ([]Result, error) {
	j, err := json.Marshal(state)
	if err != nil {
		return nil, err
	}

	fmt.Println(string(j))

	res, err := c.Request("PUT", EndpointState(selector), bytes.NewBuffer(j))
	if err != nil {
		return nil, err
	}

	return res, nil
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

	res, err := c.Request("PUT", EndpointStates(), bytes.NewBuffer(j))
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (c *Client) Toggle(selector string, duration float64) ([]Result, error) {
	j, err := json.Marshal(&Toggle{Duration: duration})
	if err != nil {
		return nil, err
	}

	res, err := c.Request("POST", EndpointToggle(selector), bytes.NewBuffer(j))
	if err != nil {
		return nil, err
	}

	return res, nil
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
