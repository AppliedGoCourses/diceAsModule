package diceprinter

import (
	"fmt"

	"github.com/appliedgocourses/dice"
)

func Roll(sides int) {
	fmt.Printf("Rolling a %d-sided die: %d\n", sides, dice.Roll(sides))
}
