package keys

import (
	"fmt"
	"os"
	"strings"
	"time"
)

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
