// Package lights handles manipulating physical lights.
//
// This package exclusively deals with calling Philips Hue
// APIs unless I find need to manipulate other types of lights
package lights

import "fmt"

func PrintLights(keys string) {
	fmt.Printf("Lights Printing '%s'", keys)
}