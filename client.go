package lifx

import (
	//"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

const UserAgent = "go-lifx"

type (
	Client struct {
		accessToken string
		Client      *http.Client
	}

	Results struct {
		Results []Result `json:results`
	}

	Result struct {
		ID     string `json:"id"`
		Label  string `json:"label"`
		Status Status `json:"status"`
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

func (c *Client) Request(method, url string, body io.Reader) (*http.Response, error) {
	req, err := c.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}

	switch resp.StatusCode {
	case http.StatusOK, http.StatusAccepted, http.StatusMultiStatus:
		return resp, nil
	}

	err, ok := errorMap[resp.StatusCode]
	if ok {
		return resp, err
	}

	return resp, nil
}

func (c *Client) UnmarshalResponse(resp *http.Response, s interface{}) error {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, &s)
	if err != nil {
		return err
	}

	return nil
}
