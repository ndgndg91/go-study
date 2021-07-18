package goroutine

import "sync"

var mutex sync.Mutex

func DepositAndWithdrawUsingMutex(account *Account) {
	mutex.Lock()
	defer mutex.Unlock()
	DepositAndWithdraw(account)
}

func MutexDrive() {
	var wg sync.WaitGroup

	account := &Account{0}
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			for {
				DepositAndWithdraw(account)
			}
			wg.Done()
		}()
	}

	wg.Wait()
}
