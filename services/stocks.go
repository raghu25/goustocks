package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/kataras/tablewriter"
	"github.com/lensesio/tableprinter"
	"github.com/raghu25/goustocks/model"
)

func GetPorfolio(username string) *model.Portfolio {
	url := "https://nodeapi.collabmantra.com/goportfolio/" + username
	method := "GET"
	payload := strings.NewReader(``)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return nil
	}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	if res.StatusCode == 404 {
		fmt.Println("not found")
		return &model.Portfolio{
			Username: username,
			Stocks:   []model.Stock{},
		}
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	portfolio := &model.Portfolio{}
	errparse := json.Unmarshal(body, portfolio)
	if errparse != nil {
		log.Fatalf("Error occured during unmarshaling. GetUsers Error: %s", errparse.Error())
	}

	return portfolio
}

//Create ....
func SavePortFolio(portfolio *model.Portfolio) *model.Portfolio {

	url := "https://nodeapi.collabmantra.com/goportfolio"
	method := "POST"

	//Converting Object to bytes
	payload, err := json.MarshalIndent(portfolio, "", " ")
	if err != nil {
		return nil
	}

	//Convert byte to paylod to a string to send it to server
	payloadToSend := strings.NewReader(string(payload))

	fmt.Println(string(payload))

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payloadToSend)

	if err != nil {
		fmt.Println(err)
		return nil
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	errparse := json.Unmarshal(body, portfolio)
	if errparse != nil {
		log.Fatalf("Error occured during unmarshaling. CreateUser Error: %s", errparse.Error())
	}

	return portfolio
}

func PrintTranctions(stocks *[]model.Stock) {

}

func PrintStocks(stocks []model.Stock) {
	printer := tableprinter.New(os.Stdout)

	// stocks_sorted := make(model.StocksType, 0, len(stocks))
	// for _, d := range stocks {
	// 	stocks_sorted = append(stocks_sorted, d)
	// }

	// sort.Sort(stocks_sorted)

	//Optionally, customize the table, import of the underline 'tablewriter' package is required for that.
	printer.BorderTop, printer.BorderBottom, printer.BorderLeft, printer.BorderRight = true, true, true, true
	printer.CenterSeparator = "│"
	printer.ColumnSeparator = "│"
	printer.RowSeparator = "─"
	printer.HeaderBgColor = tablewriter.BgBlackColor
	printer.HeaderFgColor = tablewriter.FgGreenColor
	printer.Print(stocks)
}
