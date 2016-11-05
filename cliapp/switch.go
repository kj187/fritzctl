package cliapp

import (
	"github.com/bpicode/fritzctl/fatals"
	"github.com/bpicode/fritzctl/fritz"
	"github.com/bpicode/fritzctl/logger"
	"github.com/mitchellh/cli"
)

type switchOnCommand struct {
}

func (cmd *switchOnCommand) Help() string {
	return "Switch on device. Example usage: fritzctl switch on mydevice."
}

func (cmd *switchOnCommand) Synopsis() string {
	return "Switch on device"
}

func (cmd *switchOnCommand) Run(args []string) int {
	fatals.AssertStringSliceHasAtLeast(args, 1, "Insufficient input: device name expected.")
	f := fritz.UsingClient(clientLogin())
	res, err := f.Switch(args[0], "on")
	fatals.AssertNoError(err, "Unable to switch on device:", err)
	logger.Info("Success! FRITZ!Box answered: " + res)
	return 0
}

func switchOnDevice() (cli.Command, error) {
	p := switchOnCommand{}
	return &p, nil
}

type switchOffCommand struct {
}

func (cmd *switchOffCommand) Help() string {
	return "Switch off device. Example usage: fritzctl switch on mydevice."
}

func (cmd *switchOffCommand) Synopsis() string {
	return "Switch off device"
}

func (cmd *switchOffCommand) Run(args []string) int {
	fatals.AssertStringSliceHasAtLeast(args, 1, "Insufficient input: device name expected.")
	f := fritz.UsingClient(clientLogin())
	res, err := f.Switch(args[0], "off")
	fatals.AssertNoError(err, "Unable to switch off device:", err)
	logger.Info("Success! FRITZ!Box answered: " + res)
	return 0
}

func switchOffDevice() (cli.Command, error) {
	p := switchOffCommand{}
	return &p, nil
}