package main

import "fmt"

func main() {
	//Method #1
	// colors := map[string]string{
	// 	"red":   "#ff0000",
	// 	"green": "#00ff00",
	// }

	//Mehod #2
	// var colors map[string]string

	//Method #3
	colors := make(map[string]string)

	colors["White"] = "#ffffff"

	//Delete map value
	delete(colors, "White")
	fmt.Println(colors)
}
