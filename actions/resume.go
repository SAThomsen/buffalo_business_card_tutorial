package actions

import (
	"github.com/gobuffalo/buffalo"
)

func ResumeHandler(c buffalo.Context) error {
	return c.Render(200, r.HTML("resume.html"))
}
