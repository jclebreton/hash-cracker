package comparators

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_graceful_success(t *testing.T) {
	gracefulChan := make(chan bool, 3)
	defer close(gracefulChan)

	graceful(2, gracefulChan)
	gracefulChan <- false

	assert.True(t, <-gracefulChan)
	assert.True(t, <-gracefulChan)
	assert.False(t, <-gracefulChan)
}

//func Test_worker_success(t *testing.T) {
//	gracefulChan := make(chan bool, 3)
//	dictionaryChan := make(chan string, 3)
//	ResultChan := make(chan string, 3)
//	defer close(gracefulChan)
//	defer close(dictionaryChan)
//	defer close(ResultChan)
//
//	wg := &sync.WaitGroup{}
//	c := &LBCPassword{}
//	worker(wg, gracefulChan, dictionaryChan, ResultChan, c)
//}
