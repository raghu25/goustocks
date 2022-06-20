package main

import (
	"os"
)

func main() {

	Symbol := os.Args[1]
	println("Printing stocks for ", Symbol)
	getStockPrice(Symbol)
	// argsWithProg := os.Args
	// argsWithoutProg := os.Args[1:]
	// arg := os.Args[1]
	// fmt.Println(argsWithProg)
	// fmt.Println(argsWithoutProg)
	// fmt.Println(arg)

}
