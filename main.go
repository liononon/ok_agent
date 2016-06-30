package main

import (
	//go builtin pkg
	"flag"

	//local pkg
	"github.com/OpsKitchen/ok_agent/util"
	"github.com/Sirupsen/logrus"
)

func main() {
	var baseConfigFile *string
	var debugMode *bool
	var dispatcher *Dispatcher

	//parse config file from cli argument
	baseConfigFile = flag.String("c", "/etc/ok_agent.json", "base config file path")
	debugMode = flag.Bool("d", false, "enable debug log")
	flag.Parse()

	//enable debug log
	if *debugMode {
		util.Logger.Level = logrus.DebugLevel
	}

	//dispatch
	dispatcher = &Dispatcher{
		BaseConfigFile: *baseConfigFile,
	}
	dispatcher.Dispatch()
}
