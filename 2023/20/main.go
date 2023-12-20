package main

import (
	"embed"
	"fmt"
	"strings"

	aocmath "adventofcode/pkg/math"
	aocslices "adventofcode/pkg/slices"

	"golang.org/x/exp/maps"
	"golang.org/x/exp/slices"
)

//go:embed *.txt
var f embed.FS

func main() {
	input, _ := f.ReadFile("input.txt")

	r1 := One(string(input), 1000)
	fmt.Printf("puzzle 1: %v\n", r1)
	r2 := Two(string(input))
	fmt.Printf("puzzle 2: %v\n", r2)
}

type Pulse struct {
	Pulse    string
	Sender   string
	Receiver string
}

type Module interface {
	GetName() string
	HasDestination(string) bool
	Process(pulse Pulse)
	Reset()
	Kind() string
}

type Button struct {
	Queue *PulseQueue
}

func (b Button) Push() {
	b.Queue.Enqueue(Pulse{"low", "button", "broadcaster"})
}

type Broadcaster struct {
	Name         string
	Destinations []string
	Queue        *PulseQueue
}

func (b *Broadcaster) GetName() string              { return b.Name }
func (b *Broadcaster) HasDestination(s string) bool { return slices.Contains(b.Destinations, s) }
func (b *Broadcaster) Reset()                       { return }
func (b *Broadcaster) Kind() string                 { return "broadcaster" }

// Process processes the received pulse.
// It transfers the same pulse to destinations
func (b *Broadcaster) Process(pulse Pulse) {
	for _, d := range b.Destinations {
		b.Queue.Enqueue(Pulse{pulse.Pulse, b.Name, d})
	}
}

type FlipFlop struct {
	Name         string
	Destinations []string
	Queue        *PulseQueue
	On           bool
}

func (f *FlipFlop) GetName() string              { return f.Name }
func (f *FlipFlop) HasDestination(s string) bool { return slices.Contains(f.Destinations, s) }
func (f *FlipFlop) Reset()                       { f.On = false }
func (f *FlipFlop) Kind() string                 { return "flip-flop" }

// Process processes the received pulse.
// off, receives low pulse -> on, send high
// on, receives low pulse -> off, send high
// receives high pulse -> do nothing
func (f *FlipFlop) Process(pulse Pulse) {
	if pulse.Pulse == "high" {
		return
	}

	p := "low"
	if !f.On {
		p = "high"
	}
	f.On = !f.On

	for _, d := range f.Destinations {
		f.Queue.Enqueue(Pulse{p, f.Name, d})
	}
}

type Conjunction struct {
	Name         string
	Destinations []string
	Queue        *PulseQueue
	Memory       map[string]string
	Modules      map[string]Module
}

func (c *Conjunction) GetName() string              { return c.Name }
func (c *Conjunction) HasDestination(s string) bool { return slices.Contains(c.Destinations, s) }
func (c *Conjunction) Reset() {
	for _, m := range c.inputs() {
		c.Memory[m.GetName()] = "low"
	}
}
func (c *Conjunction) Kind() string { return "conjunction" }

// Process processes the received pulse.
// remember previous received pulse for each connected module
// default is a low pulse
// receives a low pulse -> updates memory and send low pulse if every module has received a high pulse
// receives a high pulse -> updates memory and send high pulse if every module has received a low pulse
func (c *Conjunction) Process(pulse Pulse) {
	if c.Memory == nil {
		c.Memory = map[string]string{}
		for _, m := range c.inputs() {
			c.Memory[m.GetName()] = "low"
		}
	}

	c.Memory[pulse.Sender] = pulse.Pulse

	p := "high"
	hp := aocslices.CountIf(maps.Values(c.Memory), func(v string) bool { return v == "high" })
	if hp == len(c.Memory) {
		p = "low"
	}

	for _, d := range c.Destinations {
		c.Queue.Enqueue(Pulse{
			Pulse:    p,
			Sender:   c.Name,
			Receiver: d,
		})
	}
}

func (c *Conjunction) inputs() []Module {
	inputs := []Module{}
	for _, m := range c.Modules {
		if m.HasDestination(c.Name) {
			inputs = append(inputs, m)
		}
	}

	return inputs
}

type PulseQueue struct {
	q []Pulse
}

func (p *PulseQueue) Enqueue(pulse Pulse) {
	p.q = append(p.q, pulse)
}

func (p *PulseQueue) Dequeue() Pulse {
	pulse := p.q[0]
	p.q = p.q[1:]

	return pulse
}

func (p *PulseQueue) IsEmpty() bool {
	if len(p.q) == 0 {
		return true
	}

	return false
}

func One(input string, push int) int {
	input = strings.Trim(input, "\n")

	queue := &PulseQueue{}
	modules := readModuleConfiguration(input, queue)

	history := map[string]int{}
	button := Button{Queue: queue}
	for i := 0; i < push; i++ {
		button.Push()
		for !queue.IsEmpty() {
			pulse := queue.Dequeue()
			history[pulse.Pulse]++
			if m, ok := modules[pulse.Receiver]; ok {
				m.Process(pulse)
			}
		}
	}

	return history["high"] * history["low"]
}

func readModuleConfiguration(input string, queue *PulseQueue) map[string]Module {
	modules := map[string]Module{}

	for _, line := range strings.Split(input, "\n") {
		parts := strings.Fields(line)

		if parts[0] == "broadcaster" {
			m := &Broadcaster{Queue: queue}
			m.Name = "broadcaster"
			for _, d := range parts[2:] {
				m.Destinations = append(m.Destinations, strings.Trim(d, ","))
			}
			modules[m.Name] = m
		} else if string(parts[0][0]) == "%" {
			m := &FlipFlop{Queue: queue}
			m.Name = parts[0][1:]
			for _, d := range parts[2:] {
				m.Destinations = append(m.Destinations, strings.Trim(d, ","))
			}
			modules[m.Name] = m
		} else {
			m := &Conjunction{Queue: queue, Modules: modules}
			m.Name = parts[0][1:]
			for _, d := range parts[2:] {
				m.Destinations = append(m.Destinations, strings.Trim(d, ","))
			}
			modules[m.Name] = m
		}
	}

	return modules
}

func Two(input string) int {
	input = strings.Trim(input, "\n")

	queue := &PulseQueue{}
	modules := readModuleConfiguration(input, queue)

	button := Button{Queue: queue}

	cycles := map[string]int{
		"pr": 0, "bt": 0, "fv": 0, "rd": 0,
	}
	push := 0
loop:
	for {
		push++
		button.Push()
		for !queue.IsEmpty() {
			pulse := queue.Dequeue()
			if m, ok := modules[pulse.Receiver]; ok {

				if pulse.Pulse == "low" && slices.Contains(maps.Keys(cycles), pulse.Receiver) {
					cycles[m.GetName()] = push

					if aocslices.CountIf(maps.Values(cycles), func(v int) bool { return v == 0 }) == 0 {
						break loop
					}
				}

				m.Process(pulse)
			}
		}
	}

	return aocmath.LeastCommonMultiplied(cycles["pr"], cycles["bt"], cycles["fv"], cycles["rd"])
}
