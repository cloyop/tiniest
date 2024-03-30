package utils

import (
	"time"
)

var ipChannel = make(chan func())

func RunIpQuequeChannel() {
	var waiter int
	lastPinged := time.AfterFunc(time.Second*1, func() {
		waiter = 25
	})
	go func() {
		for VisitToRegist := range ipChannel {
			if waiter == 0 {
				time.Sleep(45 * time.Second)
				waiter = 25
			}
			lastPinged.Reset(time.Second * 15)
			VisitToRegist()
			waiter--
		}
	}()
}

func PassToIpChan(f func()) {
	ipChannel <- f
}
