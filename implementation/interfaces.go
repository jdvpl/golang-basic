package implementation

import (
	"fmt"

	"go.mod/interfaces"
)

func HumanBreathing(hu interfaces.Human, name string) {
	hu.Breathe()
	hu.AddingName(name)
	fmt.Printf("%s, I am %s, I am Alive= %t,  \n", hu.SayingHello(), hu.Sex(), hu.BeingAlive())
}
