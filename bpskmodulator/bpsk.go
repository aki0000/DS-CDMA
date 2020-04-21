package bpskmodulator

import (
	"gonum.org/v1/gonum/mat"
)

// Modulator modulartor for BPSK
func Modulator(datas *mat.Dense) *mat.Dense {
	var Ndata, Nuser = datas.Dims()
	// Copy a matrix due to pointer
	var copydatas *mat.Dense = mat.NewDense(Ndata, Nuser, nil)
	copydatas.Copy(datas)
	// BPSKは [1, -1] のみの要素配列 >> 2 * datas[i] - 1
	var constarray = make([]float64, 1000)
	for index := 0; index < 1000; index++ {
		constarray[index] = 1
	}
	var constmat *mat.Dense = mat.NewDense(Ndata, Nuser, constarray)
	copydatas.Scale(2, copydatas)
	var bpsk = mat.NewDense(Ndata, Nuser, nil)
	bpsk.Sub(copydatas, constmat)
	return bpsk
}

// Decision バイナリ判定器
func Decision(recdatas *mat.Dense) *mat.Dense {
	var Ndata, _ = recdatas.Dims()
	var decdatas = make([]float64, Ndata)
	for index := 0; index < Ndata; index++ {
		var data = recdatas.At(index, 0)
		if data > 0 {
			decdatas[index] = 1
		} else {
			decdatas[index] = 0
		}
	}
	return mat.NewDense(Ndata, 1, decdatas)
}
