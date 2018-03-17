package grifts

import (
	"github.com/gobuffalo/buffalo"
	"github.com/sathomsen/slidetell/actions"
)

func init() {
	buffalo.Grifts(actions.App())
}
