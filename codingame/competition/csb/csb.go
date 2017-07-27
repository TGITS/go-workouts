package main

import (
	"fmt"
	"math"
	"math/cmplx"
	"strconv"
)

const (
	CheckpointCoreSize  = 50
	MaxLoopBeforeBoost  = 3
	MaxLoopBeforeShield = 10
	PodRadius           = 400
	CheckpointRadius    = 600
)

// ToRadians convert the parameters corresponding to an angle in degrees int an angle in radians.
// radians  = degrees * PI / 180
func ToRadians(angleInDegrees float64) float64 { return angleInDegrees * math.Pi / 180 }

// ToDegrees convert the parameters corresponding to an angle in radians int an angle in degrees.
// degreess = radians * 180 / PI
func ToDegrees(angleInRadians float64) float64 { return angleInRadians * 180 / math.Pi }

//Round a float64 and convert it to an int32
func round(val float64) int32 {
	if val < 0 {
		return int32(val - 0.5)
	}
	return int32(val + 0.5)
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
		// if _, ok := checkpoints[checkpoint]; ok != true {
		// 	checkpoints[checkpoint] = struct{}{}
		// }

		// fmt.Fprintln(os.Stderr, "Debug messages...")
		// fmt.Fprintf(os.Stderr, "Turn number : %d\n", turns)
		// fmt.Fprintf(os.Stderr, "Pod coordinates : %.1f\n", pod)
		// fmt.Fprintf(os.Stderr, "Opponent pod coordinates : %.1f\n", opponentPod)
		// fmt.Fprintf(os.Stderr, "Checkpoint coordinates : %.1f\n", checkpoint)
		// fmt.Fprintf(os.Stderr, "Angle in degrees between pod and next checkpoint : %d\n", nextCheckpointAngle)
		// fmt.Fprintf(os.Stderr, "List of checkpoints : %v\n", checkpoints)
		// fmt.Fprintf(os.Stderr, "Number of checkpoints : %v\n", len(checkpoints))
		// fmt.Fprintf(os.Stderr, "Boost use ? : %t\n", *boostUsed)
		// fmt.Fprintf(os.Stderr, "Distance from pod to next checkpoint : %d\n", nextCheckpointDist)
		// fmt.Fprintf(os.Stderr, "Calculated distance between pod and checkpoint : %f\n", cmplx.Abs(checkpoint-pod))
		// fmt.Fprintf(os.Stderr, "Calculated distance between opponent pod and checkpoint : %f\n", cmplx.Abs(checkpoint-opponentPod))
		// fmt.Fprintf(os.Stderr, "Calculated distance between pod and opponent pod : %f\n", cmplx.Abs(opponentPod-pod))

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
	target := computeTarget(checkpoint, nextCheckpointDistance, pod)
	return strconv.Itoa(int(real(target))), strconv.Itoa(int(imag(target))), computeAction(checkpoint, nextCheckpointDistance, nextCheckpointAngle, pod, opponentPod, boostUsed, turns)
	//return strconv.Itoa(int(real(checkpoint))), strconv.Itoa(int(imag(checkpoint))), computeAction(checkpoint, nextCheckpointDistance, nextCheckpointAngle, pod, opponentPod, boostUsed, turns)
}

/*
computeAction return "BOOST", "SHIELD" or the thrust value
*/
func computeAction(checkpoint complex128, nextCheckpointDistance int, nextCheckpointAngle int, pod complex128, opponentPod complex128, boostUsed *bool, turns int) string {
	if !*boostUsed && turns > MaxLoopBeforeBoost && math.Abs(float64(nextCheckpointAngle)) < 10 && nextCheckpointDistance > 4000 {
		*boostUsed = true
		return "BOOST"
	}
	return computeThrust(checkpoint, nextCheckpointDistance, nextCheckpointAngle, pod, opponentPod, boostUsed, turns)
}

/*
computeThrust compute the thrust value
(defn compute-thrust [next-checkpoint-distance next-checkpoint-angle]
  (cond
    (> (Math/abs next-checkpoint-angle) 90) "0"
    ; (and (< next-checkpoint-distance 1000) (< (Math/abs next-checkpoint-angle) 30)) "50"
    ; (< (Math/abs next-checkpoint-angle) 30) "100"
    (< next-checkpoint-distance 600) "20"
    (and (< next-checkpoint-distance 1000) (> (Math/abs next-checkpoint-angle) 25)) "0"
    (< next-checkpoint-distance 1000) "40"
    (< next-checkpoint-distance 2000) "80"
    true "100"))
*/
func computeThrust(checkpoint complex128, nextCheckpointDistance int, nextCheckpointAngle int, pod complex128, opponentPod complex128, boostUsed *bool, turns int) string {
	switch {
	case math.Abs(float64(nextCheckpointAngle)) > 90 /* || nextCheckpointDistance < 600*/ :
		return "0"
	case nextCheckpointDistance < 600 && math.Abs(float64(nextCheckpointAngle)) > 45:
		return "0"
	case nextCheckpointDistance < 600 || (nextCheckpointDistance < 1000 && math.Abs(float64(nextCheckpointAngle)) > 45):
		return "20"
	case nextCheckpointDistance < 1000:
		return "40"
	// case nextCheckpointDistance < 1200:
	// 	return "60"
	// case nextCheckpointDistance < 1500:
	// 	return "80"
	// case nextCheckpointDistance < 100 && math.Abs(float64(nextCheckpointAngle)) > 25:
	// 	return "0"
	// case nextCheckpointDistance < 1000:
	// 	return "50"
	// case nextCheckpointDistance < 2000:
	// 	return "90"
	case *boostUsed && turns > 10 && cmplx.Abs(opponentPod-pod) <= 900 && math.Abs(float64(nextCheckpointAngle)) > 90:
		return "SHIELD"
	default:
		return "100"
	}
}

func computeTarget(checkpoint complex128, nextCheckpointDistance int, pod complex128) complex128 {
	if nextCheckpointDistance > 1000 {
		intermediate := complex(real(pod), imag(checkpoint))
		hypothenuse := cmplx.Abs(pod - checkpoint)
		cosTheta := cmplx.Abs(intermediate-checkpoint) / hypothenuse
		sinTheta := cmplx.Abs(pod-intermediate) / hypothenuse
		var xTarget, yTarget float64
		if real(checkpoint) > real(pod) {
			xTarget = real(checkpoint) - PodRadius*cosTheta
		} else {
			xTarget = real(checkpoint) + PodRadius*cosTheta
		}
		if imag(checkpoint) > imag(pod) {
			yTarget = imag(checkpoint) - PodRadius*sinTheta
		} else {
			yTarget = imag(checkpoint) + PodRadius*sinTheta
		}
		return complex(xTarget, yTarget)
	}
	return checkpoint
}
