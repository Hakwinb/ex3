package main

import "time"



var timerEndTime time.Time
var timerActive bool

func timerStart(duration float64) {
	now := time.Now()
	timerEndTime = now.Add(time.Duration(duration * float64(time.Second)))
	timerActive = true
}

func timerStop() {
	timerActive = false
}

func timer_timedOut() bool {
    return timerActive  &&  time.Now().After(timerEndTime)
}

