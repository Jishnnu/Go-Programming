package main

import (
	"fmt"
	"net/http"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
)

func httpserver(w http.ResponseWriter, _ *http.Request) {
	population := []opts.PieData{{Name: "Uttar Pradesh", Value: 199812341}, {Name: "Kerala", Value: 33387677}, {Name: "Karnataka", Value: 61130704}, {Name: "Sikkim", Value: 607688}}

	pieChart := charts.NewPie()
	pieChart.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{Title: "Population of Various States - Pie Chart", Subtitle: "Rendered by the http server"}))
	pieChart.AddSeries("population", population)
	pieChart.Render(w)
}

func main() {
	http.HandleFunc("/", httpserver)
	fmt.Println("Server started at port 8080")
	http.ListenAndServe(":8080", nil)
}
