package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ListLocations(pageURL *string) (RespShallowLocations, error) {
	url := baseURL + "/location-area"

	if pageURL != nil {
		url = *pageURL
	}

	var locationResp RespShallowLocations

	if cachedData, ok := c.cache.Get(url); ok {
		if err := json.Unmarshal(cachedData, &locationResp); err != nil {
			return RespShallowLocations{}, err
		}
		return locationResp, nil
	}

	res, err := http.Get(url)
	if err != nil {
		return RespShallowLocations{}, err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return RespShallowLocations{}, err
	}

	err = json.Unmarshal(data, &locationResp)
	if err != nil {
		return RespShallowLocations{}, err
	}

	c.cache.Add(url, data)

	return locationResp, nil
}
