// Package keys implements functions for capturing key presses
// from STDIN and bundling them into strings when pressed within
// 500 milliseconds of each other.
//
// This package should only be used if STDIN has been changed to
// Raw mode via [golang.org/x/term] package
package keys

import (
	"fmt"
	"os"
	"strings"
	"time"
)

// Pressed returns a channelof the type string. Each message
// on the channel is a key that is pressed and sent to STDIN.
//
// This function should only be called after the given STDIN has
// been chnaged to Raw mode via the [term] package
func Pressed() chan string {
	keysPressedCh := make(chan string)
	go func() {
		for {
			b := make([]byte, 1)
			_, err := os.Stdin.Read(b)
			if err != nil {
				fmt.Println(err)
				return
			}

			keysPressedCh <- string(b)
		}
	}()
	return keysPressedCh
}

// Bundled returns a channel of the type string where
// each message is a string of keys that have been pressed
// within 500 milliseconds of each other.
//
// This function should only be called after the given STDIN has
// been chnaged to Raw mode via the [term] package
func Bundled() chan string {
	var keyPresses chan string = Pressed()
	keyBundles := make(chan string)
	go func() {
		var buffIt bool = false
		var keyBuffer strings.Builder
		for {
			select {
			case keyPressed := <-keyPresses: {
				if buffIt {
					keyBuffer.WriteString(keyPressed)
				} else {
					buffIt = true
					keyBuffer.WriteString(keyPressed)
					time.AfterFunc(500 * time.Millisecond, func() {
						buffIt = false
						keyBundles <- keyBuffer.String()
						keyBuffer.Reset()
					})
				}
			}
			}
		}
	}()
	return keyBundles
}
