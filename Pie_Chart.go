package main

import (
	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"os"
)

func main() {
	dest := []opts.PieData{{Name: "Croatia", Value: 22}, {Name: "Bohemia", Value: 34}, {Name: "Bulgaria", Value: 18}, {Name: "Spain", Value: 10}}
	pie := charts.NewPie()

	pie.setGlobalOptions(
		charts.WithTitleOpts(opts.Title{Title: "Popular Destinations"}),
	)

	pie.AddSeries("Destinations", dest)

	f, _ := os.Create("Pie.html")
	pie.render(f)
}
