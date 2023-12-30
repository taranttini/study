package main

import "fmt"

// slices

func main() {
	s := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
	fmt.Printf("1 len=%d cap=%d %v\n", len(s), cap(s), s)

	fmt.Printf("2 len=%d cap=%d %v\n", len(s[:0]), cap(s[:0]), s[:0])
	fmt.Printf("3 len=%d cap=%d %v\n", len(s[:4]), cap(s[:4]), s[:4])
	fmt.Printf("4 len=%d cap=%d %v\n", len(s[2:]), cap(s[2:]), s[2:])

	s = append(s, 110)

	fmt.Printf("6 len=%d cap=%d %v\n", len(s), cap(s), s)
}
