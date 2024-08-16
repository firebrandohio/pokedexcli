package pokeapi

import (
	"net/http"
	"time"
)

const baseURL = "http://www.pokeapi.co/api/v2"

type Client struct {
	httpClient http.Client
}

func NewClient() Client {
	return Client{
		httpClient: http.Client{
			Timeout: time.Minute * 2,
		},
	}
}
