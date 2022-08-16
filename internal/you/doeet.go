// Package you implements a means for executing functions
// based on a matching string
//
// This package is funnier if read in an Arnold Schwarzeneger
package you

type Doeet struct {
	DoWhat map[string]func(keys string)
}

func (doeet Doeet) Now(keys string) {
	if dotheesNow, ok := doeet.DoWhat[keys]; ok {
		dotheesNow(keys)
	}
}
