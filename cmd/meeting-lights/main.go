package main

import (
	"fmt"
	"os"

	"github.com/mattacton/meeting-lights/internal/keys"
	"github.com/mattacton/meeting-lights/internal/lights"
	"github.com/mattacton/meeting-lights/internal/you"
	"golang.org/x/term"
)

func main() {
	var doeet = new(you.Doeet)
	doeet.DoWhat = map[string]you.Thees {
		"jj": lights.PrintLights,
		"a": lights.PrintLights,
	}

	// Get a continuous stream of key presses
	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		fmt.Println(err)
		return
	}
	defer term.Restore(int(os.Stdin.Fd()), oldState)

	keyCh := keys.Bundled()

	for {
		select {
		case keyBundle := <-keyCh: {
			doeet.Now(keyBundle)
			if keyBundle == "x" {
				os.Exit(0)
			}
		}
		}
	}
}