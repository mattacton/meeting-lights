package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/mattacton/meeting-lights/internal/keys"
	"github.com/mattacton/meeting-lights/internal/lights"
	"github.com/mattacton/meeting-lights/internal/you"
	"golang.org/x/term"
)

var doeet you.Doeet
var theLights lights.Lights

func getEnvs() (host, secret string, ids []string) {
	host = os.Getenv("LIGHT_HOST")
	secret = os.Getenv("LIGHT_SECRET")
	idString := os.Getenv("LIGHT_IDS")
	if len(host) == 0 || len(secret) == 0 || len(idString) == 0 {
		fmt.Println("Error LIGHT_HOST, LIGHT_SECRET, or LIGHT_IDS not set! Exiting")
		os.Exit(1)
	}
	ids = strings.Split(idString, ",")
	return
}

func init() {
	host, secret, lightIds := getEnvs()
	theLights = *lights.NewLights(host, secret, lightIds)
	theLights.OriginalStates = theLights.GetCurrentState()
	fmt.Printf("\nThe state: %+v \n\n", theLights.GetCurrentState())

	doeet = you.Doeet{
		DoWhat: map[string]func(keys string){
			"r": theLights.TurnRed,
			"b": theLights.TurnBlue,
			"g": theLights.TurnGreen,
			"n": theLights.TurnNormal,
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
		keyBundle := <-keyCh
		doeet.Now(keyBundle)
		if keyBundle == "x" {
			areWeDone = true
		}
	}

	theLights.ResetLight()
}
