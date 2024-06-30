package modeldistributionchart

import (
	"github.com/uszebr/thegamem/internal/chart/chartutil"
	"github.com/uszebr/thegamem/play/board"
	"github.com/uszebr/thegamem/play/game"
	"github.com/valyala/fasttemplate"
)

type ModelsDistributionChart struct{}

type seriesModelDistribution struct {
	series []int
	labels []string
}

func (modelsDistributionChart ModelsDistributionChart) GetChartScript(game *game.Game) string {
	//todo create cashing sytem and return value if cashed
	return modelsDistributionChart.createChartScript(game)
}

func (modelsDistributionChart ModelsDistributionChart) createChartScript(game *game.Game) string {
	boards := game.GetBoards()
	boardsQuantity := len(boards)
	var md seriesModelDistribution
	if boardsQuantity == 0 {
		md = seriesModelDistribution{[]int{}, []string{}}
	} else {
		md = calculateModelDistribution(boards[boardsQuantity-1])
	}

	t := fasttemplate.New(scriptTemplate, "<<", ">>")
	s := t.ExecuteString(map[string]interface{}{
		"seriesData": chartutil.SliceIntToString(md.series),
		"labelsData": chartutil.SliceStringToString(md.labels),
	})
	return s
}

func calculateModelDistribution(b *board.Board) seriesModelDistribution {
	resultForBoard := make(map[string]int)
	for _, p := range b.GetPlayersOneSlice() {
		playerModel := p.GetModelName()
		resultForBoard[playerModel]++
	}
	result := seriesModelDistribution{}
	for key, value := range resultForBoard {
		result.labels = append(result.labels, key)
		result.series = append(result.series, value)
	}
	return result
}

const scriptTemplate string = `
<script>
    var options = {
          series: <<seriesData>>,
          chart: {
          height: 350,
          type: 'pie',
		  toolbar: {
                show: true,
                tools: {
                    download: true,
                }
            }
        },
		 title: {
            text: 'Models Distribution On The Last  Board',
            align: 'left'
        },
        labels: <<labelsData>>,
        };

    var chart = new ApexCharts(document.querySelector("#model-distribution"), options);
    chart.render();
</script>
`
