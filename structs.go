package lifx

import (
	"net/http"
)

const API_BASE_URL = "https://api.lifx.com/v1"

type (
	State struct {
		Power      string  `json:"power,omitempty"`
		Color      string  `json:"color,omitempty"`
		Brightness float64 `json:"brightness,omitempty"`
		Duration   float64 `json:"duration,omitempty"`
		Infrared   float64 `json:"infrared,omitempty"`
		Fast       bool    `json:"fast,omitempty"`
	}

	Client struct {
		token  string
		Client *http.Client
	}

	Results struct {
		Results []Result `json:results`
	}

	Result struct {
		ID     string `json:"id"`
		Label  string `json:"label"`
		Status string `json:"status"`
	}
)
