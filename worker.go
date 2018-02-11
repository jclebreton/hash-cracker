package main

import "sync"

type PasswordComparator interface {
	SetHash(hash string)
	Compare(passwd string) bool
}

func worker(wg *sync.WaitGroup, gracefulChan chan bool, dictionaryChan chan string, ResultChan chan string, c PasswordComparator) {
	for {
		select {
		case passwd := <-dictionaryChan:
			if c.Compare(passwd) {
				ResultChan <- passwd
				wg.Done()
				return
			}
		case <-gracefulChan:
			wg.Done()
			return
		}
	}
}

func graceful(gracefulChan chan bool) {
	for i := 0; i < concurrency; i++ {
		gracefulChan <- true
	}
}
