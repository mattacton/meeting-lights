package keys

import (
	"fmt"
	"os"
)

func Pressed(keysPressedCh chan []byte) {
	for {
		b := make([]byte, 3)
		_, err := os.Stdin.Read(b)
		if err != nil {
			fmt.Println(err)
			return
		}

		keysPressedCh <- b
	}
}