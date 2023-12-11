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

type ShapeScore int

var (
	RockScore     ShapeScore = 1
	PaperScore    ShapeScore = 2
	ScissorsScore ShapeScore = 3
)

type RoundOutcome int

var (
	WinOutcome  RoundOutcome = 6
	DrawOutcome RoundOutcome = 3
	LoseOutcome RoundOutcome = 0
)

type Shapes map[string]ShapeScore

func One(input string) string {
	input = strings.Trim(input, "\n")

	opponentShapes := Shapes{
		"A": RockScore,
		"B": PaperScore,
		"C": ScissorsScore,
	}

	responseShapes := Shapes{
		"X": RockScore,
		"Y": PaperScore,
		"Z": ScissorsScore,
	}

	i := strings.Split(input, "\n")
	score := 0
	for _, round := range i {
		s := strings.Split(round, " ")
		opponentShape := opponentShapes[s[0]]
		responseShape := responseShapes[s[1]]

		score += int(responseShape)

		switch {
		case opponentShape == responseShape:
			score += int(DrawOutcome)
		case opponentShape == RockScore && responseShape == ScissorsScore:
			score += int(LoseOutcome)
		case opponentShape == ScissorsScore && responseShape == PaperScore:
			score += int(LoseOutcome)
		case opponentShape == PaperScore && responseShape == RockScore:
			score += int(LoseOutcome)
		default:
			score += int(WinOutcome)
		}
	}

	return strconv.Itoa(score)
}

func Two(input string) string {
	input = strings.Trim(input, "\n")

	opponentShapes := Shapes{
		"A": RockScore,
		"B": PaperScore,
		"C": ScissorsScore,
	}

	responseOutcomes := map[string]RoundOutcome{
		"X": LoseOutcome,
		"Y": DrawOutcome,
		"Z": WinOutcome,
	}

	i := strings.Split(input, "\n")
	score := 0

	for _, round := range i {
		if round == "" {
			continue
		}

		var responseShape ShapeScore

		s := strings.Split(round, " ")
		opponentShape := opponentShapes[s[0]]
		responseOutcome := responseOutcomes[s[1]]

		if responseOutcome == WinOutcome {
			switch {
			case opponentShape == RockScore:
				responseShape = PaperScore
			case opponentShape == ScissorsScore:
				responseShape = RockScore
			case opponentShape == PaperScore:
				responseShape = ScissorsScore
			}
		}

		if responseOutcome == LoseOutcome {
			switch {
			case opponentShape == RockScore:
				responseShape = ScissorsScore
			case opponentShape == ScissorsScore:
				responseShape = PaperScore
			case opponentShape == PaperScore:
				responseShape = RockScore
			}
		}

		if responseOutcome == DrawOutcome {
			responseShape = opponentShape
		}

		score = score + int(responseOutcome) + int(responseShape)
	}

	return strconv.Itoa(score)
}
