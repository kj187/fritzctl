package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/bpicode/fritzctl/fritz"
	"github.com/spf13/cobra"
)

var listMonitoringCmd = &cobra.Command{
	Use:     "monitoring",
	Short:   "List available smart home data for monitoring",
	Example: "fritzctl list monitoring --loglevel=error",
	RunE:    listMonitoring,
}

func init() {
	listCmd.AddCommand(listMonitoringCmd)
}

// Monitoring ...
type Monitoring struct {
	Thermostats []Thermostat
	Devices     []Device
}

// Thermostat ...
type Thermostat struct {
	Name                string
	Manufacturer        string
	Productname         string
	MeasuredTemperature string
	Offset              string
	GoalTemperature     string
	SavingTemperature   string
	ComfortTemperature  string
}

// Device ...
type Device struct {
	Name         string
	Manufacturer string
	Productname  string
	PowerW       string
	EnergyWh     string
	Celsius      string
	Offset       string
}

func listMonitoring(_ *cobra.Command, _ []string) error {
	c := homeAutoClient()
	devs, err := c.List()
	assertNoErr(err, "cannot obtain thermostats device data")

	monitoring := &Monitoring{}
	addThermostats(monitoring, devs)
	addSwitches(monitoring, devs)

	b, err := json.Marshal(monitoring)
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println(string(b))

	return nil
}

func addThermostats(monitoring *Monitoring, devs *fritz.Devicelist) {
	for _, dev := range devs.Thermostats() {
		thermostat := Thermostat{
			Name:                dev.Name,
			Manufacturer:        dev.Manufacturer,
			Productname:         dev.Productname,
			MeasuredTemperature: fmt.Sprintf("%s", dev.Thermostat.FmtMeasuredTemperature()),
			Offset:              fmt.Sprintf("%s", dev.Temperature.FmtOffset()),
			GoalTemperature:     fmt.Sprintf("%s", dev.Thermostat.FmtGoalTemperature()),
			SavingTemperature:   fmt.Sprintf("%s", dev.Thermostat.FmtSavingTemperature()),
			ComfortTemperature:  fmt.Sprintf("%s", dev.Thermostat.FmtComfortTemperature()),
		}

		monitoring.Thermostats = append(monitoring.Thermostats, thermostat)
	}
}

func addSwitches(monitoring *Monitoring, devs *fritz.Devicelist) {
	for _, dev := range devs.Switches() {
		device := Device{
			Name:         dev.Name,
			Manufacturer: dev.Manufacturer,
			Productname:  dev.Productname,
			PowerW:       fmt.Sprintf("%s", dev.Powermeter.FmtPowerW()),
			EnergyWh:     fmt.Sprintf("%s", dev.Powermeter.FmtEnergyWh()),
			Celsius:      fmt.Sprintf("%s", dev.Temperature.FmtCelsius()),
			Offset:       fmt.Sprintf("%s", dev.Temperature.FmtOffset()),
		}

		monitoring.Devices = append(monitoring.Devices, device)
	}
}
