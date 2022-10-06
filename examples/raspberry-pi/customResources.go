package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/abu-lang/goabu/memory"
	"github.com/abu-lang/goabu/physical"
	"gobot.io/x/gobot/drivers/gpio"
)

type customResources struct {
	*physical.IOresources
}

func (c *customResources) Add(t string, name string, args ...interface{}) error {
	err := c.IOresources.Add(t, name, args...)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
	}
	return err
}

func MakeCustomResources(a physical.IOadaptor) *customResources {
	res := &customResources{physical.MakeEmptyIOresources(a)}
	res.AddOutputFrame("Led", MakeLed)
	res.AddOutputFrame("L293Motor", MakeL293Motor)
	res.AddInputFrame("Button", MakeButton)
	return res
}

type Led struct {
	name string
}

func MakeLed(adaptor physical.IOadaptor, name string, args ...interface{}) (physical.IOdelegate, memory.Resources, error) {
	if len(args) != 2 {
		return nil, memory.MakeResources(), errors.New("led constructor invocation should have 4 arguments")
	}
	pin, ok := args[0].(int)
	if !ok {
		return nil, memory.MakeResources(), errors.New("third argument of led constructor should be an int specifying a pin")
	}
	active, ok := args[1].(bool)
	if !ok || active {
		return nil, memory.MakeResources(), errors.New("fourth argument of led constructor should be false")
	}
	resources := memory.MakeResources()
	resources.Bool["active"] = active
	resources.Integer["pin"] = int64(pin)
	return Led{name: name}, resources, nil
}

func (l Led) Start(adaptor physical.IOadaptor, inputs chan<- string, errors chan<- error) error {
	return nil
}

func (l Led) Modified(adaptor physical.IOadaptor, name string, resources memory.Resources, errs chan<- error) *memory.Resources {
	pin := l.name + "_pin"
	if name == pin {
		errs <- errors.New("led pins cannot be modified after initialization")
		return nil
	}
	var out byte = 0
	if resources.Bool[l.name+"_active"] {
		out = 1
	}
	err := adaptor.DigitalWrite(strconv.FormatInt(resources.Integer[pin], 10), out)
	if err != nil {
		errs <- err
	}
	return nil
}

type Button struct {
	name   string
	driver *gpio.ButtonDriver
}

func MakeButton(adaptor physical.IOadaptor, name string, args ...interface{}) (physical.IOdelegate, memory.Resources, error) {
	if len(args) != 1 {
		return nil, memory.MakeResources(), errors.New("button constructor invocation should have 3 arguments")
	}
	pin, ok := args[0].(int)
	if !ok {
		return nil, memory.MakeResources(), errors.New("third argument of button constructor should be an int specifying a pin")
	}
	resources := memory.MakeResources()
	resources.Bool["pressed"] = false
	resources.Integer["pin"] = int64(pin)
	return Button{name: name, driver: gpio.NewButtonDriver(adaptor, strconv.Itoa(pin))}, resources, nil
}

func (b Button) Start(adaptor physical.IOadaptor, inputs chan<- string, errors chan<- error) error {
	err := b.driver.Start()
	if err != nil {
		return err
	}
	go b.getButtonInput(inputs, errors)
	return nil
}

func (b Button) Modified(adaptor physical.IOadaptor, name string, resources memory.Resources, errs chan<- error) *memory.Resources {
	if name == b.name+"_pin" {
		errs <- errors.New("button pins cannot be modified after initialization")
	}
	return nil
}

func (b Button) getButtonInput(in chan<- string, errs chan<- error) {
	events := b.driver.Subscribe()
	status := false
	push := b.name + "_pressed = true;"
	release := b.name + "_pressed = false;"
	event := <-events
	for {
		var inputs chan<- string = nil
		var action string
		switch event.Name {
		case gpio.ButtonPush:
			action = push
			if !status {
				inputs = in
			}
		case gpio.ButtonRelease:
			action = release
			if status {
				inputs = in
			}
		case gpio.Error:
			errs <- fmt.Errorf("input error on button %s, received: %v", b.name, event.Data)
		}
		select {
		case inputs <- action:
			status = !status
		case event = <-events:
		}
	}
}

type L293Motor struct {
	name string
}

func MakeL293Motor(adaptor physical.IOadaptor, name string, args ...interface{}) (physical.IOdelegate, memory.Resources, error) {
	if len(args) != 2 {
		return nil, memory.MakeResources(), errors.New("l293motor constructor invocation should have 4 arguments")
	}
	forward, ok := args[0].(int)
	if !ok {
		return nil, memory.MakeResources(), errors.New("third argument of l293motor constructor should be an int specifying a pin")
	}
	backward, ok := args[1].(int)
	if !ok {
		return nil, memory.MakeResources(), errors.New("fourth argument of l293motor constructor should be an int specifying a pin")
	}
	resources := memory.MakeResources()
	resources.Integer["fPin"] = int64(forward)
	resources.Integer["bPin"] = int64(backward)
	resources.Integer["fSpeed"] = 0
	resources.Integer["bSpeed"] = 0
	return L293Motor{name: name}, resources, nil
}

func (m L293Motor) Start(adaptor physical.IOadaptor, inputs chan<- string, errors chan<- error) error {
	return nil
}

func (m L293Motor) Modified(adaptor physical.IOadaptor, name string, resources memory.Resources, errs chan<- error) *memory.Resources {
	if name == m.name+"_fPin" || name == m.name+"_bPin" {
		errs <- errors.New("l293motor pins cannot be modified after initialization")
		return nil
	}
	normalize := func(arg int64) byte {
		if arg < 0 {
			return 0
		}
		if arg > 255 {
			return 255
		}
		return byte(arg)
	}
	var ops [2]struct {
		pin   string
		speed byte
	}
	ops[0].pin = strconv.FormatInt(resources.Integer[m.name+"_fPin"], 10)
	ops[0].speed = normalize(resources.Integer[m.name+"_fSpeed"])
	ops[1].pin = strconv.FormatInt(resources.Integer[m.name+"_bPin"], 10)
	ops[1].speed = normalize(resources.Integer[m.name+"_bSpeed"])
	if ops[0].speed > ops[1].speed {
		ops[0], ops[1] = ops[1], ops[0]
	}
	for i := 0; i < 2; i++ {
		err := adaptor.PwmWrite(ops[i].pin, ops[i].speed)
		if err != nil {
			errs <- err
			return nil
		}
	}
	return nil
}
