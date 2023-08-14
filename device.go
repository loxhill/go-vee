package govee

import (
	"errors"
	"fmt"
)

var (
	ErrUnsupportedCmd = errors.New("unsupported command")
)

type DeviceControlRequest struct {
	Endpoint string
	Method   string
	Body     DeviceControlBody
}

func (d DeviceControlRequest) GetEndpoint() string {
	return d.Endpoint
}
func (d DeviceControlRequest) GetMethod() string {
	return d.Method
}
func (d DeviceControlRequest) GetBody() interface{} {
	return d.Body
}

type DeviceControlBody struct {
	Device string `json:"device"`
	Model  string `json:"model"`
	Cmd    Cmd    `json:"cmd"`
}

type Cmd struct {
	Name  string      `json:"name"`
	Value interface{} `json:"value"`
}

type DeviceListRequest struct {
	Endpoint string
	Method   string
}

func (d DeviceListRequest) GetEndpoint() string {
	return d.Endpoint
}
func (d DeviceListRequest) GetMethod() string {
	return d.Method
}
func (d DeviceListRequest) GetBody() interface{} {
	return nil
}

func (c *Client) Device(device string, model string) *Device {
	return &Device{
		Device:      device,
		Model:       model,
		SupportCmds: []string{"turn", "brightness", "color", "colorTem"},
	}
}

func (c *Client) ListDevices() DeviceListRequest {
	return DeviceListRequest{
		Endpoint: "/v1/devices",
		Method:   "GET",
	}
}

type Devices []Device

type Device struct {
	Device       string     `json:"device"`
	Model        string     `json:"model"`
	DeviceName   string     `json:"deviceName"`
	Controllable bool       `json:"controllable"`
	Retrievable  bool       `json:"retrievable"`
	SupportCmds  []string   `json:"supportCmds"`
	Properties   Properties `json:"properties"`
}

func (d *Device) TurnOn() (DeviceControlRequest, error) {
	return d.Control("turn", "on")
}

func (d *Device) TurnOff() (DeviceControlRequest, error) {
	return d.Control("turn", "off")
}

func (d *Device) SetBrightness(brightness int) (DeviceControlRequest, error) {
	if brightness < 0 || brightness > 100 {
		return DeviceControlRequest{}, fmt.Errorf("brightness must be between 0 and 100")
	}
	return d.Control("brightness", brightness)
}

func (d *Device) SetColor(color Color) (DeviceControlRequest, error) {
	return d.Control("color", color)
}

func (d *Device) SetColorTem(colorTem int) (DeviceControlRequest, error) {
	if colorTem < 1000 || colorTem > 10000 {
		return DeviceControlRequest{}, fmt.Errorf("colorTem must be between 1000 and 10000")
	}
	return d.Control("colorTem", colorTem)
}

func (d *Device) Control(cmd string, value interface{}) (DeviceControlRequest, error) {
	if !d.isCmdSupported(cmd) {
		return DeviceControlRequest{}, ErrUnsupportedCmd
	}
	return DeviceControlRequest{
		Method: "PUT",
		Body: DeviceControlBody{
			Device: d.Device,
			Model:  d.Model,
			Cmd: Cmd{
				Name:  cmd,
				Value: value,
			},
		},
		Endpoint: "/v1/devices/control",
	}, nil
}

func (d *Device) isCmdSupported(cmd string) bool {
	for _, supportedCmd := range d.SupportCmds {
		if supportedCmd == cmd {
			return true
		}
	}
	return false
}

type Properties struct {
	ColorTem ColorTem `json:"colorTem"`
}

type ColorTem struct {
	Range Range `json:"range"`
}

type Range struct {
	Min int `json:"min"`
	Max int `json:"max"`
}
