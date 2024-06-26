package main

import (
	"fmt"

	configurator "github.com/OpenCHAMI/configurator/internal"
	"github.com/OpenCHAMI/configurator/internal/util"
)

type Syslog struct{}

func (g *Syslog) GetName() string {
	return "syslog"
}

func (g *Syslog) Generate(config *configurator.Config, opts ...util.Option) (map[string][]byte, error) {
	return nil, fmt.Errorf("plugin does not implement generation function")
}

var Generator Syslog
