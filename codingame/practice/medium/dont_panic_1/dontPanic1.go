package main

import "fmt"

//import "os"

const (
	LEFT  = "LEFT"
	RIGHT = "RIGHT"
	WAIT  = "WAIT"
	BLOCK = "BLOCK"
)

/**
 * Auto-generated code below aims at helping you parse
 * the standard input according to the problem statement.
 **/

func main() {
	// nbFloors: number of floors
	// width: width of the area
	// nbRounds: maximum number of rounds
	// exitFloor: floor on which the exit is found
	// exitPos: position of the exit on its floor
	// nbTotalClones: number of generated clones
	// nbAdditionalElevators: ignore (always zero)
	// nbElevators: number of elevators
	var nbFloors, width, nbRounds, exitFloor, exitPos, nbTotalClones, nbAdditionalElevators, nbElevators int
	fmt.Scan(&nbFloors, &width, &nbRounds, &exitFloor, &exitPos, &nbTotalClones, &nbAdditionalElevators, &nbElevators)
	elevatorPosByFloor := make(map[int]int)
	for i := 0; i < nbElevators; i++ {
		// elevatorFloor: floor on which this elevator is found
		// elevatorPos: position of the elevator on its floor
		var elevatorFloor, elevatorPos int
		fmt.Scan(&elevatorFloor, &elevatorPos)
		elevatorPosByFloor[elevatorFloor] = elevatorPos
	}

	var action string

	for {
		// cloneFloor: floor of the leading clone
		// clonePos: position of the leading clone on its floor
		// direction: direction of the leading clone: LEFT or RIGHT
		var cloneFloor, clonePos int
		var direction string
		fmt.Scan(&cloneFloor, &clonePos, &direction)
		if cloneFloor < exitFloor {
			action = computeAction(clonePos, elevatorPosByFloor[cloneFloor], direction)
		} else if cloneFloor == exitFloor {
			action = computeAction(clonePos, exitPos, direction)
		}
		// fmt.Fprintln(os.Stderr, "Debug messages...")
		fmt.Println(action) // action: WAIT or BLOCK
	}
}

func computeAction(clonePos int, pos int, direction string) string {
	var action string
	switch {
	case clonePos < pos && direction == RIGHT:
		action = WAIT
	case clonePos < pos && direction == LEFT:
		action = BLOCK
	case clonePos > pos && direction == RIGHT:
		action = BLOCK
	case clonePos < pos && direction == LEFT:
		action = WAIT
	default:
		action = WAIT
	}
	return action
}
