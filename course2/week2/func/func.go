package main

import (
	"fmt"
	"math"
)

func GenDisplaceFn(a, v0, s0 float64) func(float64) float64 {
	return func(t float64) float64 {
		return (1/2.0)*a*math.Pow(t, 2) + v0*t + s0
	}
}

func main() {

	var acc, iniVel, iniDis, time float64
	fmt.Print("Enter the acceleration, inivial velocity and initial displacement: ")
	fmt.Scanf("%f %f %f", &acc, &iniVel, &iniDis)
	fn := GenDisplaceFn(acc, iniVel, iniDis)
	fmt.Print("Enter the time value: ")
	fmt.Scanf("%f", &time)
	fmt.Println(fn(time))

}
