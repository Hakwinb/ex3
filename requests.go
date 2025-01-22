package main

import "Driver-go/elevio"

const Nfloors int = 4
const Nbuttons int = 3

func requestsAbove(e elevator) bool {
	for f := e.floor + 1; f < Nfloors; f++ {
		for btn := 0; btn < Nbuttons; btn++ {
			if e.requests[f][btn] == 1 {
				return true
			}
		}
	}
	return false
}

func requestsBelow(e elevator) bool {
	for f := 0; f < e.floor; f++ {
		for btn := 0; btn < Nbuttons; btn++ {
			if e.requests[f][btn] == 1 {
				return true
			}
		}
	}
	return false
}

func requestsHere(e elevator) bool {
	for btn := 0; btn < Nbuttons; btn++ {
		if e.requests[e.floor][btn] == 1 {
			return true
		}
	}
	return false
}

type behaviourPair struct {
	motorDirection elevio.MotorDirection
	behaviour      elevatorBehaviour
}

func requestsChooseDirection(e elevator) behaviourPair {
	switch e.direction {
	case elevio.MD_Up:
		if requestsAbove(e) {
			return behaviourPair{elevio.MD_Up, EB_moving}
		} else if requestsHere(e) {
			return behaviourPair{elevio.MD_Down, EB_doorOpen}
		} else if requestsBelow(e) {
			return behaviourPair{elevio.MD_Down, EB_moving}
		}
		return behaviourPair{elevio.MD_Stop, EB_idle}

	case elevio.MD_Down:
		if requestsBelow(e) {
			return behaviourPair{elevio.MD_Down, EB_moving}
		} else if requestsHere(e) {
			return behaviourPair{elevio.MD_Up, EB_doorOpen}
		} else if requestsAbove(e) {
			return behaviourPair{elevio.MD_Up, EB_moving}
		}
		return behaviourPair{elevio.MD_Stop, EB_idle}

	case elevio.MD_Stop:

		if requestsHere(e) {
			return behaviourPair{elevio.MD_Stop, EB_doorOpen}
		} else if requestsAbove(e) {
			return behaviourPair{elevio.MD_Up, EB_moving}
		} else if requestsBelow(e) {
			return behaviourPair{elevio.MD_Down, EB_moving}
		}
		return behaviourPair{elevio.MD_Stop, EB_idle}

	default:
		return behaviourPair{elevio.MD_Stop, EB_idle}
	}
}

func requestShouldStop(e elevator) bool{
    switch e.direction {
    case elevio.MD_Down:
        return  e.requests[e.floor][elevio.BT_HallDown] == 1 ||
				e.requests[e.floor][elevio.BT_Cab] == 1 	 ||
            	!requestsBelow(e);
    case elevio.MD_Up:
        return  e.requests[e.floor][elevio.BT_Cab] == 1 ||
            	e.requests[e.floor][elevio.BT_Cab] == 1 ||
            	!requestsAbove(e);
    case elevio.MD_Stop:
		fallthrough
    default:
        return true
    }
}

func requests_shouldClearImmediately(e elevator, elevio.ButtonEvent.floor int, btn elevio.ButtonType){
    switch(e.config.clearRequestVariant){
    case CV_All:
        return e.floor == btn_floor;
    case CV_InDirn:
        return 
            e.floor == btn_floor && 
            (
                (e.dirn == D_Up   && btn_type == B_HallUp)    ||
                (e.dirn == D_Down && btn_type == B_HallDown)  ||
                e.dirn == D_Stop ||
                btn_type == B_Cab
            );  
    default:
        return 0;
    }
}


