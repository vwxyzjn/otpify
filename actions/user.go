package actions

import "github.com/gobuffalo/buffalo"

// UserLogin default implementation.
func UserLogin(c buffalo.Context) error {
	return c.Render(200, r.HTML("user/login.html"))
}
