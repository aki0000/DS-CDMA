package main

import (
	"fmt"

	"./bpskmodulator"
	"./generator"
	"gonum.org/v1/gonum/mat"
)

// Ndata データの数
const Ndata int = 1000

func main() {
	// 2値のデータ生成
	var datas *mat.Dense = generator.Generatedatabyrand(Ndata)
	// BPSK変調
	var bpsk = bpskmodulator.Modulator(datas)
	fmt.Println("------------ DS-CDMA start ---------------")
	fmt.Println(datas)
	// Decoder
	var decodedatas = bpskmodulator.Decision(bpsk)
	fmt.Println(decodedatas)
	fmt.Println("------------ BER 計測 ---------------")
	// BERの計測
	var a = mat.NewDense(1000, 1, nil)
	a.Sub(datas, decodedatas)
	//fmt.Println(a)
	//var c mat.Dense
	//c.Mul(a.T(), a)
	//fmt.Println(c)
	//var d = mat.Sum(c)
	//fmt.Println(d)
}
