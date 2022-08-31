package main

var wallets = map[int]float32{
	0: 0,
}

func addWallet(clientID int, sum float32) bool {
	if _, ok := wallets[clientID]; ok {
		return false
	}
	wallets[clientID] = sum
	return true
}

func addAmount(clientID int, sum float32) bool {
	if _, ok := wallets[clientID]; !ok || sum < 0 {
		return false
	}
	wallets[clientID] += sum
	return true
}

func subtractAmount(clientID int, subtract float32) bool {
	if _, ok := wallets[clientID]; !ok || subtract < 0 {
		return false
	}
	wallets[clientID] -= subtract
	return true
}

func showBalance(clientID int) float32 {
	if _, ok := wallets[clientID]; !ok {
		return -1
	}
	return wallets[clientID]
}

func main() {
	println()
	println(addWallet(123, 123))
	println(addWallet(123, 123))
	println(showBalance(123))
	println(showBalance(50))
	println(subtractAmount(123, 22.01))
	println(subtractAmount(123, 0.99))
	println(subtractAmount(123, 100))
	println(subtractAmount(123, -100))
	println(subtractAmount(404, 100))
	println(subtractAmount(404, -100))
	println(showBalance(123))
	println(addAmount(123, 200.99))
	println(addAmount(123, -200))
	println(addAmount(404, 200))
	println(addAmount(404, -200))
	println(showBalance(123))
}
