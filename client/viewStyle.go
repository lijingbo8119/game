package client

import (
	"os"

	"golang.org/x/term"
)

type viewStyle struct{}

func (r viewStyle) PhysicalWidth() int {
	physicalWidth, _, _ := term.GetSize(int(os.Stdout.Fd()))
	return physicalWidth
}
