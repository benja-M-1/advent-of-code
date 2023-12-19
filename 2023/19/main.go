package main

import (
	"embed"
	"fmt"
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

type rule struct {
	Category  string
	Condition string
	Value     int
	Next      string
}

type workflow struct {
	Name  string
	Rules []rule
}

type workflows map[string]workflow

type part struct {
	Ratings map[string]int
}

func (p part) Total() int {
	total := 0
	for _, value := range p.Ratings {
		total += value
	}

	return total
}

type interval struct {
	Min int
	Max int
}

func One(input string) int {
	input = strings.Trim(input, "\n")

	lists := strings.Split(input, "\n\n")
	ws := parseWorkflows(lists[0])
	parts := parseParts(lists[1])

	sum := 0
	for _, p := range parts {
		r := run(p, ws)
		if r != "A" {
			continue
		}

		sum += p.Total()
	}

	return sum
}

func parseWorkflows(input string) workflows {
	ws := workflows{}
	for _, line := range strings.Split(input, "\n") {
		w := workflow{}
		start := strings.Index(line, "{")
		w.Name = line[:start]

		for _, s := range strings.Split(line[start+1:len(line)-1], ",") {
			r := rule{}

			if c := strings.ContainsFunc(s, func(r rune) bool { return string(r) == "<" || string(r) == ">" }); !c {
				r.Next = s
				w.Rules = append(w.Rules, r)
				continue
			}

			i := strings.Index(s, ":")

			r.Category = s[:1]
			r.Condition = s[1:2]
			r.Value, _ = strconv.Atoi(s[2:i])
			r.Next = s[i+1:]

			w.Rules = append(w.Rules, r)
		}

		ws[w.Name] = w
	}

	return ws
}

func parseParts(input string) []part {
	parts := []part{}
	for _, line := range strings.Split(input, "\n") {
		p := part{Ratings: make(map[string]int)}
		for _, s := range strings.Split(line[1:len(line)-1], ",") {
			p.Ratings[s[:1]], _ = strconv.Atoi(s[2:])
		}

		parts = append(parts, p)
	}

	return parts
}

func run(p part, ws workflows) string {
	next := ws["in"]
	for {
		current := next
	CurrentLoop:
		for _, r := range current.Rules {
			if r.Condition == "" {
				if len(r.Next) == 1 {
					return r.Next

				}
				next = ws[r.Next]
				continue
			}

			for k, v := range p.Ratings {
				if k == r.Category {
					var result string
					switch r.Condition {
					case "<":
						if v < r.Value {
							result = r.Next
						}
					case ">":
						if v > r.Value {
							result = r.Next
						}
					}
					if len(result) > 1 {
						next = ws[result]
						break CurrentLoop
					} else if len(result) == 1 {
						return result
					}
				}
			}
		}
	}

	return "R"
}

func Two(input string) int {
	input = strings.Trim(input, "\n")

	lists := strings.Split(input, "\n\n")
	ws := parseWorkflows(lists[0])

	intervals := map[string]*interval{
		"x": &interval{Min: 1, Max: 4000},
		"m": &interval{Min: 1, Max: 4000},
		"a": &interval{Min: 1, Max: 4000},
		"s": &interval{Min: 1, Max: 4000},
	}

	return combinations(ws, "in", intervals)
}

func combinations(ws workflows, name string, intervals map[string]*interval) int {
	sum := 0
	restIntervals := map[string]*interval{}
	for k, v := range intervals {
		restIntervals[k] = &interval{Min: v.Min, Max: v.Max}
	}

	for _, r := range ws[name].Rules {
		currentIntervals := map[string]*interval{}
		for k, v := range restIntervals {
			currentIntervals[k] = &interval{Min: v.Min, Max: v.Max}
		}

		if r.Condition != "" {
			i := restIntervals[r.Category]
			switch r.Condition {
			case "<":
				if r.Value < i.Max {
					currentIntervals[r.Category].Max = r.Value - 1
					restIntervals[r.Category].Min = r.Value
				}
			case ">":
				if r.Value > i.Min {
					currentIntervals[r.Category].Min = r.Value + 1
					restIntervals[r.Category].Max = r.Value
				}
			}
		}

		if r.Next == "R" {
			continue
		}

		if r.Next == "A" {
			sum += calculate(currentIntervals)
		} else {
			sum += combinations(ws, r.Next, currentIntervals)
		}
	}

	return sum
}

func calculate(intervals map[string]*interval) int {
	total := 1
	for _, v := range intervals {
		total *= (v.Max - v.Min) + 1
	}

	return total
}
