package main

import "Driver-go/elevio"

type ElevatorBehaviour int

const (
	EBehIdle ElevatorBehaviour = iota
	EBehDoorOpen
	EBehMoving
)

type ClearRequestVariant int 

const (
	CVAll ClearRequestVariant = iota
	CVInDirection
)


type Elevator struct {
	Floor     int
	Direction elevio.MotorDirection
	Requests  [Nfloors][Nbuttons]int
	Behaviour ElevatorBehaviour
	ClearRV ClearRequestVariant
}

type ElevOutputDevice struct {
}

