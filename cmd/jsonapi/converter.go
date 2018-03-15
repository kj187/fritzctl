package jsonapi

import "github.com/bpicode/fritzctl/fritz"

type ThermostatsMapper interface {
	Convert([]fritz.Device) ThermostatList
}

type thMapper struct {
}

func (m *thMapper) Convert(ds []fritz.Device) ThermostatList {
	l := ThermostatList{}
	l.NumberOfItems = uint64(len(ds))
	for _, d := range ds {
		l.Thermostats = append(l.Thermostats, m.convertOne(d))
	}
	return l
}
func (m *thMapper) convertOne(d fritz.Device) Thermostat {
	thermostat := Thermostat{}
	thermostat.Identifier = d.Identifier
	thermostat.ID = d.ID
	thermostat.Name = d.Name
	thermostat.FirmwareVersion = d.Fwversion
	thermostat.Manufacturer = d.Manufacturer
	thermostat.ProductName = d.Productname
	return thermostat
}

func NewThermostatsMapper() ThermostatsMapper {
	return &thMapper{}
}
