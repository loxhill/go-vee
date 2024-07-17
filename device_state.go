package govee

type DeviceStateRequest struct {
	Endpoint string
	Method   string
	Body     DeviceStateRequestBody
}

type DeviceStateRequestBody struct {
	RequestID string                        `json:"requestId"`
	Payload   DeviceStateRequestBodyPayload `json:"payload"`
}

type DeviceStateRequestBodyPayload struct {
	SKU    string `json:"sku"`
	Device string `json:"device"`
}

func (d DeviceStateRequest) GetEndpoint() string {
	return d.Endpoint
}
func (d DeviceStateRequest) GetMethod() string {
	return d.Method
}
func (d DeviceStateRequest) GetBody() interface{} {
	return d.Body
}
