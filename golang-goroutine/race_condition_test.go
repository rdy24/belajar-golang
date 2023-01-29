package golanggoroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestRaceCondition(t *testing.T) {
	counter := 0

	for i := 0; i < 1000; i++ {
		go func() {
			for j := 1; j < 100; j++ {
				counter++
			}
		}()
	}

	time.Sleep(5 * time.Second)
	fmt.Println("Counter", counter)
}

func TestRaceConditionMutex(t *testing.T) {
	counter := 0
	mutex := &sync.Mutex{}

	for i := 0; i < 1000; i++ {
		go func() {
			for j := 1; j <= 100; j++ {
				mutex.Lock()
				counter++
				mutex.Unlock()
			}
		}()
	}

	time.Sleep(5 * time.Second)
	fmt.Println("Counter", counter)

}

type BankAccount struct {
	RWMutex sync.RWMutex
	Balance int
}

func (account *BankAccount) AddBalance(value int) {
	account.RWMutex.Lock()
	account.Balance += value
	account.RWMutex.Unlock()
}

func (account *BankAccount) GetBalance() int {
	account.RWMutex.RLock()
	defer account.RWMutex.RUnlock()
	return account.Balance
}

func TestBankAccount(t *testing.T) {
	account := BankAccount{}

	for i := 0; i < 1000; i++ {
		go func() {
			for j := 1; j <= 100; j++ {
				account.AddBalance(1)
				fmt.Println("Balance", account.GetBalance())
			}
		}()
	}
	time.Sleep(5 * time.Second)
	fmt.Println("Total Balance = ", account.GetBalance())
}

type UserBalance struct {
	sync.Mutex
	Name   string
	Balance int
}

func (user *UserBalance) Lock() {
	user.Mutex.Lock()
}

func (user *UserBalance) Unlock() {
	user.Mutex.Unlock()
}

func (user *UserBalance) Change(amount int) {
	user.Balance += amount
}

func Transfer(user1, user2 *UserBalance, amount int) {
	user1.Lock()
	fmt.Println(user1.Name, "locked")
	user1.Change(-amount)
	
	time.Sleep(1 * time.Second)

	user2.Lock()
	fmt.Println(user2.Name, "locked")
	user2.Change(amount)

	time.Sleep(1 * time.Second)

	user1.Unlock()
	user2.Unlock()
}

func TestDeadlock(t *testing.T) {
	user1 := UserBalance{Name: "User 1", Balance: 10000}
	user2 := UserBalance{Name: "User 2", Balance: 100000}

	go Transfer(&user1, &user2, 1000)
	go Transfer(&user2, &user1, 1000)

	time.Sleep(5 * time.Second)
	fmt.Println(user1.Name, user1.Balance)
	fmt.Println(user2.Name, user2.Balance)
}