package main

import (
	"context"
	"fmt"

	finnhub "github.com/Finnhub-Stock-API/finnhub-go/v2"
)

func getStockPrice(Symbol string) {
	cfg := finnhub.NewConfiguration()
	cfg.AddDefaultHeader("X-Finnhub-Token", "caml0g2ad3i5q2j8l7k0")
	finnhubClient := finnhub.NewAPIClient(cfg).DefaultApi

	res, _, err := finnhubClient.SymbolSearch(context.Background()).Q(Symbol).Execute()
	if err != nil {
		fmt.Printf("%+v\n", err)
	}
	fmt.Printf("%+v\n", res.Result)
}
