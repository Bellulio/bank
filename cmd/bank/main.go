package main
import (
	
	"bank/pkg/bank/types"
)


var cards []types.Card


func main() {
	cards = []types.Card {
		{
			PAN: "5058 xxxx xxxx 7777",
			Balance: 777_77,
			Active: true,
		},
		{
			PAN: "5058 xxxx xxxx 8888",
			Balance: 888_88,
			Active: true,
		},
		{
			PAN: "5058 xxxx xxxx 9999",
			Balance: 999_99,
			Active: true,
		},
	} 
}

