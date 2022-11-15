package main

import (
	"fmt"
	"math"
)

// function that count the area and circumference of a sphere
func sphereStuff(r int) (d int, v float32, p float32) {
	d = r * 2
	v = 4.0 / 3.0 * math.Pi * float32(r^3.0)
	p = 4.0 * math.Pi * float32(r^2.0)

	return
}

func main() {
	sphere_d, sphere_v, sphere_p := sphereStuff(10)
	fmt.Printf("diameter: %v, circumference: %v, area: %v\n", sphere_d, sphere_v, sphere_p)
}
