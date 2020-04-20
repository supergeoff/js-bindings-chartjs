package chartjs

import "github.com/gopherjs/gopherjs/js"

type Config struct {
	*js.Object
	Type    string   `js:"type"`
	Data    *Data    `js:"data"`
	Options *Options `js:"options"`
}

func NewConfig(type_ string, data *Data, options *Options) *Config {
	config := &Config{Object: js.Global.Get("Object").New()}
	config.Type = type_
	config.Data = data
	config.Options = options
	return config
}
