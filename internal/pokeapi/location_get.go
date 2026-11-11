package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) GetLocation(locationName string) (Location, error) {
	url := baseURL + "/location-area/" + locationName

	var loc Location

	if cachedData, ok := c.cache.Get(url); ok {
		if err := json.Unmarshal(cachedData, &loc); err != nil {
			return Location{}, err
		}
		return loc, nil
	}

	res, err := http.Get(url)
	if err != nil {
		return Location{}, err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return Location{}, err
	}

	err = json.Unmarshal(data, &loc)
	if err != nil {
		return Location{}, err
	}

	c.cache.Add(url, data)

	return loc, nil
}
