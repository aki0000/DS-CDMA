package main

import (
	"fmt"

	"./bpskmodulator"
	"./generator"
	"gonum.org/v1/gonum/mat"
)

// Ndata number of data
const Ndata int = 1000

func main() {
	fmt.Println("------------ Transmiiter ---------------")
	// Creating binary data
	var datas *mat.Dense = generator.Generatedatabyrand(Ndata)
	// BPSK Modulartor
	var bpsk = bpskmodulator.Modulator(datas)
	fmt.Println("------------ Channel ---------------")

	fmt.Println("------------ Receiver ---------------")
	var decodedatas = bpskmodulator.Decision(bpsk)
	fmt.Println(decodedatas)

	fmt.Println("------------ Measuring BER ---------------")
	// BERの計測
	var BER float64 = MeasuringinstrumentforBER(datas, decodedatas)
	fmt.Println(BER)
}

// MeasuringinstrumentforBER BERの計測器
func MeasuringinstrumentforBER(datas *mat.Dense, decodedata *mat.Dense) float64 {
	var Ndata, Nuser = datas.Dims()
	// Formula for BER Sum(a - b)^2
	var matrix = mat.NewDense(1000, 1, nil)
	matrix.Sub(datas, decodedata)
	matrix.MulElem(matrix, matrix)
	var error = matrix.At(0, 0)
	return error / float64((Ndata * Nuser))
}
