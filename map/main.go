package main

import "fmt"

func main() {

	// 1st method of declaring map
	// colors := map[string]string{
	// 	"red":  "#ff0000",
	// 	"blue": "#67bf21",
	// }

	// 2nd method of declaring map
	// var colors map[string]string

	// 3rd method of declaring a map
	colors := make(map[string]string)

	colors["white"] = "#ffff"
	colors["blue"] = "#23ade"
	colors["green"] = "#ff12e"

	// delete key and value in a map
	// delete(colors, 10)

	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}
