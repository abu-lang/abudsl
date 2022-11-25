package main

import (
	"fmt"
	"os"

	"github.com/abu-lang/goabu/memory"
)

type coordsResources struct {
	memory.Resources
}

func MakeCoordsResources() coordsResources {
	return coordsResources{memory.MakeResources()}
}

func (c coordsResources) Add(t string, name string, lat, lon int64) error {
	if t != "coords" {
		err := fmt.Errorf("unknown type %s", t)
		fmt.Fprintln(os.Stderr, err.Error())
		return err
	}
	c.Integer[name+"_lat"] = lat
	c.Integer[name+"_lon"] = lon
	return nil
}

func (c coordsResources) Copy() memory.ResourceController {
	return coordsResources{
		c.Resources.Copy().GetResources(),
	}
}
