package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/bpicode/fritzctl/cmd/jsonapi"
	"github.com/bpicode/fritzctl/cmd/printer"
	"github.com/bpicode/fritzctl/console"
	"github.com/bpicode/fritzctl/fritz"
	"github.com/bpicode/fritzctl/logger"
	"github.com/bpicode/fritzctl/stringutils"
	"github.com/spf13/cobra"
)

var listThermostatsCmd = &cobra.Command{
	Use:     "thermostats",
	Short:   "List the available smart home thermostats",
	Long:    "List the available smart home devices [thermostats] and associated data.",
	Example: "fritzctl list thermostats",
	RunE:    listThermostats,
}

func init() {
	listThermostatsCmd.Flags().StringP("output", "o", "", "specify output format")
	listCmd.AddCommand(listThermostatsCmd)
}

func listThermostats(cmd *cobra.Command, _ []string) error {
	c := homeAutoClient()
	devs, err := c.List()
	assertNoErr(err, "cannot obtain thermostats device data")
	data, err := remapThermostats(cmd, devs.Thermostats())
	assertNoErr(err, "failed to print results for thermostats device data")
	logger.Success("Device data:")
	printer.Print(data, os.Stdout)
	return nil
}

func remapThermostats(cmd *cobra.Command, ds []fritz.Device) (interface{}, error) {
	o := cmd.Flag("output").Value.String()
	switch o {
	case "":
		t := thermostatsTable()
		appendThermostats(ds, t)
		return t, nil
	case "json":
		mapper := jsonapi.NewThermostatsMapper()
		return mapper.Convert(ds), nil
	default:
		return nil, fmt.Errorf("no output engine found for '%s'", o)
	}
}

var errorCodesVsDescriptions = map[string]string{
	"":  "",
	"0": "",
	"1": " Thermostat adjustment not possible. Is the device mounted correctly?",
	"2": " Valve plunger cannot be driven far enough. Possible solutions: Open and close the plunger a couple of times by hand. Check if the battery is too weak.",
	"3": " Valve plunger cannot be moved. Is it blocked?",
	"4": " Preparing installation.",
	"5": " Device in mode 'INSTALLATION'. It can be mounted now.",
	"6": " Device is adjusting to the valve plunger.",
}

func thermostatsTable() *console.Table {
	return console.NewTable(console.Headers(
		"NAME",
		"PRODUCT",
		"PRESENT",
		"LOCK (BOX/DEV)",
		"MEASURED",
		"OFFSET",
		"WANT",
		"SAVING",
		"COMFORT",
		"NEXT",
		"STATE",
		"BATTERY",
	))
}

func appendThermostats(devs []fritz.Device, table *console.Table) {
	for _, dev := range devs {
		columns := thermostatColumns(dev)
		table.Append(columns)
	}
}

func thermostatColumns(dev fritz.Device) []string {
	var columnValues []string
	columnValues = appendMetadata(columnValues, dev)
	columnValues = appendRuntimeFlags(columnValues, dev)
	columnValues = appendTemperatureValues(columnValues, dev)
	columnValues = appendRuntimeWarnings(columnValues, dev)
	return columnValues
}

func appendMetadata(cols []string, dev fritz.Device) []string {
	return append(cols, dev.Name, fmt.Sprintf("%s %s", dev.Manufacturer, dev.Productname))
}

func appendRuntimeFlags(cols []string, dev fritz.Device) []string {
	return append(cols,
		console.IntToCheckmark(dev.Present),
		console.StringToCheckmark(dev.Thermostat.Lock)+"/"+console.StringToCheckmark(dev.Thermostat.DeviceLock))
}

func appendRuntimeWarnings(cols []string, dev fritz.Device) []string {
	return append(cols, errorCode(dev.Thermostat.ErrorCode),
		console.Stoc(dev.Thermostat.BatteryLow).Inverse().String())
}

func appendTemperatureValues(cols []string, dev fritz.Device) []string {
	return append(cols,
		fmtUnit(dev.Thermostat.FmtMeasuredTemperature, "°C"),
		fmtUnit(dev.Temperature.FmtOffset, "°C"),
		fmtUnit(dev.Thermostat.FmtGoalTemperature, "°C"),
		fmtUnit(dev.Thermostat.FmtSavingTemperature, "°C"),
		fmtUnit(dev.Thermostat.FmtComfortTemperature, "°C"),
		fmtNextChange(dev.Thermostat.NextChange))
}
func fmtNextChange(n fritz.NextChange) string {
	return stringutils.DefaultIfEmpty(n.FmtTimestamp(time.Now()), "?") +
		" -> " +
		stringutils.DefaultIfEmpty(n.FmtGoalTemperature(), "?") +
		"°C"
}

func errorCode(ec string) string {
	checkMark := console.Stoc(ec).Inverse()
	return checkMark.String() + errorCodesVsDescriptions[ec]
}
