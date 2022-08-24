package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/raghu25/goustocks/cli"
	"github.com/raghu25/goustocks/login"
	"github.com/raghu25/goustocks/services"
)

func main() {

	cli := &cli.CLI{}

	_login := login.LoginStruct{
		Cli: cli,
	}

	fmt.Println(_login)

reAsk:
	user := _login.Login()
	fmt.Println("user:", user.Username)
	if user.IsNew {
		user = services.RegisterUser(user)
		if user == nil {
			goto reAsk
		}
	} else {
		user = services.LoginUser(user)
		if user == nil {
			fmt.Println("Login failed pls retry again")
			goto reAsk
		}
	}
	portfolio := services.GetPorfolio(user.Username)
reMenu:
	command := Menu()
	switch command {
	case 1:
		// cli.Clear()
		// stocks = svc.AddStock(stocks)
		// Boot()
	case 2:
		// stocks = svc.SellStock(stocks)
		// Boot()
	case 3:
		//services.PrintTranctions(portfolio.Stocks)
		// Boot()
	case 4:
		cli.Clear()
		services.PrintStocks(portfolio.Stocks)
		cli.ReadKey()
		goto reMenu
	case 5:
		// //cli.Clear()
		// svc.ShowPnL()
		// //cli.ReadKey()
		// Boot()
	case 6:
		fmt.Println("*******************")
		fmt.Println("Shutting down.....")
	case 0:
		cli.Clear()
		fmt.Println("Input not understood try again.")
		goto reMenu
	}
	services.SavePortFolio(portfolio)
}

func Menu() int {

	fmt.Println("Press 1 to record purches of stock")
	fmt.Println("Press 2 to record sell of stock")
	fmt.Println("Press 3 to print a stock and show all the transactions")
	fmt.Println("Press 4 to list all stocks")
	fmt.Println("Press 5 to do the P&L calc")
	fmt.Println("Press 6 to exit")

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("-> ")
	text, _ := reader.ReadString('\n')
	// convert CRLF to LF
	text = strings.Replace(text, "\n", "", -1)
	text = strings.Replace(text, "\r", "", -1)

	i, _ := strconv.Atoi(text)

	return i
}
