package govee

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

const Version = "0.6.0"

func New(apiKey string) *Client {
	return &Client{
		APIKey: apiKey,
	}
}

type Client struct {
	APIKey string
}

// ListDevices fetches a list of devices associated with your Govee account.
func (c *Client) GetDevices() ([]Device, error) {
	req := DiscoverDevicesRequest{
		Endpoint: "/router/api/v1/user/devices",
		Method:   "GET",
	}
	resp, err := c.run(req)
	if err != nil {
		return nil, fmt.Errorf("could not get devices: %v", err)
	}
	return resp.Data, nil
}

// GetDeviceState fetches the state of a single device.
func (c *Client) GetDeviceState(device Device) (Device, error) {
	req := DeviceStateRequest{
		Endpoint: "/router/api/v1/device/state",
		Method:   "POST",
		Body: DeviceStateRequestBody{
			RequestID: "test",
			Payload: DeviceStateRequestBodyPayload{
				SKU:    device.SKU,
				Device: device.Device,
			},
		},
	}
	resp, err := c.run(req)
	if err != nil {
		return Device{}, fmt.Errorf("could not get device state for device %v: %v", device.Device, err)
	}
	return resp.Payload, nil
}

// ControlDevice sends control instructions for a particular device ID.
func (c *Client) ControlDevice(device Device, capabilityType, instance string, value interface{}) error {
	req := DeviceControlRequest{
		Endpoint: "/router/api/v1/device/control",
		Method:   "POST",
		Body: DeviceControlRequestBody{
			RequestID: "test",
			Payload: DeviceControlRequestBodyPayload{
				SKU:    device.SKU,
				Device: device.Device,
				Capability: DeviceControlRequestCapability{
					Type:     capabilityType,
					Instance: instance,
					Value:    value,
				},
			},
		},
	}
	resp, err := c.run(req)
	if err != nil {
		return fmt.Errorf("could not control device %v: %v", device.Device, err)
	}
	if resp.Code != 200 {
		return fmt.Errorf("could not control device %v: %v", device.Device, resp.Msg)
	}
	return nil
}

func (c *Client) run(request GoveeRequest) (GoveeResponse, error) {
	client := &http.Client{}

	var req *http.Request
	switch request.GetMethod() {
	case "GET":
		req = c.getRequest(request)
	case "POST":
		req = c.postRequest(request)
	}

	req.Header.Set("Govee-API-Key", c.APIKey)
	req.Header.Set("User-Agent", fmt.Sprintf("go-vee/%s", Version))
	resp, err := client.Do(req)
	if resp.StatusCode != 200 {
		return GoveeResponse{}, fmt.Errorf("govee request error, non-200 response: %w", err)
	}
	if err != nil {
		return GoveeResponse{}, fmt.Errorf("govee request error: %w", err)
	}
	defer resp.Body.Close()

	var response GoveeResponse
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return GoveeResponse{}, fmt.Errorf("cannot parse govee response: %w", err)
	}
	return response, nil
}

func (c *Client) getRequest(request GoveeRequest) *http.Request {
	url := "https://openapi.api.govee.com" + request.GetEndpoint()
	req, _ := http.NewRequest(request.GetMethod(), url, nil)
	return req
}

func (c *Client) postRequest(request GoveeRequest) *http.Request {
	jsonBody, _ := json.Marshal(request.GetBody())
	req, _ := http.NewRequest(request.GetMethod(), "https://openapi.api.govee.com"+request.GetEndpoint(), bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	return req
}

type GoveeRequest interface {
	GetEndpoint() string
	GetMethod() string
	GetBody() interface{}
}

// GoveeResponse is the struct representing the response back from all Govee
// endpoints. Unfortunately, some endpoints return different structures hence
// the different parameters.
type GoveeResponse struct {
	RequestID string   `json:"requestId"`
	Code      int      `json:"code"`
	Msg       string   `json:"msg"`
	Message   string   `json:"message"`
	Data      []Device `json:"data"`
	Payload   Device   `json:"payload"`
}
