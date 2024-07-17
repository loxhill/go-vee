package govee

type Devices []Device

type Device struct {
	SKU          string       `json:"sku"`
	Device       string       `json:"device"`
	Capabilities []Capability `json:"capabilities"`
}

type Capability struct {
	Type       string               `json:"type"`
	Instance   string               `json:"instance"`
	Parameters CapabilityParameters `json:"parameters"`
	State      CapabilityState      `json:"state"`
}

type CapabilityParameters struct {
	DataType string                      `json:"dataType"`
	Options  []CapabilityParameterOption `json:"options"`
}

type CapabilityParameterOption struct {
	Name  string `json:"name"`
	Value int16  `json:"value"`
}

type CapabilityState struct {
	Value interface{} `json:"value"`
}
