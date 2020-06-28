package charts

import (
	"github.com/go-echarts/go-echarts/opts"
	"github.com/go-echarts/go-echarts/render"
	"github.com/go-echarts/go-echarts/types"
)

// ThemeRiver represents a theme river chart.
type ThemeRiver struct {
	BaseConfiguration
	MultiSeries
}

// Type returns the chart type.
func (ThemeRiver) Type() string { return types.ChartThemeRiver }

// NewThemeRiver creates a new theme river chart.
func NewThemeRiver() *ThemeRiver {
	chart := &ThemeRiver{}
	chart.initBaseConfiguration()
	chart.Renderer = render.NewChartRender(chart, chart.Validate)
	chart.HasSingleAxis = true
	return chart
}

// AddSeries adds new data sets.
func (c *ThemeRiver) AddSeries(name string, data []opts.ThemeRiverData, opts ...SeriesOpts) *ThemeRiver {
	series := SingleSeries{Name: name, Type: types.ChartThemeRiver, Data: data}
	series.configureSeriesOpts(opts...)
	c.MultiSeries = append(c.MultiSeries, series)
	return c
}

// SetGlobalOptions sets options for the ThemeRiver instance.
func (c *ThemeRiver) SetGlobalOptions(opts ...GlobalOpts) *ThemeRiver {
	c.BaseConfiguration.setBaseGlobalOptions(opts...)
	return c
}

// Validate
func (c *ThemeRiver) Validate() {
	c.Assets.Validate(c.AssetsHost)
}

// Render renders the chart and writes the output to given writer.
//func (c *ThemeRiver) Render(w io.Writer) error {
//	c.Validate()
//	return render.ChartRender(c, w)
//}
