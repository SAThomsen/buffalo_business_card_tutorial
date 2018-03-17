package actions

import (
	"github.com/gobuffalo/buffalo"
)

func ContactHandler(c buffalo.Context) error {
	return c.Render(200, r.HTML("contact.html"))
}
