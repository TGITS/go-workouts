package main

import "fmt"
//import "os"

/**
 * Auto-generated code below aims at helping you parse
 * the standard input according to the problem statement.
 * ---
 * Hint: You can use the debug stream to print initialTX and initialTY, if Thor seems not follow your orders.
 **/

func main() {
    // lightX: the X position of the light of power
    // lightY: the Y position of the light of power
    // initialTX: Thor's starting X position
    // initialTY: Thor's starting Y position
    var lightX, lightY, initialTX, initialTY int
    fmt.Scan(&lightX, &lightY, &initialTX, &initialTY)
    
    dx := initialTX - lightX
    dy := initialTY - lightY
    horizontalMove := ""
    verticalMove := ""
    
    for {
        // remainingTurns: The remaining amount of turns Thor can move. Do not remove this line.
        var remainingTurns int
        fmt.Scan(&remainingTurns)
        
        if dx > 0 {
            dx = dx + 1;
            horizontalMove = "W"
        } else if dx < 0 {
            dx = dx - 1;
            horizontalMove = "E"
        } else {
            horizontalMove = ""
        }
        
        if dy > 0 {
            dy = dy - 1;
            verticalMove = "N"
        } else if dy < 0 {
            dy = dy + 1;
            verticalMove = "S"
        } else {
            verticalMove = ""
        }
        
        // fmt.Fprintln(os.Stderr, "Debug messages...")
        
        // A single line providing the move to be made: N NE E SE S SW W or NW
        fmt.Println(verticalMove + horizontalMove)
    }
}