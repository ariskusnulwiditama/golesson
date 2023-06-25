package mutexlesson

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

type SyncList struct {
	m     sync.Mutex
	slice []interface{}
}

func NewSyncList(cap int) *SyncList {
	return &SyncList{
		sync.Mutex{},
		make([]interface{}, cap),
	}
}

func (l *SyncList) Load(i int) interface{} {
	l.m.Lock()
	defer l.m.Unlock()
	return l.slice[i]
}

func (l *SyncList) Append(val interface{}) {
	l.m.Lock()
	defer l.m.Unlock()
	l.slice = append(l.slice, val)
}

func (l *SyncList) Store(i int, val interface{}) {
	l.m.Lock()
	defer l.m.Unlock()
	l.slice[i] = val
}

func TestMutex(t *testing.T) {
	names := []string{"Alan", "Joe", "Jack", "Ben", "Ellen", "Lisa", "Carl", "Steve", "Anton", "Popo"}
	l := NewSyncList(0)
	wg := &sync.WaitGroup{}
	wg.Add(10)

	for i := 0; i < 10; i++ {
		go func(idx int) {
			l.Append(names[idx])
			wg.Done()
		}(i)
	}
	wg.Wait()

	for i := 0; i < 10; i++ {
		fmt.Printf("Val: %v stored at idx: %d\n", l.Load(i), i)
	}
}

type BankAccount struct {
	RWMutex sync.RWMutex
	Balance int
}

func (account *BankAccount) AddBalance(amount int) {
	account.RWMutex.Lock()
	account.Balance = amount + account.Balance
	account.RWMutex.Unlock()
}

func (account *BankAccount) GetBallance() int {
	account.RWMutex.RLock()
	res := account.Balance
	account.RWMutex.RUnlock()
	return res
}

func TestRWMutex(t *testing.T) {
	acc := BankAccount{}

	for c := 0; c < 100; c++ {
		go func() {
			for k := 0; k < 100; k++ {
				acc.AddBalance(1)
				fmt.Println(acc.GetBallance())
			}
		}()
	}

	time.Sleep(5 * time.Second)
	fmt.Println("total balance: ", acc.GetBallance())
}

type UserBalance struct {
	sync.Mutex
	Name    string
	Balance int
}

func (u *UserBalance) Lock() {
	u.Mutex.Lock()
}

func (u *UserBalance) UnLock() {
	u.Mutex.Unlock()
}

func (u *UserBalance) Change(amount int) {
	u.Balance = u.Balance + amount
}

func Transfer(user1 *UserBalance, user2 *UserBalance, amuount int) {
	user1.Lock()
	fmt.Println("lock user 1", user1.Name)
	user1.Change(-amuount)
	time.Sleep(1 * time.Second)
	user2.Lock()
	fmt.Println("lock user 2", user2.Name)
	user2.Change(amuount)
	time.Sleep(1 * time.Second)

	user1.UnLock()
	user2.UnLock()
}

func TestDeadLock(t *testing.T) {
	user1 := UserBalance {
		Name: "Joko",
		Balance: 1000000,
	}

	user2 := UserBalance {
		Name: "Koko",
		Balance: 1000000,
	}

	go Transfer(&user1, &user2, 100000)
	go Transfer(&user1, &user2, 200000)
	time.Sleep(3 * time.Second)

	fmt.Println("User ", user1.Name, "balance ", user1.Balance)
	fmt.Println("User ", user2.Name, "balance ", user2.Balance)
}