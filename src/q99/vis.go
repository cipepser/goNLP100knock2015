package q99

import (
	"errors"
	"image/color"
	"math"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/vg"
	"gonum.org/v1/plot/vg/draw"
)

// Point is a plotting point with drawn text.
type Point struct {
	// X and Y are the point.
	X, Y float64

	// Text is plotted name.
	Text string

	// TextStyle is the style of the plotting text.
	draw.TextStyle
}

// Scatter implements the Plotter interface, plotting
// a text for each of a set of points.
type Scatter []Point

// NewScatter creates a new Scatter for the given Tree.
func NewScatter(x, y []float64, ts []string, cs []color.Color) (*Scatter, error) {
	if len(x) != len(y) {
		return nil, errors.New("x and y must have same length")
	}
	if len(x) != len(ts) {
		return nil, errors.New("x and ts must have same length")
	}
	if len(x) != len(cs) {
		return nil, errors.New("x and cs must have same length")
	}
	font, err := vg.MakeFont(plot.DefaultFont, vg.Points(12))
	if err != nil {
		return nil, errors.New("fail to make font")
	}

	s := Scatter{}
	for i := range x {
		s = append(s, Point{
			X:    x[i],
			Y:    y[i],
			Text: ts[i],
			TextStyle: draw.TextStyle{
				Color:  cs[i],
				Font:   font,
				XAlign: draw.XCenter,
				YAlign: draw.YBottom,
			},
		})
	}

	return &s, nil
}

// Plot implements the Plot method of the plot.Plotter interface.
func (s *Scatter) Plot(c draw.Canvas, plt *plot.Plot) {
	trX, trY := plt.Transforms(&c)
	for _, p := range *s {
		c.FillText(p.TextStyle, vg.Point{X: trX(p.X), Y: trY(p.Y)}, p.Text)
	}
}

// DataRange implements the DataRange method
// of the plot.DataRanger interface.
func (s *Scatter) DataRange() (xmin, xmax, ymin, ymax float64) {
	xmin = math.Inf(1)
	xmax = math.Inf(-1)
	ymin = math.Inf(1)
	ymax = math.Inf(-1)

	for _, p := range *s {
		xmin = math.Min(xmin, p.X)
		xmax = math.Max(xmax, p.X)
		ymin = math.Min(ymin, p.Y)
		ymax = math.Max(ymax, p.Y)
	}

	return xmin * 1.1, xmax * 1.1, ymin * 1.1, ymax * 1.1
}
