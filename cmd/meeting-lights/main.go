package main

import (
	"fmt"
	"os"

	"github.com/mattacton/meeting-lights/internal/keys"
	"github.com/mattacton/meeting-lights/internal/lights"
	"github.com/mattacton/meeting-lights/internal/you"
	"golang.org/x/term"
)

var doeet you.Doeet

func init() {
	doeet = you.Doeet {
		DoWhat: map[string]you.Thees {
			"jj": lights.PrintLights,
			"a": lights.PrintLights,
		},
	}
}

func main() {

	// Get a continuous stream of key presses
	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		fmt.Println(err)
		return
	}
	defer term.Restore(int(os.Stdin.Fd()), oldState)

	keyCh := keys.Bundled()

	areWeDone := false
	for !areWeDone {
		select {
		case keyBundle := <-keyCh: {
			doeet.Now(keyBundle)
			if keyBundle == "x" {
				areWeDone = true
			}
		}
		}
	}
}