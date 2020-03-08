package perlinnoise

import (
	"testing"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

func Test(t *testing.T) {
	p, err := plot.New()
	if err != nil {
		panic(err)
	}

	p.Title.Text = "perlin noise"
	p.X.Label.Text = "time"
	p.Y.Label.Text = "noise"

	size := 5000
	pts := make(plotter.XYs, size)

	rand := New(1000, 2000.0, 0.4, 5)

	for i := 0; i < size; i++ {
		pts[i].X = float64(i)
		pts[i].Y = rand.Generate()
	}

	err = plotutil.AddLinePoints(p, pts)
	if err != nil {
		panic(err)
	}

	if err := p.Save(10*vg.Inch, 5*vg.Inch, "perlinnoise.png"); err != nil {
		panic(err)
	}
}
