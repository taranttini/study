package main

import (
	"fmt"

	"/User/cache"
	//"./cache"
)

func main() {
	// Create a new Cache instance
	myCache := cache.New[string, int]()

	// Set key-value pairs in the cache
	myCache.Set("one", 1)
	myCache.Set("two", 2)
	myCache.Set("three", 3)

	// Retrieve values from the cache
	value, found := myCache.Get("two")
	if found {
		fmt.Printf("Value for key 'two': %v\n", value)
	} else {
		fmt.Println("Key 'two' not found in the cache")
	}

	// Pop a key from the cache
	poppedValue, found := myCache.Pop("three")
	if found {
		fmt.Printf("Popped value for key 'three': %v\n", poppedValue)
	} else {
		fmt.Println("Key 'three' not found in the cache")
	}

	// Remove a key from the cache
	myCache.Remove("one")

	// Try to retrieve a removed key
	removedValue, found := myCache.Get("one")
	if found {
		fmt.Printf("Value for key 'one': %v\n", removedValue)
	} else {
		fmt.Println("Key 'one' not found in the cache (after removal)")
	}
}
