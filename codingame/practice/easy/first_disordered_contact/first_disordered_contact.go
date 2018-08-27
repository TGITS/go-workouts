package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	//"strings"
	//"strconv"
)

/**
 * Auto-generated code below aims at helping you parse
 * the standard input according to the problem statement.
 **/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(make([]byte, 1000000), 1000000)

	var N int
	scanner.Scan()
	fmt.Sscan(scanner.Text(), &N)

	scanner.Scan()
	MESSAGE := scanner.Text()

	fmt.Fprintln(os.Stderr, fmt.Sprintf("Input action : %d", N))
	fmt.Fprintln(os.Stderr, fmt.Sprintf("Input Message : %s", MESSAGE))
	/*
	 * An integer N indicating the number of times the message was transformed.
	 * If N is positive you have to decode i.e. retrieve the original message.
	 * If N is negative you have to encode i.e. transform the message.
	 */
	answer := MESSAGE
	if N < 0 {
		N = int(math.Abs(float64(N)))
		for i := 0; i < N; i++ {
			answer = Encode2(answer)
		}
	} else {
		for i := 0; i < N; i++ {
			answer = Decode(answer)
		}
	}
	fmt.Println(answer) // Write answer to stdout
}

func RegroupBySize(message string) ([]string, []int) {
	messageLen := len(message)
	groupSize := 1
	index := 0
	var syllables []string
	var syllablesSize []int
	for {
		if index+groupSize < messageLen {
			syllables = append(syllables, message[index:index+groupSize])
			syllablesSize = append(syllablesSize, groupSize)
		} else {
			syllables = append(syllables, message[index:])
			syllablesSize = append(syllablesSize, (messageLen - index))
			break
		}
		index += groupSize
		groupSize++
	}
	return syllables, syllablesSize
}

func Encode2(message string) string {
	group, _ := RegroupBySize(message)
	inFront := false
	encoded := ""
	for _, val := range group {
		if inFront {
			encoded = val + encoded
		} else {
			encoded = encoded + val
		}
		inFront = !inFront
	}
	return encoded
}

func Encode(message string) string {
	encoded := message[0:1]
	index := 1
	nbOfCharsToProcessed := 2
	//messageLenProcessed := 1
	inFront := true
	lenMessage := len(message)
	for index < lenMessage {
		if inFront {
			if index+nbOfCharsToProcessed > lenMessage {
				encoded = message[index:] + encoded
				break
			} else {
				encoded = message[index:index+nbOfCharsToProcessed] + encoded
			}
			inFront = false
		} else {
			endIndex := index + nbOfCharsToProcessed
			if index+nbOfCharsToProcessed > lenMessage {
				encoded = encoded + message[index:]
				break
			} else {
				encoded = encoded + message[index:endIndex]
			}
			inFront = true
		}
		index += nbOfCharsToProcessed
		//messageLenProcessed += nbOfCharsToProcessed
		nbOfCharsToProcessed++
	}

	return encoded
}

func FindGroupSize(message string) []int {
	messageLen := len(message)
	groupSize := 1
	index := 0
	var syllablesSize []int
	for {
		if index+groupSize < messageLen {
			syllablesSize = append(syllablesSize, groupSize)
		} else {
			syllablesSize = append(syllablesSize, (messageLen - index))
			break
		}
		index += groupSize
		groupSize++
	}
	return syllablesSize
}

//Reverse taken from http://golangcookbook.com/chapters/arrays/reverse/
func Reverse(numbers []int) []int {
	for i := 0; i < len(numbers)/2; i++ {
		j := len(numbers) - i - 1
		numbers[i], numbers[j] = numbers[j], numbers[i]
	}
	return numbers
}

func Decode(message string) string {
	//Cas dégénéré avec une chaine de longueur 1
	if len(message) == 1 {
		return message
	}

	fmt.Fprintln(os.Stderr, fmt.Sprintf("Original encoded message : %s", message))
	encoded := message
	decoded := ""
	groupSize := Reverse(FindGroupSize(message))
	fmt.Fprintln(os.Stderr, fmt.Sprintf("List of the group size in reverse order : %v", groupSize))
	groupSizeLen := len(groupSize)
	fmt.Fprintln(os.Stderr, fmt.Sprintf("Length of the group size : %d", groupSizeLen))

	takeInFront := true
	if groupSizeLen%2 != 0 {
		takeInFront = false
	}

	fmt.Fprintln(os.Stderr, fmt.Sprintf("Initial value of the decoded string : %s", decoded))
	fmt.Fprintln(os.Stderr, fmt.Sprintf("Initial value of the encoded string : %s", encoded))
	for _, val := range groupSize {
		maxEncodedIndex := len(encoded) - 1
		if takeInFront {
			decoded = encoded[0:val] + decoded
			encoded = encoded[val:]
		} else {
			decoded = encoded[maxEncodedIndex-val+1:] + decoded
			encoded = encoded[:maxEncodedIndex-val+1]
		}
		takeInFront = !takeInFront
		fmt.Fprintln(os.Stderr, fmt.Sprintf("Intermediate value of the decoded string : %s", decoded))
		fmt.Fprintln(os.Stderr, fmt.Sprintf("Intermediate value of the encoded string : %s", encoded))
	}
	fmt.Fprintln(os.Stderr, fmt.Sprintf("Final value of the decoded string : %s", decoded))
	fmt.Fprintln(os.Stderr, fmt.Sprintf("Final value of the encoded string : %s", encoded))
	return decoded
}
