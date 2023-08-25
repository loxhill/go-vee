package govee

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

const Version = "0.0.1"

func New(apiKey string) *Client {
	return &Client{
		APIKey: apiKey,
	}
}

type Client struct {
	APIKey string
}

func (c *Client) Run(request GoveeRequest) (GoveeResponse, error) {
	client := &http.Client{}

	var req *http.Request
	switch request.GetMethod() {
	case "GET":
		req = c.getRequest(request)
	case "PUT":
		req = c.putRequest(request)
	}

	req.Header.Set("Govee-API-Key", c.APIKey)
	req.Header.Set("User-Agent", fmt.Sprintf("go-vee/%s", Version))
	resp, err := client.Do(req)
	if err != nil {
		return GoveeResponse{}, fmt.Errorf("govee request error: %w", err)
	}
	defer resp.Body.Close()

	var response GoveeResponse
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return GoveeResponse{}, fmt.Errorf("cannot parse govee response: %w", err)
	}
	if resp.StatusCode != 200 {
		return GoveeResponse{}, fmt.Errorf("govee request error: %s", response.Message)
	}

	return response, nil
}

func (c *Client) getRequest(request GoveeRequest) *http.Request {
	params := request.GetParams()
	url := "https://developer-api.govee.com" + request.GetEndpoint()
	if len(params) > 0 {
		url += "?"
		for key, value := range params {
			url += key + "=" + value + "&"
		}
		url = strings.TrimSuffix(url, "&")
	}
	req, _ := http.NewRequest(request.GetMethod(), url, nil)
	return req
}

func (c *Client) putRequest(request GoveeRequest) *http.Request {
	jsonBody, _ := json.Marshal(request.GetBody())
	req, _ := http.NewRequest(request.GetMethod(), "https://developer-api.govee.com"+request.GetEndpoint(), bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	return req
}

type GoveeRequest interface {
	GetEndpoint() string
	GetMethod() string
	GetBody() interface{}
	GetParams() map[string]string
}

type GoveeResponse struct {
	Code    int          `json:"code"`
	Message string       `json:"message"`
	Data    ResponseData `json:"data"`
}

func (g GoveeResponse) Devices() Devices {
	return g.Data.Devices
}
func (g GoveeResponse) Device() string {
	return g.Data.Device
}
func (g GoveeResponse) Model() string {
	return g.Data.Model
}
func (g GoveeResponse) Properties() []map[string]interface{} {
	return g.Data.Properties
}

type ResponseData struct {
	Device     string                   `json:"device"`
	Model      string                   `json:"model"`
	Properties []map[string]interface{} `json:"properties"`
	Devices    []Device                 `json:"devices"`
}
