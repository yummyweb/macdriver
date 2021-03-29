package core

import "fmt"

type NSSize struct {
	Width  float64
	Height float64
}

func (sz NSSize) String() string {
	return fmt.Sprintf("(%v, %v)", sz.Width, sz.Height)
}
