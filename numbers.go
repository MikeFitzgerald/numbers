package main

func main() {

	const factor = 3 // factor is compatible with any numeric type
	i := 20000 // i is of type int by inference
	i *= factor
	j := int16(20) // j is of type int16
	i += int(j) // Type must match so conversion is required
	k := uint8(0) // Same as: var k uint8
	k = uint8(i) // Succeeds, but k's value is truncated to 8 bits
	fmt.Println(i, j, k) // Prints: 60020 20 116

}