package jsonapi

// NextChange corresponds to the next HKR switch event.
type ThermostatList struct {
	NumberOfItems uint64       `json:"numberOfItems"`
	Thermostats   []Thermostat `json:"thermostats"` // The temperature to switch to. Same unit convention as in Thermostat.Measured.
}

// Device models a smart home device. This corresponds to
// the single entries of the json that the FRITZ!Box returns.
// codebeat:disable[TOO_MANY_IVARS]
type Thermostat struct {
	Identifier          string     `json:"identifier"`          // A unique ID like AIN, MAC address, etc.
	ID                  string     `json:"id"`                  // Internal device ID of the FRITZ!Box.
	FirmwareVersion     string     `json:"firmwareVersion"`     // Firmware version of the device.
	Manufacturer        string     `json:"manufacturer"`        // Manufacturer of the device, usually set to "AVM".
	ProductName         string     `json:"productName"`         // Name of the product, empty for unknown or undefined devices.
	Connected           bool       `json:"connected"`           // Device connected (1) or not (0).
	Name                string     `json:"name"`                // The name of the device. Can be assigned in the web gui of the FRITZ!Box. Temperature     Temperature `json:"temperature"`     // Only filled with sensible data for devices with a temperature sensor.
	MeasuredTemperature float64    `json:"measuredTemperature"` // Temperature measured in °C.
	GoalTemperature     string     `json:"goalTemperature"`     // Desired temperature, user controlled.
	SavingTemperature   string     `json:"savingTemperature"`   // Energy saving temperature.
	ComfortTemperature  string     `json:"comfortTemperature"`  // Comfortable temperature.
	NextChange          NextChange `json:"nextChange"`          // The next scheduled temperature change.
	Lock                string     `json:"lock"`                // Switch locked (box defined)? 1/0 (empty if not known or if there was an error).
	DeviceLock          string     `json:"devicelock"`          // Switch locked (device defined)? 1/0 (empty if not known or if there was an error).
	ErrorCode           string     `json:"errorCode"`           // Error codes: 0 = OK, 1 = ... see https://avm.de/fileadmin/user_upload/Global/Service/Schnittstellen/AHA-HTTP-Interface.pdf.
	BatteryLow          string     `json:"batteryLow"`          // "0" if the battery is OK, "1" if it is running low on capacity.
}

// NextChange corresponds to the next HKR switch event.
type NextChange struct {
	At          string `json:"timestamp"`   // Timestamp (epoch time) when the next temperature switch is scheduled.
	Temperature string `json:"temperature"` // The temperature to switch to. Same unit convention as in Thermostat.Measured.
}

// codebeat:enable[TOO_MANY_IVARS]
