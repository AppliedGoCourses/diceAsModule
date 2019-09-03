package diceprinter

import (
	"fmt"

	"github.com/appliedgocourses/dice"
	"github.com/common-nighthawk/go-figure"
)

func Roll(sides int) {
	fmt.Printf("Rolling a %d-sided die: %d\n", sides, dice.Roll(sides))
}

func Pretty(sides int) {
	out := fmt.Sprintf("%d-sided roll: %d", sides, dice.Roll(sides))
	f := figure.NewFigure(out, "", true)
	f.Print()
}
