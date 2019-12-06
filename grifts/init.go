package grifts

import (
	"github.com/dkeza/goblog/actions"

	"github.com/gobuffalo/buffalo"
)

func init() {
	buffalo.Grifts(actions.App())
}
