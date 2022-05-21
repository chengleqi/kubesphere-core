package controllers

import (
	"sync"
)

func listAndWatch(ctl Controller, stopChan chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()
	go ctl.sync(stopChan)
	for {
		select {
		case <-stopChan:
			return
		default:
		}
	}
}
