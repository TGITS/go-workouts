//From https://blog.golang.org/strings
package main

import (
	"fmt"
	"strings"
)

func main() {
	const sample = "\xbd\xb2\x3d\xbc\x20\xe2\x8c\x98"

	fmt.Println("Println:")
	fmt.Println(sample)

	fmt.Println("Byte loop with %x:")
	for i := 0; i < len(sample); i++ {
		fmt.Printf("%x ", sample[i])
	}
	fmt.Printf("\n")

	//[Exercise: Loop over the string using the %q format on each byte. What does the output tell you?]
	fmt.Println("Byte loop with %q:")
	for i := 0; i < len(sample); i++ {

		fmt.Printf("%q ", sample[i])
	}
	fmt.Printf("\n")

	fmt.Println("Printf with %x:")
	fmt.Printf("%x\n", sample)

	fmt.Println("Printf with % x:")
	fmt.Printf("% x\n", sample)

	fmt.Println("Printf with %q:")
	fmt.Printf("%q\n", sample)

	fmt.Println("Printf with %+q:")
	fmt.Printf("%+q\n", sample)

	//[Exercise: Modify the examples above to use a slice of bytes instead of a string. Hint: Use a conversion to create the slice.]
	fmt.Println("################################################################")
	fmt.Println("Same exercise but with an array of bytes created from the string")
	fmt.Println("################################################################")
	fmt.Println()
	sample2 := strings.Fields(sample)

	fmt.Println("Println:")
	fmt.Println(sample2)

	fmt.Println("Byte loop with %x:")
	for i := 0; i < len(sample2); i++ {
		fmt.Printf("%x ", sample2[i])
	}
	fmt.Printf("\n")

	//[Exercise: Loop over the string using the %q format on each byte. What does the output tell you?]
	fmt.Println("Byte loop with %q:")
	for i := 0; i < len(sample2); i++ {

		fmt.Printf("%q ", sample2[i])
	}
	fmt.Printf("\n")

	fmt.Println("Printf with %x:")
	fmt.Printf("%x\n", sample2)

	fmt.Println("Printf with % x:")
	fmt.Printf("% x\n", sample2)

	fmt.Println("Printf with %q:")
	fmt.Printf("%q\n", sample2)

	fmt.Println("Printf with %+q:")
	fmt.Printf("%+q\n", sample2)

}
