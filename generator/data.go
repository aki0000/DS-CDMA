package generator

import (
	"math/rand"
	"time"

	"gonum.org/v1/gonum/mat"
)

// Generatedatabyrand 乱数によるデータ生成器
func Generatedatabyrand(Ndata int) *mat.Dense {
	var datas []float64
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < Ndata; i++ {
		var data float64 = float64(rand.Intn(2))
		datas = append(datas, data)
	}
	return mat.NewDense(Ndata, 1, datas)
}
