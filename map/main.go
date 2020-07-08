package main

import "fmt"

func main() {
	// colors := map[string]string{
	// 	"red":   "#ff0000",
	// 	"green": "#4bf745",
	// }
	// fmt.Println(colors)

	// First approach for creating an empty map
	// var colors map[string]string
	// Second approach for creating an empty map
	colors := make(map[string]string)
	colors["white"] = "#ffffff"
	// Deleting key value pairs
	delete(colors, "white")
	fmt.Println(colors)
}
