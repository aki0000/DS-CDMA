package bpskmodulator

import (
	"gonum.org/v1/gonum/mat"
)

// Modulator BPSK変調器
func Modulator(datas *mat.Dense) *mat.Dense {
	// BPSKは [1, -1] のみの要素配列 >> 2 * datas[i] - 1
	var constarray = make([]float64, 1000)
	for index := 0; index < 1000; index++ {
		constarray[index] = 1
	}
	var constmat *mat.Dense = mat.NewDense(1000, 1, constarray)
	var codedatas = datas
	codedatas.Scale(2, codedatas)
	var bpsk = mat.NewDense(1000, 1, nil)
	bpsk.Sub(codedatas, constmat)
	return bpsk
}

// Decision バイナリ判定器
func Decision(recdatas *mat.Dense) *mat.Dense {
	var decdatas = make([]float64, 1000)
	for index := 0; index < 1000; index++ {
		var data = recdatas.At(index, 0)
		if data > 0 {
			decdatas[index] = 1
		} else {
			decdatas[index] = 0
		}
	}
	return mat.NewDense(1000, 1, decdatas)
}
