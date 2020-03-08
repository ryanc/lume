package lifx

import (
	//"crypto/tls"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

const UserAgent = "go-lifx"

type (
	Client struct {
		accessToken string
		Client      *http.Client
	}

	Result struct {
		ID     string `json:"id"`
		Label  string `json:"label"`
		Status Status `json:"status"`
	}

	Error struct {
		Field   string   `json:"field"`
		Message []string `json:"message"`
	}

	Warning struct {
		Warning string `json:"warning"`
	}

	Response struct {
		Error    string    `json:"error"`
		Errors   []Error   `json:"errors"`
		Warnings []Warning `json:"warnings"`
		Results  []Result  `json:"results"`
	}
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
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("User-Agent", UserAgent)
	return
}

func (c *Client) setState(selector string, state State) (*http.Response, error) {
	var (
		err  error
		j    []byte
		req  *http.Request
		resp *http.Response
	)

	if j, err = json.Marshal(state); err != nil {
		return nil, err
	}

	if req, err = c.NewRequest("PUT", EndpointState(selector), bytes.NewBuffer(j)); err != nil {
		return nil, err
	}

	if resp, err = c.Client.Do(req); err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Client) setStates(selector string, states States) (*http.Response, error) {
	var (
		err  error
		j    []byte
		req  *http.Request
		resp *http.Response
	)

	if j, err = json.Marshal(states); err != nil {
		return nil, err
	}

	if req, err = c.NewRequest("PUT", EndpointStates(), bytes.NewBuffer(j)); err != nil {
		return nil, err
	}

	if resp, err = c.Client.Do(req); err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Client) toggle(selector string, duration float64) (*http.Response, error) {
	var (
		err  error
		j    []byte
		req  *http.Request
		resp *http.Response
	)

	if j, err = json.Marshal(&Toggle{Duration: duration}); err != nil {
		return nil, err
	}

	if req, err = c.NewRequest("POST", EndpointToggle(selector), bytes.NewBuffer(j)); err != nil {
		return nil, err
	}

	if resp, err = c.Client.Do(req); err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Client) validateColor(color Color) (*http.Response, error) {
	var (
		err  error
		req  *http.Request
		resp *http.Response
		q    url.Values
	)

	if req, err = c.NewRequest("GET", EndpointColor(), nil); err != nil {
		return nil, err
	}

	q = req.URL.Query()
	q.Set("string", color.ColorString())
	req.URL.RawQuery = q.Encode()

	if resp, err = c.Client.Do(req); err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Client) listLights(selector string) (*http.Response, error) {
	var (
		err  error
		req  *http.Request
		resp *http.Response
	)

	if req, err = c.NewRequest("GET", EndpointListLights(selector), nil); err != nil {
		return nil, err
	}

	if resp, err = c.Client.Do(req); err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Client) stateDelta(selector string, delta StateDelta) (*http.Response, error) {
	var (
		err  error
		j    []byte
		req  *http.Request
		resp *http.Response
	)

	if j, err = json.Marshal(delta); err != nil {
		return nil, err
	}

	if req, err = c.NewRequest("POST", EndpointStateDelta(selector), bytes.NewBuffer(j)); err != nil {
		return nil, err
	}

	if resp, err = c.Client.Do(req); err != nil {
		return nil, err
	}

	return resp, nil
}
