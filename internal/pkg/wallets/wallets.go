package wallets

import (
	"fmt"
	"sync"
)

type Wallets struct {
	balances map[uint64]float64
	mutex    sync.RWMutex
}

func NewWallets() Wallets {
	return Wallets{balances: make(map[uint64]float64)}
}

func (w *Wallets) AddWallet(clientID uint64, sum float64) string {
	w.mutex.Lock()
	res := "Wallet successfully added!"
	if _, ok := w.balances[clientID]; ok {
		res = "Unsuccessfully. The client already exists!"
	}
	w.balances[clientID] = sum
	w.mutex.Unlock()
	return res
}

func (w *Wallets) AddAmount(clientID uint64, sum float64) string {
	w.mutex.Lock()
	res := "Amount " + fmt.Sprintf("%f", sum) + " successfully added!"
	if _, ok := w.balances[clientID]; !ok || sum < 0 {
		res = "Unsuccessfully. The client doesn't exist!"
	}
	w.balances[clientID] += sum
	w.mutex.Unlock()
	return res
}

func (w *Wallets) SubtractAmount(clientID uint64, subtract float64) string {
	w.mutex.Lock()
	res := "Amount " + fmt.Sprintf("%f", subtract) + " successfully subtracted!"
	if _, ok := w.balances[clientID]; !ok || subtract < 0 {
		res = "Unsuccessfully. The client doesn't exist!"
	}
	w.balances[clientID] -= subtract
	w.mutex.Unlock()
	return res
}

func (w *Wallets) ShowBalance(clientID uint64) string {
	w.mutex.Lock()
	res := "The client doesn't exist!"
	if _, ok := w.balances[clientID]; ok {
		res = "Your balance: " + fmt.Sprintf("%f", w.balances[clientID])
	}
	w.mutex.Unlock()
	return res
}
