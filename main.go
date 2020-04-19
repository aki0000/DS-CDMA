package main

import (
	"fmt"

	"./generator"
)

// Ndata データの数
const Ndata int = 1000

func main() {
	// 2値のデータ生成
	var datas = generator.Generatedatabyrand(Ndata)
	fmt.Println(datas)
	// BPSK変調

	fmt.Println("------------ DS-CDMA start ---------------")
}
