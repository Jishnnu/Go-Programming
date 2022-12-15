package main

import (
	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"os"
)

func main() {
	sals := []opts.LineData{{Value: 567}, {Value: 612}, {Value: 800}, {Value: 980}, {Value: 1410}, {Value: 2350}}
	ages := []string{"18", "20", "25", "30", "40", "50"}
	line := charts.NewLine()

	line.setGlobalOptions(
		charts.WithTitleOpts(opts.Title{Title: "AVERAGE SALARAY PER AGE", Subtitle: "ABC COMPANY"}),
	)
	line.setAxis(ages)
	line.AddSeries("SALARIES", sals)
	f, _ := os.Create("line.html")
	line.Render(f)
}
