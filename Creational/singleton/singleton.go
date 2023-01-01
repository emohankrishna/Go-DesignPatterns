package singleton

import (
	"fmt"
	"sync"
)

var lock = &sync.Mutex{}

type Singleton struct{}

var instance *Singleton

func getInstance() *Singleton {
	if instance == nil {
		fmt.Println("Creating Singleton Instance...")
		lock.Lock()
		instance = &Singleton{}
		defer lock.Unlock()
		return instance
	} else {
		fmt.Println("Singleton Instance is already available")
	}
	return instance
}
