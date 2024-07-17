package govee

type DiscoverDevicesRequest struct {
	Endpoint string
	Method   string
}

func (d DiscoverDevicesRequest) GetEndpoint() string {
	return d.Endpoint
}
func (d DiscoverDevicesRequest) GetMethod() string {
	return d.Method
}
func (d DiscoverDevicesRequest) GetBody() interface{} {
	return nil
}
