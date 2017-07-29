package main

import (
	"fmt"
	"sort"
	"strconv"
)

//import "os"

/**
 * Auto-generated code below aims at helping you parse
 * the standard input according to the problem statement.
 **/
/*
* Premièrement trouver la surface plane unique sur laquelle on peut atterrir
* Ensuite déplacer l'appareil au-dessus
* Enfin descendre l'appareil à la surface
* Il faudra certainement optimiser ensuite
* Il faut faire attention à l'altitude et la hauteur maximum du sol que l'on pourra rencontrer : où sont les pics ?
 */

const (
	tiltAngleIncrement = 15
)

type Point struct {
	X, Y int
}

type ByX []Point

func (ps ByX) Len() int {
	return len(ps)
}
func (ps ByX) Less(i, j int) bool {
	return ps[i].X < ps[j].X
}
func (ps ByX) Swap(i, j int) {
	ps[i], ps[j] = ps[j], ps[i]
}

type ByYDescending []Point

func (ps ByYDescending) Len() int {
	return len(ps)
}
func (ps ByYDescending) Less(i, j int) bool {
	return ps[i].Y > ps[j].Y
}
func (ps ByYDescending) Swap(i, j int) {
	ps[i], ps[j] = ps[j], ps[i]
}

type Plane struct {
	StartX, EndX, Y int
}

type FlightParameter struct {
	X, Y, HSpeed, VSpeed, Fuel, Rotate, Power int
}

func main() {
	// surfaceN: the number of points used to draw the surface of Mars.
	var surfaceN int
	fmt.Scan(&surfaceN)
	points := make([]Point, surfaceN, surfaceN)
	pointsByY := make(map[int][]Point)
	for i := 0; i < surfaceN; i++ {
		// landX: X coordinate of a surface point. (0 to 6999)
		// landY: Y coordinate of a surface point. By linking all the points together in a sequential fashion, you form the surface of Mars.
		var landX, landY int
		fmt.Scan(&landX, &landY)
		point := Point{landX, landY}
		points = append(points, point)
		if p, ok := pointsByY[landY]; ok {
			pointsByY[landY] = append(p, point)
		} else {
			pointsByY[landY] = []Point{point}
		}
	}

	var flatGround Plane
	for k, v := range pointsByY {
		if length := len(v); length > 1 {
			sort.Sort(ByX(v))
			flatGround = Plane{v[0].X, v[length-1].X, k}
		}
	}

	loopCount := 0
	flightParameters := make(map[int]FlightParameter)
	for {
		// hSpeed: the horizontal speed (in m/s), can be negative.
		// vSpeed: the vertical speed (in m/s), can be negative.
		// fuel: the quantity of remaining fuel in liters.
		// rotate: the rotation angle in degrees (-90 to 90).
		// power: the thrust power (0 to 4).
		var X, Y, hSpeed, vSpeed, fuel, rotate, power int
		fmt.Scan(&X, &Y, &hSpeed, &vSpeed, &fuel, &rotate, &power)
		flightParameters[loopCount] = FlightParameter{X, Y, hSpeed, vSpeed, fuel, rotate, power}
		// fmt.Fprintln(os.Stderr, "Debug messages...")

		// rotate power. rotate is the desired rotation angle. power is the desired thrust power.
		newRotate, newPower := ComputeRotateAndPower(flightParameters[loopCount], flatGround, loopCount)
		loopCount++
		fmt.Printf("%s %s\n", newRotate, newPower)
	}
}

func ComputeRotateAndPower(parameter FlightParameter, flatGround Plane, loopCount int) (string, string) {
	var rotationToApply, powerToApply string
	rotationToApply = "0"
	powerToApply = "0"
	//Where am I ?
	if parameter.X <= flatGround.StartX {
		//Lander on the left side, we need to increase X
		//We need to go right by decreasing the tilt angle and applying power
		//Est-ce qu'entre parameter.X et flatGround.StartX il y a des pics élevés ?
		//Si des pics plus haut que parameter.X
	} else if parameter.X > flatGround.StartX && parameter.X < flatGround.EndX {
		//Lander above the flat ground
		//We need to nullify hSpeed and to nullify tilt angle
		//Then we need to limit the vSpeed
		if parameter.HSpeed < 20 || parameter.HSpeed > -20 {
			currentVSpeed := parameter.VSpeed
			currentPower := parameter.Power
			currentRotate := parameter.Rotate
			switch {
			case currentVSpeed > 39:
				powerToApply = strconv.Itoa(max(currentPower-1, 0))
			case currentVSpeed < -39:
				powerToApply = strconv.Itoa(min(currentPower+1, 4))
			default:
				powerToApply = "0"
			}
			switch {
			//On peut avoir une rotation initiale et ne pas être sur un multiple de 15
			case currentRotate > 0 && currentRotate%tiltAngleIncrement != 0:
				rotationToApply = strconv.Itoa(currentRotate - currentRotate%tiltAngleIncrement)
			case currentRotate < 0 && -currentRotate%tiltAngleIncrement != 0:
				rotationToApply = strconv.Itoa(currentRotate - (-currentRotate % tiltAngleIncrement))
			case currentRotate > 0:
				rotationToApply = strconv.Itoa(currentRotate - tiltAngleIncrement)
			case currentRotate < 0:
				rotationToApply = strconv.Itoa(currentRotate + tiltAngleIncrement)
			default:
				rotationToApply = "0"
			}
		} else {
			//HSpeed trop élevé, il faut réduire la vitesse latéral en réduisant l'accelération et en basculant l'angle
			//Puis si nécessaire il faudra à nouveau accélérer
		}
	} else if parameter.X >= flatGround.EndX {
		//On doit aller sur la gauche

	}
	return rotationToApply, powerToApply
}

// ComputeThrustPower compute the thrust power that needs to be applied
func computeThrustPower(parameter FlightParameter) string {
	currentVSpeed := parameter.VSpeed
	currentPower := parameter.Power
	switch {
	case currentVSpeed > 39:
		currentPower--
		return strconv.Itoa(max(currentPower, 0))
	case currentVSpeed < -39:
		currentPower++
		return strconv.Itoa(min(currentPower, 4))
	default:
		return "0"
	}
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
