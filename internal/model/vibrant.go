package vibrant

import "fmt"

type Ladybug struct {
	Name string
}

func (l *Ladybug) String() {
	fmt.Println(l.Name)
}
