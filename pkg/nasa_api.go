package nasa

import (
	"encoding/json"
	"fmt"
	"net/http"

	log "github.com/sirupsen/logrus"
)

const (
	BaseURL = "https://api.nasa.gov/planetary/apod?"
)

type Client struct {
	APIKey string
}

func New(key string) Client {
	return Client{APIKey: key}
}

type QueryOptions struct {
	Date      string
	StartDate string
	EndDate   string
	Count     int
}

func makeRequest(queryUrl string) (*http.Response, error) {

	log.Info("Making request to ", queryUrl)
	resp, err := http.Get(queryUrl)
	log.Info(resp.Status)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Client) GetApod(options QueryOptions) (ApiResponse, error) {
	var apiResp ApiResponse
	var queryUrl string

	// construct query
	if options.Date != "" {
		// create validation for dates in the struct
		queryUrl = fmt.Sprintf("%sapi_key=%s&date=%s", BaseURL, c.APIKey, options.Date)
	} else {
		return ApiResponse{}, fmt.Errorf("only date queryoption is supported with getapod")
	}

	resp, err := makeRequest(queryUrl)
	if err != nil {
		return ApiResponse{}, nil
	}
	defer resp.Body.Close()

	if err = json.NewDecoder(resp.Body).Decode(&apiResp); err != nil {
		return ApiResponse{}, err
	}

	return apiResp, nil
}

func (c *Client) GetApods(options QueryOptions) (ImagesResponse, error) {
	var apiResp ImagesResponse
	var queryUrl string

	if options.StartDate != "" && options.EndDate != "" {
		queryUrl = fmt.Sprintf("%sapi_key=%s&start_date=%s&end_date=%s", BaseURL, c.APIKey, options.StartDate, options.EndDate)
	} else if options.Count != 0 {
		queryUrl = fmt.Sprintf("%sapi_key=%s&count=%d", BaseURL, c.APIKey, options.Count)
	} else {
		return ImagesResponse{}, fmt.Errorf("query option not allowed for getapods")
	}

	resp, err := makeRequest(queryUrl)
	if err != nil {
		return ImagesResponse{}, err
	}

	defer resp.Body.Close()

	if err = json.NewDecoder(resp.Body).Decode(&apiResp.Images); err != nil {
		return ImagesResponse{}, err
	}

	return apiResp, nil
}
