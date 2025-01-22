package main

import "Driver-go/elevio"

type elevatorBehaviour int //liten bokstav??

const (
	EB_idle elevatorBehaviour = iota
	EB_doorOpen
	EB_moving
)

type elevator struct { //Hvorfor liten e
	floor     int
	direction elevio.MotorDirection
	requests  [Nfloors][Nbuttons]int
	behaviour elevatorBehaviour
}

type ElevOutputDevice struct {
}

type ClearRequestVariant int 

const (
	CVAll ClearRequestVariant = iota
	CVInMD
)
