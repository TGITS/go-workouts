package main

import (
	"bufio"
	"fmt"
	"math/cmplx"
	"os"
	"strconv"
	"strings"
)

/**
 * Auto-generated code below aims at helping you parse
 * the standard input according to the problem statement.
 **/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(make([]byte, 1000000), 1000000)

	scanner.Scan()
	c := scanner.Text()
	fmt.Fprintln(os.Stderr, "Given input 'c' : ", c)
	myC := stringToComplex(c)
	fmt.Fprintln(os.Stderr, "Created complex from String : ", myC)
	var m int
	scanner.Scan()
	fmt.Sscan(scanner.Text(), &m)
	fmt.Fprintln(os.Stderr, "Given input 'm' : ", m)

	two := complex(2.0, 0.0)
	i := 0
	for f := complex(0.0, 0.0); i < m && cmplx.Abs(f) < 2; i++ {
		f = cmplx.Pow(f, two) + myC
	}

	fmt.Println(i)
}

func stringToComplex(s string) complex128 {
	var values []string
	coeffR := 1.0
	coeffIm := 1.0

	if strings.HasPrefix(s, "-") {
		coeffR = -1.0
		s = strings.TrimPrefix(s, "-")
	}

	countPlus := strings.Count(s, "+")

	if countPlus == 1 {
		values = strings.Split(s, "+")
	} else {
		values = strings.Split(s, "-")
		coeffIm = -1.0
	}

	fmt.Fprintln(os.Stderr, "Values : ", values)
	fmt.Fprintln(os.Stderr, "CoeffR : ", coeffR)
	fmt.Fprintln(os.Stderr, "CoeffIm : ", coeffIm)
	values[1] = strings.TrimSuffix(values[1], "i")
	re, _ := strconv.ParseFloat(values[0], 64)
	im, _ := strconv.ParseFloat(values[1], 64)
	fmt.Fprintln(os.Stderr, "re : ", re)
	fmt.Fprintln(os.Stderr, "im : ", im)
	return complex(coeffR*re, coeffIm*im)
}
