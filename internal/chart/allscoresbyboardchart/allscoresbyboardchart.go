package allscoresbyboardchart

import (
	"github.com/uszebr/thegamem/internal/chart/chartcache"
	"github.com/uszebr/thegamem/internal/chart/chartutil"
	"github.com/uszebr/thegamem/play/game"
	"github.com/valyala/fasttemplate"
)

type AllScoresByBoardChart struct {
	cache *chartcache.ChartCache
}

func New() *AllScoresByBoardChart {
	return &AllScoresByBoardChart{cache: chartcache.NewChartCache()}
}

func (allscoresbyboardchart *AllScoresByBoardChart) GetChartScript(game *game.Game) string {
	script, ok := allscoresbyboardchart.cache.GetCache(game.GetUUID(), game.GetBoardsQuantity())
	if ok {
		return script
	}
	calculatedScript := allscoresbyboardchart.createChartScript(game)
	allscoresbyboardchart.cache.SetCache(game.GetUUID(), game.GetBoardsQuantity(), calculatedScript)
	return calculatedScript
}

func (allscoresbyboardchart *AllScoresByBoardChart) createChartScript(game *game.Game) string {
	boards := game.GetBoards()
	data := calculateScores(game)

	t := fasttemplate.New(scriptTeplate, "<<", ">>")
	s := t.ExecuteString(map[string]interface{}{
		"dataX":       chartutil.SliceIntToString(data),
		"categoriesX": chartutil.SliceBorderToString(0, len(boards)),
	})
	return s
}

func calculateScores(game *game.Game) []int {
	boards := game.GetBoards()
	if len(boards) == 0 {
		return []int{}
	}
	resultScore := make([]int, 0, len(boards))
	for _, b := range boards {
		resultScore = append(resultScore, b.GetBoardScoresSum())
	}
	return resultScore
}

const scriptTeplate string = `
<script>
 var options = {
          series: [{
          name: 'Board Score',
          data: <<dataX>>
        }],
          chart: {
          height: 350,
          type: 'bar',
        },
        plotOptions: {
          bar: {
            borderRadius: 2,
            dataLabels: {
              position: 'top', // top, center, bottom
            },
          }
        },
        dataLabels: {
          enabled: false,
          offsetY: -20,
          style: {
            fontSize: '12px',
            colors: ["#304758"]
          }
        },
        
        xaxis: {
          categories: <<categoriesX>>,
          position: 'bottom',
          axisBorder: {
            show: true
          },
          axisTicks: {
            show: false
          },
          crosshairs: {
            fill: {
              type: 'gradient',
              gradient: {
                colorFrom: '#D8E3F0',
                colorTo: '#BED1E6',
                stops: [0, 100],
                opacityFrom: 0.4,
                opacityTo: 0.5,
              }
            }
          },
          tooltip: {
            enabled: true,
          }
        },
        yaxis: {
          axisBorder: {
            show: true,
          },
          axisTicks: {
            show: true,
          },
          labels: {
            show: true,
          }
        
        },
        title: {
          text: 'Sum of all scores on boards',
          floating: true,
          offsetY: 330,
          align: 'center',
          style: {
            color: '#444'
          }
        }
        };

        var chart = new ApexCharts(document.querySelector("#board-scores"), options);
        chart.render();
</script>
`
