package lifx

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func NewSession(token string) *Session {
	return &Session{
		token:  token,
		Client: &http.Client{},
	}
}

func (s *Session) NewRequest(method, url string, body io.Reader) *http.Request {
	req, _ := http.NewRequest(method, url, body)
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", s.token))
	return req
}

func (s *Session) SetState(selector string, state *State) error {
	j, _ := json.Marshal(state)
	req := s.NewRequest("PUT", EndpointState(selector), bytes.NewBuffer(j))
	resp, err := s.Client.Do(req)
	if err != nil {
		return err
	}
	fmt.Println(resp)
	return nil
}
