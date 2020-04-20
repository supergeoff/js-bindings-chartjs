package chartjs

import "github.com/gopherjs/gopherjs/js"

type Options struct {
	*js.Object
	Scales              *Scales `js:"scales"`
	Responsive          bool    `js:"responsive"`
	MaintainAspectRatio bool    `js:"maintainAspectRatio"`
	Title               *Title  `js:"title"`
}

type Scales struct {
	*js.Object
	YAxes []*Axe `js:"yAxes"`
	XAxes []*Axe `js:"xAxes"`
}

type Axe struct {
	*js.Object
	Type       string      `js:"type"`
	Ticks      *Ticks      `js:"ticks"`
	GridLines  *GridLines  `js:"gridLines"`
	ScaleLabel *ScaleLabel `js:"scaleLabel"`
}

type Ticks struct {
	*js.Object
	Display     bool `js:"display"`
	BeginAtZero bool `js:"beginAtZero"`
}

type GridLines struct {
	*js.Object
	Display         bool `js:"display"`
	DrawBorder      bool `js:"drawBorder"`
	DrawOnChartArea bool `js:"drawOnChartArea"`
	DrawTicks       bool `js:"drawTicks"`
}

type ScaleLabel struct {
	*js.Object
	Display bool `js:"display"`
}

type Title struct {
	*js.Object
	Text    string `js:"text"`
	Display bool   `js:"display"`
}

func NewOptions() *Options {
	return &Options{Object: js.Global.Get("Object").New()}
}

func (o *Options) AddScales(scales *Scales) *Options {
	o.Scales = scales
	return o
}

func (o *Options) AddTitle(title *Title) *Options {
	o.Title = title
	return o
}

func (o *Options) SetResponsive(b bool) *Options {
	o.Responsive = b
	return o
}

func (o *Options) SetMaintainAspectRatio(b bool) *Options {
	o.MaintainAspectRatio = b
	return o
}

func NewScales() *Scales {
	scales := &Scales{Object: js.Global.Get("Object").New()}
	scales.XAxes = []*Axe{}
	scales.YAxes = []*Axe{}
	return scales
}

func (s *Scales) AddXAxe(axe *Axe) *Scales {
	s.XAxes = append(s.XAxes, axe)
	return s
}

func (s *Scales) AddYAxe(axe *Axe) *Scales {
	s.YAxes = append(s.YAxes, axe)
	return s
}

func NewAxe(type_ string) *Axe {
	axe := &Axe{Object: js.Global.Get("Object").New()}
	axe.Type = type_
	return axe
}

func (a *Axe) AddTicks(ticks *Ticks) *Axe {
	a.Ticks = ticks
	return a
}

func (a *Axe) AddGridLines(gridlines *GridLines) *Axe {
	a.GridLines = gridlines
	return a
}

func (a *Axe) AddScaleLabel(scalelabel *ScaleLabel) *Axe {
	a.ScaleLabel = scalelabel
	return a
}

func NewTicks() *Ticks {
	return &Ticks{Object: js.Global.Get("Object").New()}
}

func (t *Ticks) SetDisplay(b bool) *Ticks {
	t.Display = b
	return t
}

func (t *Ticks) SetBeginAtZero(b bool) *Ticks {
	t.BeginAtZero = b
	return t
}

func NewGridLines() *GridLines {
	return &GridLines{Object: js.Global.Get("Object").New()}
}

func (g *GridLines) SetDisplay(b bool) *GridLines {
	g.Display = b
	return g
}

func (g *GridLines) SetDrawBorder(b bool) *GridLines {
	g.DrawBorder = b
	return g
}

func (g *GridLines) SetDrawOnChartArea(b bool) *GridLines {
	g.DrawOnChartArea = b
	return g
}

func (g *GridLines) SetDrawTicks(b bool) *GridLines {
	g.DrawOnChartArea = b
	return g
}

func NewScaleLabel() *ScaleLabel {
	return &ScaleLabel{Object: js.Global.Get("Object").New()}
}

func (s *ScaleLabel) SetDisplay(b bool) *ScaleLabel {
	s.Display = b
	return s
}

func NewTitle(text string) *Title {
	title := &Title{Object: js.Global.Get("Object").New()}
	title.Text = text
	return title
}

func (t *Title) SetDisplay(b bool) *Title {
	t.Display = b
	return t
}
