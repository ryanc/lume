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
