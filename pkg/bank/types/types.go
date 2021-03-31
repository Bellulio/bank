package types

type Money int64

// Card представляет информацию о платежной карте
type Card struct {
	ID          int
	Name        string
	PAN         string
	Balance     Money
	Currency    string 
	Active 		bool 
}




type PaymentSource struct { 
	Type string // card
    Number string // номер вида '5058 xxxx xxxx 8888'
	Balance Money // баланс в дирамах
}