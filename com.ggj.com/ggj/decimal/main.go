package main

import (
	"fmt"
	"io/ioutil"

	"github.com/shopspring/decimal"
	"github.com/skip2/go-qrcode"
)

func main() {
	//png,_:=qrcode.Encode()
	var v1 = decimal.NewFromFloat(0.871)
	var v2 = decimal.NewFromFloat(1.13)

	var v3 = v1.Add(v2)

	fmt.Println(v3)

	png, _ := qrcode.Encode("hello world", qrcode.Medium, 256)
	ioutil.WriteFile("1.png", png, 0666)

}
