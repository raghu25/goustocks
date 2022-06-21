package main

import (
	"fmt"
	"os"

	"github.com/raghu25/goustocks/login"
)

func main() {

	user := login.Login()
	fmt.Println(user)
	if len(os.Args) < 2 {
		fmt.Println("you need to pass atleast one symbol")
		fmt.Println("goustock.exe GOOG")
		panic(1)
	}
	//argsWithoutProg := os.Args[1:]

	// for i := 0; i < len(argsWithoutProg); i++ {
	// 	fmt.Println(argsWithoutProg[i])

	// }

	// for _, Symbol := range argsWithoutProg {
	// 	println("Printing stocks for ", Symbol)
	// 	//TODO : implement go ruitines here...
	// 	getStockPrice(Symbol)
	// }

	//	const name string = "darshan"
	var name string = "darshan"
	name = "vijay"
	fmt.Print(name)

	// argsWithProg := os.Args
	// argsWithoutProg := os.Args[1:]
	// arg := os.Args[1]
	// fmt.Println(argsWithProg)
	// fmt.Println(argsWithoutProg)
	// fmt.Println(arg)

}
