package modelsbyboardchart

import (
	"github.com/uszebr/thegamem/internal/chart/chartutil"
	"github.com/uszebr/thegamem/play/game"
	"github.com/valyala/fasttemplate"
)

type ModelsByBoardChart struct{}

type seriesModelQuantity struct {
	Name string
	Data []int
}

func (modelsByBoard ModelsByBoardChart) GetChartScript(game *game.Game) string {
	//todo create cashing sytem and return value if cashed
	return modelsByBoard.createChartScript(game)
}

func (modelsByBoard ModelsByBoardChart) createChartScript(game *game.Game) string {
	boards := game.GetBoards()
	seriesModelQuantities := calculateSeriesModelQuanteties(game)
	t := fasttemplate.New(scriptTeplate, "<<", ">>")
	s := t.ExecuteString(map[string]interface{}{
		"seriesData":  seriesData(seriesModelQuantities),
		"categoriesX": chartutil.SliceBorderToString(0, len(boards)),
	})
	return s
}

// calculateing model quantity data based on current game
//
//	ex. return: []seriesModelQuantity{
//		{"alwaysgreen", []int{10, 41, 35, 51, 49, 62, 69, 91, 148}},
//		{"alwaysred", []int{44, 51, 25, 21, 29, 72, 65, 71, 48}},
//	}
func calculateSeriesModelQuanteties(game *game.Game) []seriesModelQuantity {
	initialModels := game.GetInitialModels()
	// calculating result as map
	result := make([]seriesModelQuantity, 0, len(initialModels))
	//initial structure, all models with empty slice
	for _, model := range initialModels {
		result = append(result, seriesModelQuantity{Name: model, Data: []int{}})
	}

	for _, b := range game.GetBoards() {
		resultForBoard := make(map[string]int, len(initialModels))
		for _, p := range b.GetPlayersOneSlice() {
			playerModel := p.GetModelName()
			resultForBoard[playerModel]++
		}

		// adding results from board to the result
		for i := range result {
			quantityForBoard := resultForBoard[result[i].Name]
			result[i].Data = append(result[i].Data, quantityForBoard)
		}

	}
	return result
}

// converting series data to js data
func seriesData(modelsQuantities []seriesModelQuantity) string {
	result := ""
	t := fasttemplate.New(seriesOneTemplate, "<<", ">>")
	for _, modelNameQuantity := range modelsQuantities {
		result += t.ExecuteString(map[string]interface{}{
			"name": modelNameQuantity.Name,
			"data": chartutil.SliceIntToString(modelNameQuantity.Data),
		})
	}
	return result
}

const seriesOneTemplate string = `{name:"<<name>>",data:<<data>>},`
const scriptTeplate string = `
<script>
    var options = {
        series: [
         <<seriesData>>
        ],
        chart: {
            height: 350,
            type: 'line',
            zoom: {
                enabled: false
            }
        },
        dataLabels: {
            enabled: false
        },
        stroke: {
            curve: 'smooth'
        },
        title: {
            text: 'Models Quantity By Boards',
            align: 'left'
        },
        grid: {
            borderColor: '#ECECEC',
            xaxis: {
            lines: {
                    show: true,
                }
            },
            yaxis: {
                lines: {
                show: true,
                }
            },
            },
        xaxis: {
			title: {
				text: 'Board'
			},
            categories: <<categoriesX>>,
        },
 		yaxis: {
			title: {
				text: 'Players quantity'
			}
        }
    };

    var chart = new ApexCharts(document.querySelector("#chart"), options);
    chart.render();
	</script>
`
