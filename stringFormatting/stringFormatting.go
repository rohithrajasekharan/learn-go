package main

import (
	"fmt"
	"os"
)

type point struct {
	x, y int
}

func main() {
	p := point{1, 2}
	fmt.Printf("struct 1: %v\n", p)
	fmt.Printf("struct 2: %+v\n", p)
	fmt.Printf("struct 3: %#v\n", p)
	fmt.Printf("type: %T\n", p)
	fmt.Printf("bool: %t\n", true)
	fmt.Printf("int: %d\n", 3)
	fmt.Printf("binary: %b\n", 32)
	fmt.Printf("char: %c\n", 33)
	fmt.Printf("hex: %x\n", 456)
	fmt.Printf("float1: %f\n", 78.9)
	fmt.Printf("float2: %e\n", 12300000.9)
	fmt.Printf("float3: %E\n", 78000000.9)
	fmt.Printf("string1: %s\n", "\"string\"")
	fmt.Printf("string2: %q\n", "\"string\"")
	fmt.Printf("string3: %x\n", "hex this")
	fmt.Printf("Pointer: %p\n", &p)
	fmt.Printf("width1: |%6d|%6d|\n", 12, 345)
	fmt.Printf("width2: |%6.2f|%6.2f|\n", 1.2, 3.45)
	fmt.Printf("width3: |%6s|%6s|\n", "foo", "b")
	fmt.Printf("width3: |%-6s|%-6s|\n", "foo", "b")

	s := fmt.Sprintf("sprintf: a %s", "string")
	fmt.Println(s)

	fmt.Fprintf(os.Stderr, "io: an %s\n", "error")
}
