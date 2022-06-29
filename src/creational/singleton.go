package creational

import "sync"

type singleton struct {
	str string
}

var (
	once     sync.Once
	instance *singleton
)

const Test = "test"

func getInstance() *singleton {
	once.Do(func() {
		instance = &singleton{str: Test}
	})

	return instance
}
