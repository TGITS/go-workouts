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
func ToRadians(degrees int) int { return round(float64(degrees) * math.Pi / 180.0) }

// ToDegrees convert the parameters corresponding to an angle in radians int an angle in degrees.
// degreess = radians * 180 / PI
func ToDegrees(radians int) int { return round(float64(radians) * 180.0 / math.Pi) }

//Round a float64 and convert it to an int32
func round(val float64) int {
	if val < 0 {
		return int(val - 0.5)
	}
	return int(val + 0.5)
}

//Representing the coordinate of a Game Element
type Point struct {
	X, Y int
}

//Calculating the Distance
func (p1 *Point) Distance(p2 *Point) int {
	x := p2.X - p1.X
	y := p2.Y - p1.Y
	return int((round(math.Sqrt(float64(x*x + y*y)))))
}

//representing a Pod in CSB
type Pod struct {
	Point
	AngleToCP    int
	DistanceToCp int
	Thrust       bool
	Shield       int // with possible value 0 (not recently used or more than 3), 3 (just used), 2 (used 1 turn ago), 1 (used 2 turns ago)
}

type Checkpoint struct {
	Point
}

/**
 * Auto-generated code below aims at helping you parse
 * the standard input according to the problem statement.
 **/

func main() {

	// checkpoints := make(map[complex128]struct{})
	var checkpoint complex128
	var pod complex128
	var opponentPod complex128
	var xTarget, yTarget, thrust string
	// var checkpointAngle, podAngle, opponentPodAngle, angleBetweenPods float64
	turns := 0
	boostUsed := new(bool)
	*boostUsed = false
	// thrustByTurn := make(map[int]string)
	// thrustByTurn[turns] = "0"
	for {
		// nextCheckpointX: x position of the next check point
		// nextCheckpointY: y position of the next check point
		// nextCheckpointDist: distance to the next checkpoint
		// nextCheckpointAngle: angle between your pod orientation and the direction of the next checkpoint
		var x, y, nextCheckpointX, nextCheckpointY, nextCheckpointDist, nextCheckpointAngle int
		fmt.Scan(&x, &y, &nextCheckpointX, &nextCheckpointY, &nextCheckpointDist, &nextCheckpointAngle)

		var opponentX, opponentY int
		fmt.Scan(&opponentX, &opponentY)

		turns++
		pod = complex(float64(x), float64(y))
		checkpoint = complex(float64(nextCheckpointX), float64(nextCheckpointY))
		opponentPod = complex(float64(opponentX), float64(opponentY))
		//distanceBetweenPods := int(cmplx.Abs(opponentPod - pod))
		// if _, ok := checkpoints[checkpoint]; ok != true {
		// 	checkpoints[checkpoint] = struct{}{}
		// }

		// fmt.Fprintln(os.Stderr, "Debug messages...")
		// fmt.Fprintf(os.Stderr, "Turn number : %d\n", turns)
		fmt.Fprintf(os.Stderr, "Pod coordinates : %.1f\n", pod)
		fmt.Fprintf(os.Stderr, "Opponent pod coordinates : %.1f\n", opponentPod)
		fmt.Fprintf(os.Stderr, "Checkpoint coordinates : %.1f\n", checkpoint)
		fmt.Fprintf(os.Stderr, "Angle in degrees between pod and next checkpoint : %d\n", nextCheckpointAngle)
		// fmt.Fprintf(os.Stderr, "List of checkpoints : %v\n", checkpoints)
		// fmt.Fprintf(os.Stderr, "Number of checkpoints : %v\n", len(checkpoints))
		// fmt.Fprintf(os.Stderr, "Boost use ? : %t\n", *boostUsed)
		// fmt.Fprintf(os.Stderr, "Distance from pod to next checkpoint : %d\n", nextCheckpointDist)
		// fmt.Fprintf(os.Stderr, "Calculated distance between pod and checkpoint : %f\n", cmplx.Abs(checkpoint-pod))
		//fmt.Fprintf(os.Stderr, "Calculated distance between pod and opponent pod : %d\n", distanceBetweenPods)

		// _, checkpointAngle = cmplx.Polar(checkpoint)
		// _, podAngle = cmplx.Polar(pod)
		// _, opponentPodAngle = cmplx.Polar(opponentPod)
		// fmt.Fprintf(os.Stderr, "Calculated phases in radians for checkpoint, pod and opponentPod : %.1f %.1f %.1f\n", checkpointAngle, podAngle, opponentPodAngle)
		// _, angleBetweenPods = cmplx.Polar((pod - checkpoint) / (opponentPod - checkpoint))
		// fmt.Fprintf(os.Stderr, "Calculated angle in radians between the 2 vectors (Checkpoint, Pod) and (Checkpoint, OpponentPod) : %.3f\n", angleBetweenPods)
		// fmt.Fprintf(os.Stderr, "Calculated angle in degrees between the 2 vectors (Checkpoint, Pod) and (Checkpoint, OpponentPod) : %.3f\n", ToDegrees(angleBetweenPods))
		//ToDegrees
		// You have to output the target position
		// followed by the power (0 <= thrust <= 100) or "BOOST" or "SHIELD"
		// i.e.: "x y thrust"
		xTarget, yTarget, thrust = computeValues(checkpoint, nextCheckpointDist, nextCheckpointAngle, pod, opponentPod, boostUsed, turns)
		// thrustByTurn[turns] = thrust
		fmt.Printf("%s %s %s\n", xTarget, yTarget, thrust)
	}
}

