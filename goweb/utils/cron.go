package utils

import (
	"time"
)

// type MyDailyCron struct {
// 	stopChan chan bool
// }

func StartDailyCron(execHour, execMin int, execFunc func()) {
	go func() {
		for {
			now := time.Now()
			next := now.Add(time.Hour * 24)
			next = time.Date(next.Year(), next.Month(), next.Day(), execHour, execMin, 0, 0, next.Location())
			t := time.NewTimer(next.Sub(time.Now()))
			<-t.C
			execFunc()
		}
	}()
}
