package main

//variables to track money transfer
//bank and cash are separately
var bankBalance float64
var cashBalance float64
//one big counter for all money
var moneyBalance float64 = cashBalance+bankBalance

func main() {

}

//function to add money to the bank account
func addMoneyBank(amount float64){
	bankBalance=bankBalance+amount
}
//function to add money to the cash counter
func addMoneyCash(amount float64){
	cashBalance=cashBalance+amount
}
//refresh the money counter
func refreshMoney(){
	moneyBalance=cashBalance+bankBalance
}
