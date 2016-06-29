package main

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"

	"github.com/james-maloney/datatable"
	"github.com/leekchan/accounting"
)

var homeTmpl = template.Must(template.New("index").Parse(tmpl))

func main() {
	http.HandleFunc("/", HomePage)
	http.HandleFunc("/api/pie", PieChart)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

var usd = accounting.Accounting{Symbol: "$", Precision: 2}

func PieChart(w http.ResponseWriter, r *http.Request) {
	tbl := datatable.New()
	tbl.Meta = datatable.PieOptions{
		Title:  "Expenses",
		Is3D:   true,
		Legend: "left",
		Width:  500,
	}

	tbl.AddColumn(&datatable.Column{
		Type:  datatable.String,
		Label: "Item",
	}, &datatable.Column{
		Type:  datatable.Number,
		Label: "Cost",
	})

	tbl.AddRow([]*datatable.Cell{{
		Value: "Food",
	}, {
		Value:  151.5,
		Format: usd.FormatMoney(151.5),
	}}).AddRow([]*datatable.Cell{{
		Value: "Car Insurance",
	}, {
		Value:  75,
		Format: usd.FormatMoney(75),
	}}).AddRow([]*datatable.Cell{{
		Value: "Cat Food",
	}, {
		Value:  15,
		Format: usd.FormatMoney(15),
	}}).AddRow([]*datatable.Cell{{
		Value: "Mortgage",
	}, {
		Value:  1500,
		Format: usd.FormatMoney(1500),
	}})

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tbl)
}

func HomePage(w http.ResponseWriter, r *http.Request) {
	homeTmpl.Execute(w, nil)
}

var tmpl = `
<!doctype html>
<html>
	<head>
		<script src="https://code.jquery.com/jquery-3.0.0.js" integrity="sha256-jrPLZ+8vDxt2FnE1zvZXCkCcebI/C8Dt5xyaQBjxQIo=" crossorigin="anonymous"></script>
		<script type="text/javascript" src="https://www.gstatic.com/charts/loader.js"></script>
		
	  <script type="text/javascript">
		google.charts.load('current', {packages: ['corechart']});
		google.charts.setOnLoadCallback(drawChart);

		function getPieChart(cb) {
			$.ajax({
				url: "/api/pie",
				method: "GET",
				dataType: "json",
			}).done(function(data) {
				cb(data);
			});
		}

		function drawChart() {
		  getPieChart(function(data) {
			  var dt = new google.visualization.DataTable(data);
		  var chart = new google.visualization.PieChart(document.getElementById('myPieChart'));
			  chart.draw(dt, data.meta);
		  });
		}
	  </script>

	</head>
	<body>

	<div id="myPieChart"/>
	</body>
</html>
`
