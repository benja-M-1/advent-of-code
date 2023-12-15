package main

import (
	"embed"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

//go:embed *.txt
var f embed.FS

func main() {
	input, _ := f.ReadFile("input.txt")

	r1 := One(string(input))
	fmt.Printf("puzzle 1: %v\n", r1)
	r2 := Two(string(input))
	fmt.Printf("puzzle 2: %v\n", r2)
}

func One(input string) int {
	sum := 0
	for _, step := range strings.Split(strings.Trim(input, "\n"), ",") {
		sum += hash(step)
	}

	return sum
}

func hash(step string) int {
	var current int
	for _, c := range []byte(step) {
		current = ((current + int(c)) * 17) % 256
	}
	return current
}

type Lens struct {
	Label       string
	FocalLength int
}

type Box struct {
	ID     int
	Lenses []Lens
}

func (b *Box) Slot(lens Lens) int {
	return slices.IndexFunc(b.Lenses, func(l Lens) bool {
		return l.Label == lens.Label
	})
}

func (b *Box) Remove(lens Lens) {
	slot := b.Slot(lens)
	if slot == -1 {
		return
	}

	b.Lenses = append(b.Lenses[:slot], b.Lenses[slot+1:]...)
}

func (b *Box) Add(lens Lens) {
	slot := b.Slot(lens)

	// Replace the focal length of the existing lens in the box
	if slot != -1 {
		b.Lenses[slot].FocalLength = lens.FocalLength
		return
	}

	// Add the lens to the box
	b.Lenses = append(b.Lenses, lens)
}

func Two(input string) int {
	boxes := make([]*Box, 256)

	for _, step := range strings.Split(strings.Trim(input, "\n"), ",") {
		var (
			operation string
			lens      Lens
		)

		if i := strings.Index(step, "-"); i != -1 {
			operation = "-"
			lens = Lens{Label: step[:i]}
		}

		if i := strings.Index(step, "="); i != -1 {
			operation = "="
			focalLength, _ := strconv.Atoi(step[i+1:])
			lens = Lens{Label: step[:i], FocalLength: focalLength}
		}

		id := hash(lens.Label)
		box := boxes[id]
		if box == nil {
			box = &Box{ID: id}
			boxes[id] = box
		}

		switch operation {
		case "-":
			box.Remove(lens)
		case "=":
			box.Add(lens)
		}
	}

	focusingPower := 0
	for _, box := range boxes {
		if box != nil {
			for slot, lens := range box.Lenses {
				focusingPower += (1 + box.ID) * (slot + 1) * lens.FocalLength
			}
		}
	}

	return focusingPower
}
