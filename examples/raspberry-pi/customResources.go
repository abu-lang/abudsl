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
	if !ok {
		return nil, memory.MakeResources(), errors.New("fourth argument of led constructor should be a boolean")
	}
	resources := memory.MakeResources()
	resources.Bool["active"] = active
	resources.Integer["pin"] = int64(pin)
	var out byte = 0
	if active {
		out = 1
	}
	err := adaptor.DigitalWrite(strconv.Itoa(pin), out)
	if err != nil {
		return nil, memory.MakeResources(), err
	}
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
	name         string
	startPressed bool
	driver       *gpio.ButtonDriver
}

func MakeButton(adaptor physical.IOadaptor, name string, args ...interface{}) (physical.IOdelegate, memory.Resources, error) {
	if len(args) != 2 {
		return nil, memory.MakeResources(), errors.New("button constructor invocation should have 4 arguments")
	}
	pin, ok := args[0].(int)
	if !ok {
		return nil, memory.MakeResources(), errors.New("third argument of button constructor should be an int specifying a pin")
	}
	pressed, ok := args[1].(bool)
	if !ok {
		return nil, memory.MakeResources(), errors.New("fourth argument of button constructor should be a boolean")
	}
	resources := memory.MakeResources()
	resources.Bool["pressed"] = pressed
	resources.Integer["pin"] = int64(pin)
	return Button{name: name,
		startPressed: pressed,
		driver:       gpio.NewButtonDriver(adaptor, strconv.Itoa(pin)),
	}, resources, nil
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
	status := b.startPressed
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

func (m L293Motor) setSpeed(adaptor physical.IOadaptor, fPin, bPin string, fSpeed, bSpeed int64) error {
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
	ops[0].pin = fPin
	ops[0].speed = normalize(fSpeed)
	ops[1].pin = bPin
	ops[1].speed = normalize(bSpeed)
	if ops[0].speed > ops[1].speed {
		ops[0], ops[1] = ops[1], ops[0]
	}
	for i := 0; i < 2; i++ {
		err := adaptor.PwmWrite(ops[i].pin, ops[i].speed)
		if err != nil {
			return err
		}
	}
	return nil
}

func MakeL293Motor(adaptor physical.IOadaptor, name string, args ...interface{}) (physical.IOdelegate, memory.Resources, error) {
	if len(args) != 4 {
		return nil, memory.MakeResources(), errors.New("l293motor constructor invocation should have 6 arguments")
	}
	forward, ok := args[0].(int)
	if !ok {
		return nil, memory.MakeResources(), errors.New("third argument of l293motor constructor should be an int specifying a pin")
	}
	backward, ok := args[1].(int)
	if !ok {
		return nil, memory.MakeResources(), errors.New("fourth argument of l293motor constructor should be an int specifying a pin")
	}
	fSpeed, ok := args[2].(int)
	if !ok {
		return nil, memory.MakeResources(), errors.New("fifth argument of l293motor constructor should be an int")
	}
	bSpeed, ok := args[3].(int)
	if !ok {
		return nil, memory.MakeResources(), errors.New("sixth argument of l293motor constructor should be an int")
	}
	resources := memory.MakeResources()
	resources.Integer["fPin"] = int64(forward)
	resources.Integer["bPin"] = int64(backward)
	resources.Integer["fSpeed"] = int64(fSpeed)
	resources.Integer["bSpeed"] = int64(bSpeed)
	res := L293Motor{name: name}
	err := res.setSpeed(adaptor, strconv.Itoa(forward), strconv.Itoa(backward), int64(fSpeed), int64(bSpeed))
	if err != nil {
		return nil, memory.MakeResources(), err
	}
	return res, resources, nil
}

func (m L293Motor) Start(adaptor physical.IOadaptor, inputs chan<- string, errors chan<- error) error {
	return nil
}

func (m L293Motor) Modified(adaptor physical.IOadaptor, name string, resources memory.Resources, errs chan<- error) *memory.Resources {
	if name == m.name+"_fPin" || name == m.name+"_bPin" {
		errs <- errors.New("l293motor pins cannot be modified after initialization")
		return nil
	}
	err := m.setSpeed(adaptor,
		strconv.FormatInt(resources.Integer[m.name+"_fPin"], 10),
		strconv.FormatInt(resources.Integer[m.name+"_bPin"], 10),
		resources.Integer[m.name+"_fSpeed"],
		resources.Integer[m.name+"_bSpeed"],
	)
	if err != nil {
		errs <- err
	}
	return nil
}
