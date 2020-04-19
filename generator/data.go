package generator

import (
	"math/rand"
	"time"

	"gonum.org/v1/gonum/mat"
)

// Generatedatabyrand 乱数によるデータ生成器
func Generatedatabyrand(Ndata int) NewDense {
	var datas []int
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < Ndata; i++ {
		var data int = rand.Intn(2)
		datas = append(datas, data)
	}
	return mat.NewDense(1, 1000, datas)
}
