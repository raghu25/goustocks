package model

import "time"

type Portfolio struct {
	Username string `json:"_id"`
	Stocks   []Stock
}
type Stock struct {
	Symbol      string  `header:"Symbol"`
	BuyPrice    float64 `header:"BuyPrice"`
	BuyValue    float64 `header:"BuyValue"`
	BuyDate     time.Time
	BuyDateStr  string  `header:"BuyDate"`
	Qty         int     `header:"Qty"`
	SellPrice   float64 `header:"SellPrice"`
	SellValue   float64 `header:"Sell Value"`
	SellDate    time.Time
	SellDateStr string `header:"SellDate"`
	//	GrossProfite float64
	ProfitBooked float64 `header:"ProfitBooked"`
	TransType    string  `header:"TransType"`
	IsIntraDay   bool    `header:"IsIntraDay"`
	IsSold       bool    `header:"IsSold"`
	LTP          float64 `header:"LTP"`
	Transactions []Stock
	Pnl          float64 `header:"P&L"`
}
