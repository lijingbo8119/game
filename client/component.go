package client

import "fmt"

type component interface {
	fmt.Stringer
	init() component
}
