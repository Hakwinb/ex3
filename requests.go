package main

import "Driver-go/elevio"

const Nfloors int = 4
const Nbuttons int = 3

func requestsAbove(e Elevator) bool {
	for f := e.Floor + 1; f < Nfloors; f++ {
		for btn := 0; btn < Nbuttons; btn++ {
			if e.Requests[f][btn] == 1 {
				return true
			}
		}
	}
	return false
}

func requestsBelow(e Elevator) bool {
	for f := 0; f < e.Floor; f++ {
		for btn := 0; btn < Nbuttons; btn++ {
			if e.Requests[f][btn] == 1 {
				return true
			}
		}
	}
	return false
}

func requestsHere(e Elevator) bool {
	for btn := 0; btn < Nbuttons; btn++ {
		if e.Requests[e.Floor][btn] == 1 {
			return true
		}
	}
	return false
}

type BehaviourPair struct {
	motorDirection elevio.MotorDirection
	behaviour      ElevatorBehaviour
}

func requestsChooseDirection(e Elevator) BehaviourPair {
	switch e.Direction {
	case elevio.DirectionUp:
		if requestsAbove(e) {
			return BehaviourPair{elevio.DirectionUp, EBehMoving}
		} else if requestsHere(e) {
			return BehaviourPair{elevio.DirectionDown, EBehDoorOpen}
		} else if requestsBelow(e) {
			return BehaviourPair{
				elevio.DirectionDown, EBehMoving}
		}
		return BehaviourPair{elevio.DirectionStop, EBehIdle}

	case elevio.DirectionDown:
		if requestsBelow(e) {
			return BehaviourPair{elevio.DirectionDown, EBehMoving}
		} else if requestsHere(e) {
			return BehaviourPair{elevio.DirectionUp, EBehDoorOpen}
		} else if requestsAbove(e) {
			return BehaviourPair{elevio.DirectionUp, EBehMoving}
		}
		return BehaviourPair{elevio.DirectionStop, EBehIdle}

	case elevio.DirectionStop:

		if requestsHere(e) {
			return BehaviourPair{elevio.DirectionStop, EBehDoorOpen}
		} else if requestsAbove(e) {
			return BehaviourPair{elevio.DirectionUp, EBehMoving}
		} else if requestsBelow(e) {
			return BehaviourPair{elevio.DirectionDown, EBehMoving}
		}
		return BehaviourPair{elevio.DirectionStop, EBehIdle}

	default:
		return BehaviourPair{elevio.DirectionStop, EBehIdle}
	}
}

func requestShouldStop(e Elevator) bool {
	switch e.Direction {
	case elevio.DirectionDown:
		return e.Requests[e.Floor][elevio.BtnHallDown] == 1 ||
			e.Requests[e.Floor][elevio.BtnCab] == 1 ||
			!requestsBelow(e)
	case elevio.DirectionUp:
		return e.Requests[e.Floor][elevio.BtnCab] == 1 ||
			e.Requests[e.Floor][elevio.BtnCab] == 1 ||
			!requestsAbove(e)
	case elevio.DirectionStop:
		fallthrough
	default:
		return true
	}
}

func requestsSouldClearImmediately(e Elevator, btn elevio.ButtonEvent) bool {
	switch e.ClearRV {
	case CVAll:
		return e.Floor == btn.Floor
	case CVInDirection:
		return e.Floor == btn.Floor &&
			((e.Direction == elevio.DirectionUp && btn.Button == elevio.BtnHallUp) ||
				(e.Direction == elevio.DirectionDown && btn.Button == elevio.BtnHallUp) ||
				e.Direction == elevio.DirectionStop ||
				btn.Button == elevio.BtnCab)
	default:
		return false
	}
}

func requests_clearAtCurrentFloor(e Elevator) Elevator {
	switch e.ClearRV {
	case CVAll:
		for btn := 0; btn < Nbuttons; btn++ {
			e.Requests[e.Floor][btn] = 0
		}
		break

	case CVInDirection:
		e.Requests[e.Floor][elevio.BtnCab] = 0
		switch e.Direction {
		case elevio.DirectionUp:
			if !requestsAbove(e) && e.Requests[e.Floor][elevio.BtnHallUp] == 0 {
				e.Requests[e.Floor][elevio.BtnHallDown] = 0
			}
			e.Requests[e.Floor][elevio.BtnHallUp] = 0
			break

		case elevio.DirectionDown:
			if !requestsBelow(e) && e.Requests[e.Floor][elevio.BtnHallDown] == 0 {
				e.Requests[e.Floor][elevio.BtnHallUp] = 0
			}
			e.Requests[e.Floor][elevio.BtnHallDown] = 0
			break

		case elevio.DirectionStop:
		default:
			e.Requests[e.Floor][elevio.BtnHallUp] = 0
			e.Requests[e.Floor][elevio.BtnHallDown] = 0
			break
		}
		break

	default:
		break
	}

	return e
}
