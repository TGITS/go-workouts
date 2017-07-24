package main

import "fmt"
import "os"
import "bufio"
import "strings"
import "math"
import "strconv"

/**
 * Auto-generated code below aims at helping you parse
 * the standard input according to the problem statement.
 **/

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Buffer(make([]byte, 1000000), 1000000)

    // n: the number of temperatures to analyse
    var n int
    scanner.Scan()
    fmt.Sscan(scanner.Text(),&n)

    scanner.Scan()
    temps := scanner.Text() // the n temperatures expressed as integers ranging from -273 to 5526
    fmt.Fprintln(os.Stderr, temps)

    temperatures := strings.Fields(temps)
    fmt.Fprintln(os.Stderr, temperatures, len(temperatures))

    nearest_zero_t := 0
    if len(temperatures) > 0 {
        temp, _ := strconv.Atoi(temperatures[0])
        min_abs_t := math.Abs(float64(temp))
        nearest_zero_t = temp
        for _, t := range temperatures {
            t_in_int, _ := strconv.Atoi(t)
            abs_t_in_int := math.Abs(float64(t_in_int))
            if min_abs_t > abs_t_in_int {
                nearest_zero_t = t_in_int
                min_abs_t = abs_t_in_int
            }
            if min_abs_t == abs_t_in_int {
                if t_in_int > 0 {
                    nearest_zero_t = t_in_int
                }
            }
        }
    }


    // fmt.Fprintln(os.Stderr, "Debug messages...")
    fmt.Println(nearest_zero_t)// Write answer to stdout
}
