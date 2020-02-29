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

func NewClient(token string) *Client {
	tr := &http.Transport{
		//TLSNextProto: make(map[string]func(authority string, c *tls.Conn) http.RoundTripper),
	}
	return &Client{
		token:  token,
		Client: &http.Client{Transport: tr},
	}
}

func (s *Client) NewRequest(method, url string, body io.Reader) (req *http.Request, err error) {
	req, err = http.NewRequest(method, url, body)
	if err != nil {
		return
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", s.token))
	return
}

func (s *Client) Request(method, url string, body io.Reader) ([]Result, error) {
	req, err := s.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	resp, err := s.Client.Do(req)
	if err != nil {
		return nil, err
	}

	switch resp.StatusCode {
	case http.StatusNotFound:
		return nil, errors.New("Selector did not match any lights")
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
	return nil, nil
}

func (s *Client) SetState(selector string, state State) ([]Result, error) {
	j, err := json.Marshal(state)
	if err != nil {
		return nil, err
	}

	res, err := s.Request("PUT", EndpointState(selector), bytes.NewBuffer(j))
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *Client) SetStates(states States) ([]Result, error) {
	j, err := json.Marshal(states)
	if err != nil {
		return nil, err
	}

	res, err := s.Request("PUT", EndpointStates(), bytes.NewBuffer(j))
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *Client) Toggle(selector string, duration float64) ([]Result, error) {
	j, err := json.Marshal(&Toggle{Duration: duration})
	if err != nil {
		return nil, err
	}

	res, err := s.Request("POST", EndpointToggle(selector), bytes.NewBuffer(j))
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *Client) PowerOff(selector string) ([]Result, error) {
	return s.SetState(selector, State{Power: "off"})

}

func (s *Client) FastPowerOff(selector string) {
	s.SetState(selector, State{Power: "off", Fast: true})
}

func (s *Client) PowerOn(selector string) ([]Result, error) {
	return s.SetState(selector, State{Power: "on"})
}

func (s *Client) FastPowerOn(selector string) {
	s.SetState(selector, State{Power: "on", Fast: true})
}
