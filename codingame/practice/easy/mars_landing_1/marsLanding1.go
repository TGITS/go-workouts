package main

import (
	"fmt"
	"strconv"
	//"os"
)

/**
 * Auto-generated code below aims at helping you parse
 * the standard input according to the problem statement.
 **/

func main() {
	// surfaceN: the number of points used to draw the surface of Mars.
	var surfaceN int
	fmt.Scan(&surfaceN)

	for i := 0; i < surfaceN; i++ {
		// landX: X coordinate of a surface point. (0 to 6999)
		// landY: Y coordinate of a surface point. By linking all the points together in a sequential fashion, you form the surface of Mars.
		var landX, landY int
		fmt.Scan(&landX, &landY)
	}
	for {
		// hSpeed: the horizontal speed (in m/s), can be negative.
		// vSpeed: the vertical speed (in m/s), can be negative.
		// fuel: the quantity of remaining fuel in liters.
		// rotate: the rotation angle in degrees (-90 to 90).
		// power: the thrust power (0 to 4).
		var X, Y, hSpeed, vSpeed, fuel, rotate, power int
		fmt.Scan(&X, &Y, &hSpeed, &vSpeed, &fuel, &rotate, &power)

		// fmt.Fprintln(os.Stderr, "Debug messages...")

		// 	For a landing to be successful, the ship must:

		// land on flat ground
		// land in a vertical position (tilt angle = 0°)
		// vertical speed must be limited ( ≤ 40m/s in absolute value)
		// horizontal speed must be limited ( ≤ 20m/s in absolute value)

		// 2 integers: rotate power. rotate is the desired rotation angle (should be 0 for level 1), power is the desired thrust power (0 to 4).
		fmt.Printf("0 %s\n", computeThrustPower(power, vSpeed))
	}
}

// ComputeThrustPower compute the thrust power that needs to be applied
func computeThrustPower(currentPower int, currentVSpeed int) string {
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
