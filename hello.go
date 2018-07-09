package main

import (
	"fmt"
	"math"
)

// package main

var a int = 12
var b int = 12

func main() {

	c := a + b
	d := math.Abs(float64(c))
	fmt.Println(d)

}
