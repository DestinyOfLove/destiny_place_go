package main

import "fmt"

const boilingF = 212.0

func main() {
	fmt.Printf("%gC is %gF\n", fToC(20.0), 20.0)
	fmt.Printf("%gC is %gF\n", fToC(boilingF), boilingF)
}

func fToC(f float64) float64 {
	return (f - 32) * 5 / 9
}
