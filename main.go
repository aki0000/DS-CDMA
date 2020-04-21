package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"

	"./bpskmodulator"
	"./generator"
	"gonum.org/v1/gonum/mat"
)

// Ndata number of data
const Ndata int = 1000

// SNR Signal/Noise Ratio
const SNR float64 = 10 ^ (0 / 10)

// Norm norm
const Norm float64 = 1 / SNR

func main() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println("------------ Transmiiter ---------------")
	// Creating binary data
	var datas *mat.Dense = generator.Generatedatabyrand(Ndata)
	// BPSK Modulartor
	var bpsk = bpskmodulator.Modulator(datas)
	fmt.Println("------------ Channel ---------------")
	var noise = make([]float64, Ndata)
	for index := 0; index < Ndata; index++ {
		noise[index] = math.Sqrt(Norm/2) * float64(rand.Intn(Ndata))
	}
	var noisemat = mat.NewDense(Ndata, 1, noise)
	var receive = mat.NewDense(1000, 1, nil)
	receive.Copy(bpsk)
	receive.Add(receive, noisemat)
	fmt.Println("------------ Receiver ---------------")
	var decodedatas = bpskmodulator.Decision(receive)
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
