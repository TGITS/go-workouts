package main

import (
	"fmt"
	"math"
	"math/cmplx"
	"os"
	//"os"
	"sort"
	"strconv"
)

const (
	CheckpointCoreSize = 50
	//MaxLoopBeforeBoost = 3
	//MaxLoopBeforeShield = 10
	PodRadius        = 400
	CheckpointRadius = 600
)

// ToRadians convert the parameters corresponding to an angle in degrees int an angle in radians.
// radians  = degrees * PI / 180
func ToRadians(degrees float64) float64 { return float64(degrees) * math.Pi / 180.0 }

// ToDegrees convert the parameters corresponding to an angle in radians int an angle in degrees.
// degreess = radians * 180 / PI
func ToDegrees(radians float64) float64 { return float64(radians) * 180.0 / math.Pi }

// Round rounds a float64, converts it to an int and returns it
func Round(val float64) int {
	if val < 0 {
		return int(val - 0.5)
	}
	return int(val + 0.5)
}

func computeAngle(a complex128, b complex128, hypothenuse float64) (float64, float64, float64) {
	c := complex(real(a), imag(b))
	// hypothenuse := cmplx.Abs(a - b)
	cosTheta := cmplx.Abs(c-b) / hypothenuse
	sinTheta := cmplx.Abs(a-c) / hypothenuse
	theta := ToDegrees(math.Acos(cosTheta))
	return cosTheta, sinTheta, theta
}

// TurnInfo is a struct that gather all the game information of a turn
type TurnInfo struct {
	Checkpoint  complex128
	PlayerPod   complex128
	OpponentPod complex128
	Angle       int //Angle between the player pod and the next Checkpoint in degrees
	Distance    int //Distance between the player pod and the next Checkpoint

}

// NewTurnInfo to create and initialize a new TurnInfo (a pointer to TurnInfo)
func NewTurnInfo(x, y, nextCheckpointX, nextCheckpointY, nextCheckpointDist, nextCheckpointAngle, opponentX, opponentY int) *TurnInfo {
	turnInfo := new(TurnInfo)
	turnInfo.Checkpoint = complex(float64(nextCheckpointX), float64(nextCheckpointY))
	turnInfo.PlayerPod = complex(float64(x), float64(y))
	turnInfo.OpponentPod = complex(float64(opponentX), float64(opponentY))
	turnInfo.Angle = nextCheckpointAngle
	turnInfo.Distance = nextCheckpointDist
	return turnInfo
}

/*
computeThrust compute the thrust value
*/
func (turnInfo *TurnInfo) computeThrust() string {
	switch {
	case turnInfo.Distance < 1000:
		return "20"
	case turnInfo.Distance < 600:
		return "00"
	case math.Abs(float64(turnInfo.Angle)) >= 90:
		return "0"
	case math.Abs(float64(turnInfo.Angle)) >= 72:
		return "20"
	case math.Abs(float64(turnInfo.Angle)) >= 54:
		return "40"
	case math.Abs(float64(turnInfo.Angle)) >= 36:
		return "60"
	case math.Abs(float64(turnInfo.Angle)) >= 18:
		return "80"
	default:
		return "100"
	}
}

