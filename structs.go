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

	Session struct {
		BaseUrl string
		token   string
		Client  *http.Client
	}
)
