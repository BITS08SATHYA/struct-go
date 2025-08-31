package main

import "fmt"

func main() {

	//var colors map[string]string

	//colors := make(map[int]string)

	//colors[1] = "#FF0000"
	//colors[2] = "#0000FF"
	//
	//delete(colors, 1)

	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
		"white": "#ffffff",
	}

	//fmt.Println(colors)

	printMap(colors)
}

func printMap(c map[string]string) {
	for key, value := range c {
		fmt.Println("Hex code for ", key, " is", value)
	}
}