func (turnInfo *TurnInfo) computeTarget() (complex128, int, int) {
	// intermediate := complex(real(pod), imag(checkpoint))
	// hypothenuse := cmplx.Abs(pod - checkpoint)
	// cosTheta := cmplx.Abs(intermediate-checkpoint) / hypothenuse
	// sinTheta := cmplx.Abs(pod-intermediate) / hypothenuse
	//cosTheta, sinTheta, theta := computeAngle(pod, checkpoint, float64(nextCheckpointDistance))
	cosTheta, sinTheta, _ := computeAngle(turnInfo.PlayerPod, turnInfo.Checkpoint, float64(turnInfo.Distance))
	// fmt.Fprintf(os.Stderr, "cosTheta : %.1f\n", cosTheta)
	// fmt.Fprintf(os.Stderr, "sinTheta : %.1f\n", sinTheta)
	// fmt.Fprintf(os.Stderr, "theta : %.1f\n", theta)
	// fmt.Fprintf(os.Stderr, "hypothenuse : %.1f\n", hypothenuse)
	var xTarget, yTarget float64
	distances := make([]float64, 5, 5)
	angles := make([]float64, 5, 5)
	targetsByDistance := make(map[float64]complex128)
	targetsByAngle := make(map[float64]complex128)

	distances[0] = float64(turnInfo.Distance)
	angles[0] = math.Abs(float64(turnInfo.Angle))
	targetsByDistance[distances[0]] = turnInfo.Checkpoint
	targetsByAngle[angles[0]] = turnInfo.Checkpoint
	index := 1
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			if i%2 != 0 {
				xTarget = real(turnInfo.Checkpoint) - PodRadius*cosTheta
			} else {
				xTarget = real(turnInfo.Checkpoint) + PodRadius*cosTheta
			}

			if j%2 != 0 {
				yTarget = imag(turnInfo.Checkpoint) + PodRadius*sinTheta
			} else {
				yTarget = imag(turnInfo.Checkpoint) - PodRadius*sinTheta
			}
			target := complex(xTarget, yTarget)
			distances[index] = cmplx.Abs(turnInfo.PlayerPod - target)
			_, _, angle := computeAngle(turnInfo.PlayerPod, target, distances[index])
			targetsByDistance[distances[index]] = target
			angles[index] = math.Abs(angle)
			targetsByAngle[angles[index]] = target
			index++
		}
	}
	sort.Float64s(distances)
	sort.Float64s(angles)
	fmt.Fprintf(os.Stderr, "Distances from pod to potential target : %.1f\n", distances)
	fmt.Fprintf(os.Stderr, "targetsByDistance : %v\n", targetsByDistance)
	fmt.Fprintf(os.Stderr, "Absolute value of the angle between pod and potential target : %.1f\n", angles)
	fmt.Fprintf(os.Stderr, "targetsByAngle : %v\n", targetsByAngle)
	//return targetsByDistance[distances[0]]
	return targetsByAngle[angles[0]], int(distances[0]), int(angles[0])
}

// GameInfo is a struct that gather all the global game information relative to the game
type GameInfo struct {
	Checkpoints []complex128
	History     map[int]*TurnInfo
	Boost       bool //True if a boost is available, false if the thrust of the player pod has already been used
	Shield      int  //To mark if the shield has been used. With possible value 0 (not recently used or more than 3 turns ; the initial value), 3 (just used), 2 (used 1 turn ago), 1 (used 2 turns ago)
	Turn        int  // 0 is the initial value
	Lap         int  //The lap in which we are. This value has to be computed

}

// NewGameInfo To create and initialize a new GameInfo (a pointer to GameInfo)
// Essentially the game board for CSB in silver league is a slice of checkpoints
// This slice is poupulated turns by turns, we know we are in the second lap when we cannot had a new CP
// We know we are in the third lap when we have reached the first checkpoint
func NewGameInfo() *GameInfo {
	gameInfo := new(GameInfo)
	gameInfo.Boost = true
	gameInfo.Shield = 3
	gameInfo.Turn = 0
	gameInfo.Lap = 0
	gameInfo.Checkpoints = make([]complex128, 0, 10)
	gameInfo.History = make(map[int]*TurnInfo)
	return gameInfo
}

// UpdateHistory is a method that had a new TurnInfo in the GameInfo. Before it update the turn count, the checkpoint list and the Lap count if necessary
func (gameInfo *GameInfo) UpdateHistory(turnInfo *TurnInfo) {
	gameInfo.UpdateTurn()
	if gameInfo.Lap == 0 {
		gameInfo.UpdateCheckpoints(turnInfo.Checkpoint)
	}
	gameInfo.History[gameInfo.Turn] = turnInfo
}

// PreviousTurn is a method that return the counter value of the previous turn
func (gameInfo *GameInfo) PreviousTurn() int {
	return gameInfo.Turn - 1
}

// NextTurn is a method that return the counter value of the next turn
func (gameInfo *GameInfo) NextTurn() int {
	return gameInfo.Turn + 1
}

// UpdateTurn is a method that updates the turn counter and return an int representing this new value
func (gameInfo *GameInfo) UpdateTurn() {
	gameInfo.Turn++
}

// UpdateCheckpoints is a method that update the list of checkpoints if necessary
func (gameInfo *GameInfo) UpdateCheckpoints(checkpoint complex128) {
	newCheckpoint, newLap := gameInfo.In(checkpoint)
	if !newCheckpoint {
		gameInfo.Checkpoints = append(gameInfo.Checkpoints, checkpoint)
	}
	if newLap {
		gameInfo.Lap++
	}
}

