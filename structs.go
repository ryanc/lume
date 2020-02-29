package lifx

import (
	"net/http"
)

const API_BASE_URL = "https://api.lifx.com/v1"

type (
	Status string

	State struct {
		Power      string  `json:"power,omitempty"`
		Color      string  `json:"color,omitempty"`
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

	Toggle struct {
		Duration float64 `json:"duration,omitempty"`
	}
)

const (
	OK       Status = "ok"
	TimedOut Status = "timed_out"
	Offline  Status = "offline"
)

func (s Status) Success() bool {
	return s == OK
}
