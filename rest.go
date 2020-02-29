package lifx

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

func NewSession(token string) *Session {
	return &Session{
		token:  token,
		Client: &http.Client{},
	}
}

func (s *Session) NewRequest(method, url string, body io.Reader) (req *http.Request, err error) {
	req, err = http.NewRequest(method, url, body)
	if err != nil {
		return
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", s.token))
	return
}

func (s *Session) Request(method, url string, body io.Reader) ([]Result, error) {
	req, err := s.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	resp, err := s.Client.Do(req)
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
	return nil, nil
}

func (s *Session) SetState(selector string, state *State) ([]Result, error) {
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

func (s *Session) Toggle(selector string, duration float64) ([]Result, error) {
	m := make(map[string]interface{})
	m["duration"] = duration
	j, err := json.Marshal(m)
	if err != nil {
		return nil, err
	}

	res, err := s.Request("POST", EndpointToggle(selector), bytes.NewBuffer(j))
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *Session) PowerOff(selector string) ([]Result, error) {
	return s.SetState(selector, &State{Power: "off"})

}

func (s *Session) FastPowerOff(selector string) {
	s.SetState(selector, &State{Power: "off", Fast: true})
}

func (s *Session) PowerOn(selector string) ([]Result, error) {
	return s.SetState(selector, &State{Power: "on"})
}

func (s *Session) FastPowerOn(selector string) {
	s.SetState(selector, &State{Power: "on", Fast: true})
}