func computeValues(checkpoint complex128, nextCheckpointDistance int, nextCheckpointAngle int, pod complex128, opponentPod complex128, boostUsed *bool, turns int) (string, string, string) {
	target, distance, angle := computeTarget(checkpoint, nextCheckpointDistance, nextCheckpointAngle, pod)
	return strconv.Itoa(int(real(target))), strconv.Itoa(int(imag(target))), computeAction(target, distance, angle, pod, opponentPod, boostUsed, turns)
	//return strconv.Itoa(int(real(checkpoint))), strconv.Itoa(int(imag(checkpoint))), computeAction(checkpoint, nextCheckpointDistance, nextCheckpointAngle, pod, opponentPod, boostUsed, turns)
}

/*
computeAction return "BOOST", "SHIELD" or the thrust value
*/
func computeAction(checkpoint complex128, nextCheckpointDistance int, nextCheckpointAngle int, pod complex128, opponentPod complex128, boostUsed *bool, turns int) string {
	// if !*boostUsed && turns > MaxLoopBeforeBoost && math.Abs(float64(nextCheckpointAngle)) < 10 && nextCheckpointDistance > 4000 {
	if !*boostUsed && math.Abs(float64(nextCheckpointAngle)) < 10 && nextCheckpointDistance > 4000 {
		*boostUsed = true
		return "BOOST"
	}
	return computeThrust(checkpoint, nextCheckpointDistance, nextCheckpointAngle, pod, opponentPod, boostUsed, turns)
}

/*
computeThrust compute the thrust value
*/
func computeThrust(checkpoint complex128, nextCheckpointDistance int, nextCheckpointAngle int, pod complex128, opponentPod complex128, boostUsed *bool, turns int) string {
	switch {
	case math.Abs(float64(nextCheckpointAngle)) >= 90:
		return "0"
	case math.Abs(float64(nextCheckpointAngle)) >= 72:
		return "20"
	case math.Abs(float64(nextCheckpointAngle)) >= 54:
		return "40"
	case math.Abs(float64(nextCheckpointAngle)) >= 36:
		return "60"
	case math.Abs(float64(nextCheckpointAngle)) >= 18:
		return "80"
	default:
		return "100"
	}
}

func computeAngle(a complex128, b complex128, hypothenuse float64) (float64, float64, float64) {
	c := complex(real(a), imag(b))
	// hypothenuse := cmplx.Abs(a - b)
	cosTheta := cmplx.Abs(c-b) / hypothenuse
	sinTheta := cmplx.Abs(a-c) / hypothenuse
	theta := ToDegrees(math.Acos(cosTheta))
	return cosTheta, sinTheta, theta
}

func computeTarget(checkpoint complex128, nextCheckpointDistance int, nextCheckpointAngle int, pod complex128) (complex128, int, int) {
	// intermediate := complex(real(pod), imag(checkpoint))
	// hypothenuse := cmplx.Abs(pod - checkpoint)
	// cosTheta := cmplx.Abs(intermediate-checkpoint) / hypothenuse
	// sinTheta := cmplx.Abs(pod-intermediate) / hypothenuse
	//cosTheta, sinTheta, theta := computeAngle(pod, checkpoint, float64(nextCheckpointDistance))
	cosTheta, sinTheta, _ := computeAngle(pod, checkpoint, float64(nextCheckpointDistance))
	// fmt.Fprintf(os.Stderr, "cosTheta : %.1f\n", cosTheta)
	// fmt.Fprintf(os.Stderr, "sinTheta : %.1f\n", sinTheta)
	// fmt.Fprintf(os.Stderr, "theta : %.1f\n", theta)
	// fmt.Fprintf(os.Stderr, "hypothenuse : %.1f\n", hypothenuse)
	var xTarget, yTarget float64
	distances := make([]float64, 5, 5)
	angles := make([]float64, 5, 5)
	targetsByDistance := make(map[float64]complex128)
	targetsByAngle := make(map[float64]complex128)

	distances[0] = float64(nextCheckpointDistance)
	angles[0] = math.Abs(float64(nextCheckpointAngle))
	targetsByDistance[distances[0]] = checkpoint
	targetsByAngle[angles[0]] = checkpoint
	index := 1
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			if i%2 != 0 {
				xTarget = real(checkpoint) - PodRadius*cosTheta
			} else {
				xTarget = real(checkpoint) + PodRadius*cosTheta
			}

			if j%2 != 0 {
				yTarget = imag(checkpoint) + PodRadius*sinTheta
			} else {
				yTarget = imag(checkpoint) - PodRadius*sinTheta
			}
			target := complex(xTarget, yTarget)
			distances[index] = cmplx.Abs(pod - target)
			_, _, angle := computeAngle(pod, target, distances[index])
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
