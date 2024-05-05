package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocationAreas(pageURL *string) (LocationAreasResponse, error) {
	endpoint := "/location-area"
	fullURL := baseURL + endpoint

	if pageURL != nil {
		fullURL = *pageURL
	}

	data, ok := c.cache.Get(fullURL)
	if ok {
		// cache hit
		locationAreasResponse := LocationAreasResponse{}
		error := json.Unmarshal(data, &locationAreasResponse)
		if error != nil {
			return LocationAreasResponse{}, error
		}

		return locationAreasResponse, nil

	}

	request, error := http.NewRequest("GET", fullURL, nil)
	if error != nil {
		return LocationAreasResponse{}, error
	}
	response, error := c.httpClient.Do(request)
	if error != nil {
		return LocationAreasResponse{}, error
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		return LocationAreasResponse{}, fmt.Errorf("bad status code: %v", response.StatusCode)
	}

	data, error = io.ReadAll(response.Body)
	if error != nil {
		return LocationAreasResponse{}, error
	}

	locationAreasResponse := LocationAreasResponse{}
	error = json.Unmarshal(data, &locationAreasResponse)
	if error != nil {
		return LocationAreasResponse{}, error
	}
	c.cache.Add(fullURL, data)

	return locationAreasResponse, nil
}
