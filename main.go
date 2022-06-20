package main

import (
	"fmt"
	"os"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("you need to pass atleast one symbol")
		fmt.Println("goustock.exe GOOG")
		panic(1)
	}
	argsWithoutProg := os.Args[1:]
	fmt.Println(argsWithoutProg)
	//Symbol := os.Args[1]
	for _, Symbol := range argsWithoutProg {
		println("Printing stocks for ", Symbol)
		//TODO : implement go ruitines here...
		getStockPrice(Symbol)
	}

	// argsWithProg := os.Args
	// argsWithoutProg := os.Args[1:]
	// arg := os.Args[1]
	// fmt.Println(argsWithProg)
	// fmt.Println(argsWithoutProg)
	// fmt.Println(arg)

}
