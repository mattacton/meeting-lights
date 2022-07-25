package you

type Thees func(keys string)

type Doeet struct {
	DoWhat map[string]Thees
}

func (doeet Doeet) Now(keys string) {
	if dothees, ok := doeet.DoWhat[keys]; ok {
		dothees(keys)
	}
}