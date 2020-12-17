package grifts

import (
	"web_framework/goh_github/simple-go-vue-demo/actions"

	"github.com/gobuffalo/buffalo"
)

func init() {
	buffalo.Grifts(actions.App())
}
