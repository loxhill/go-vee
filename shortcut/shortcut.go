package shortcut

type Shortcut struct {
	CapabilityType string
	Instance       string
	Value          interface{}
}

func TurnOn() Shortcut {
	return Shortcut{
		CapabilityType: "devices.capabilities.on_off",
		Instance:       "powerSwitch",
		Value:          1,
	}
}

func TurnOff() Shortcut {
	return Shortcut{
		CapabilityType: "devices.capabilities.on_off",
		Instance:       "powerSwitch",
		Value:          0,
	}
}
