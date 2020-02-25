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
		return []Result{}, err
	}

	resp, err := s.Client.Do(req)
	if err != nil {
		return []Result{}, err
	}

	switch resp.StatusCode {
	case http.StatusAccepted:
		return []Result{}, nil
	case http.StatusMultiStatus:
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return []Result{}, err
		}

		r := Results{}
		err = json.Unmarshal(body, &r)
		if err != nil {
			return []Result{}, err
		}

		return r.Results, nil
	}
	return []Result{}, nil
}

func (s *Session) SetState(selector string, state *State) ([]Result, error) {
	j, err := json.Marshal(state)
	if err != nil {
		return []Result{}, err
	}

	res, err := s.Request("PUT", EndpointState(selector), bytes.NewBuffer(j))
	if err != nil {
		return []Result{}, err
	}

	return res, nil
}

func (s *Session) PowerOff(selector string) {
	s.SetState(selector, &State{Power: "off"})
}
