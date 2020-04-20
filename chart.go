package chartjs

import (
	"github.com/gopherjs/gopherjs/js"
)

type Chart struct {
	*js.Object
}

func NewChart(ctx *js.Object, config *Config) *Chart {
	return &Chart{Object: js.Global.Get("Chart").New(ctx, config)}
}

func GetChart(name string) *Chart {
	return &Chart{Object: js.Global.Get("document").Get(name)}
}

func (c *Chart) Destroy() {
	c.Call("destroy", nil)
}
