package poker

import "fmt"

type CardPattern interface {
	fmt.Stringer
	Name() string
	Cards() Cards
	Valid() bool
	Same(s CardPattern) bool
	Equal(s CardPattern) bool
	Greeter(s CardPattern) bool
	Lesser(s CardPattern) bool
	Factory(Cards) CardPattern
}
