package variables

import (
	"fmt"
	"time"
)

var Name string
var Status bool
var Amount float32
var Time time.Time

func ShowStrings() {
	Name = "Juan"
	Status = true
	Amount = 10.5
	Time = time.Now()
	fmt.Println("Name=", Name, "Status=", Status, "Amount=", Amount, "Time=", Time)
}
