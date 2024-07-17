package govee

type DeviceControlRequest struct {
	Endpoint string
	Method   string
	Body     DeviceControlRequestBody
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

type DeviceControlRequestBody struct {
	RequestID string                          `json:"requestId"`
	Payload   DeviceControlRequestBodyPayload `json:"payload"`
}

type DeviceControlRequestBodyPayload struct {
	SKU        string                         `json:"sku"`
	Device     string                         `json:"device"`
	Capability DeviceControlRequestCapability `json:"capability"`
}

type DeviceControlRequestCapability struct {
	Type     string      `json:"type"`
	Instance string      `json:"instance"`
	Value    interface{} `json:"value"`
}
