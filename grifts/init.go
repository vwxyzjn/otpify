package grifts

import (
	"github.com/gobuffalo/buffalo"
	"github.com/vwxyzjn/otpify/actions"
)

func init() {
	buffalo.Grifts(actions.App())
}
