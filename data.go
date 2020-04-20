package chartjs

import "github.com/gopherjs/gopherjs/js"

type Data struct {
	*js.Object
	Labels   []string   `js:"labels"`
	Datasets []*Dataset `js:"datasets"`
}

type Dataset struct {
	*js.Object
	Label           string        `js:"label"`
	Data            []interface{} `js:"data"`
	BackgroundColor []string      `js:"backgroundColor"`
	BorderColor     []string      `js:"borderColor"`
	BorderWidth     int           `js:"borderWidth"`
}

func NewData() *Data {
	data := &Data{Object: js.Global.Get("Object").New()}
	data.Datasets = []*Dataset{}
	return data
}

func (d *Data) SetLabels(labels []string) *Data {
	d.Labels = labels
	return d
}

func (d *Data) AddDataset(dataset *Dataset) *Data {
	d.Datasets = append(d.Datasets, dataset)
	return d
}

func NewDataset(label string, data []interface{}) *Dataset {
	dataset := &Dataset{Object: js.Global.Get("Object").New()}
	dataset.Label = label
	dataset.Data = data
	return dataset
}

func (ds *Dataset) SetBackgroundColor(colors []string) *Dataset {
	ds.BackgroundColor = colors
	return ds
}

func (ds *Dataset) SetBorderColor(colors []string) *Dataset {
	ds.BorderColor = colors
	return ds
}

func (ds *Dataset) SetBorderWidth(width int) *Dataset {
	ds.BorderWidth = width
	return ds
}
