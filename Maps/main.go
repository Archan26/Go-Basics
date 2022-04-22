package main

import "fmt"

func main() {
	//Method #1
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#00ff00",
		"white": "#ffffff",
	}

	//Mehod #2
	// var colors map[string]string

	//Method #3
	// colors := make(map[string]string)
	// colors["White"] = "#ffffff"

	//Delete map value
	// delete(colors, "white")
	// fmt.Println(colors)

	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex for ", color, " is ", hex)
	}
}
