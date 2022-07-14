package main

import (
	"fmt"
	"os"

	"github.com/mattacton/meeting-lights/internal/keys"
	"golang.org/x/term"
)

func main() {
	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		fmt.Println(err)
		return
	}
	defer term.Restore(int(os.Stdin.Fd()), oldState)

	keyCh := make(chan []byte)
	go keys.Pressed(keyCh)

	for {
		keyPressed := <-keyCh
		fmt.Printf("%s", string(keyPressed))
		if 'x' == keyPressed[0] {
			return
		}
	}
}