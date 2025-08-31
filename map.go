package main

import "fmt"

func main() {

	//var colors map[string]string

	colors := make(map[int]string)

	//colors := map[string]string{
	//	"red":   "#ff0000",
	//	"green": "#4bf745",
	//}

	colors[1] = "#FF0000"
	colors[2] = "#0000FF"

	delete(colors, 1)

	fmt.Println(colors)
}