// In is a method that checks if the complex number is in the slice Checkpoints and if it is a new turn
func (gameInfo *GameInfo) In(complex complex128) (bool, bool) {
	for index, value := range gameInfo.Checkpoints {
		if complex == value {
			if index == 0 {
				return true, true
			}
			return true, false
		}
	}
	return false, false
}

func (gameInfo *GameInfo) updateBoost() {
	gameInfo.Boost = false
}

func (gameInfo *GameInfo) getCurrentTurnInfo() *TurnInfo {
	return gameInfo.History[gameInfo.Turn]
}

// DumpInfo is a method to dump the information on stderr
func (gameInfo *GameInfo) DumpInfo() {
	fmt.Fprintf(os.Stderr, "Turn : %d\n", gameInfo.Turn)
	fmt.Fprintf(os.Stderr, "Lap : %d\n", gameInfo.Lap)
	fmt.Fprintf(os.Stderr, "Checkpoints : %v\n", gameInfo.Checkpoints)
	fmt.Fprintf(os.Stderr, "Boost use ? : %t\n", gameInfo.Boost)
	turnInfo := gameInfo.getCurrentTurnInfo()
	fmt.Fprintf(os.Stderr, "Checkpoint coordinates : %.1f\n", turnInfo.Checkpoint)
	fmt.Fprintf(os.Stderr, "Palyer Pod coordinates : %.1f\n", turnInfo.PlayerPod)
	fmt.Fprintf(os.Stderr, "Opponent pod coordinates : %.1f\n", turnInfo.OpponentPod)
	fmt.Fprintf(os.Stderr, "Angle in degrees between pod and next checkpoint : %d\n", turnInfo.Angle)
	fmt.Fprintf(os.Stderr, "Distance from pod to next checkpoint : %d\n", turnInfo.Distance)
}

func (gameInfo *GameInfo) computeValues() (string, string, string) {
	turnInfo := gameInfo.getCurrentTurnInfo()
	target, _, _ := turnInfo.computeTarget()
	return strconv.Itoa(int(real(target))), strconv.Itoa(int(imag(target))), gameInfo.computeAction()
	//return strconv.Itoa(int(real(checkpoint))), strconv.Itoa(int(imag(checkpoint))), computeAction(checkpoint, nextCheckpointDistance, nextCheckpointAngle, pod, opponentPod, boostUsed, turns)
}

/*
computeAction return "BOOST", "SHIELD" or the thrust value
*/
func (gameInfo *GameInfo) computeAction() string {
	turnInfo := gameInfo.getCurrentTurnInfo()
	if gameInfo.Boost && math.Abs(float64(turnInfo.Angle)) < 10 && turnInfo.Distance > 4000 {
		gameInfo.updateBoost()
		return "BOOST"
	}
	return turnInfo.computeThrust()
}

/**
 * Auto-generated code below aims at helping you parse
 * the standard input according to the problem statement.
 **/

func main() {

	var xTarget, yTarget, thrust string
	gameInfo := NewGameInfo()
	for {
		// nextCheckpointX: x position of the next check point
		// nextCheckpointY: y position of the next check point
		// nextCheckpointDist: distance to the next checkpoint
		// nextCheckpointAngle: angle between your pod orientation and the direction of the next checkpoint
		var x, y, nextCheckpointX, nextCheckpointY, nextCheckpointDist, nextCheckpointAngle int
		fmt.Scan(&x, &y, &nextCheckpointX, &nextCheckpointY, &nextCheckpointDist, &nextCheckpointAngle)

		var opponentX, opponentY int
		fmt.Scan(&opponentX, &opponentY)

		turnInfo := NewTurnInfo(x, y, nextCheckpointX, nextCheckpointY, nextCheckpointDist, nextCheckpointAngle, opponentX, opponentY)
		gameInfo.UpdateHistory(turnInfo)

		gameInfo.DumpInfo()

		// You have to output the target position
		// followed by the power (0 <= thrust <= 100) or "BOOST" or "SHIELD"
		// i.e.: "x y thrust"
		xTarget, yTarget, thrust = gameInfo.computeValues()
		// thrustByTurn[turns] = thrust
		fmt.Printf("%s %s %s\n", xTarget, yTarget, thrust)
	}
}
