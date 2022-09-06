package main

import (
	"avito-tech/internal/pkg/wallets"
	"fmt"
)

func main() {
	balances := wallets.NewWallets()
	fmt.Println(balances.AddWallet(0, 0))
	fmt.Println(balances.AddWallet(123, 123))
	fmt.Println(balances.ShowBalance(123))
	fmt.Println(balances.ShowBalance(50))
	fmt.Println(balances.SubtractAmount(123, 22.01))
	fmt.Println(balances.SubtractAmount(123, 0.99))
	fmt.Println(balances.SubtractAmount(123, 100))
	fmt.Println(balances.SubtractAmount(123, -100))
	fmt.Println(balances.SubtractAmount(404, 100))
	fmt.Println(balances.SubtractAmount(404, -100))
	fmt.Println(balances.ShowBalance(123))
	fmt.Println(balances.AddAmount(123, 200.99))
	fmt.Println(balances.AddAmount(123, -200))
	fmt.Println(balances.AddAmount(404, 200))
	fmt.Println(balances.AddAmount(404, -200))
	fmt.Println(balances.ShowBalance(123))
}
