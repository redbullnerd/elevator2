
// Gruppe12, Knut Hvamb & Christopher Benjamin Westlye, NTNU spring 2013

package elevator

import "elevdriver"
import "fmt"
// import "time"

type Event int

const (
	ORDER Event = iota
	STOP
	OBSTRUCTION
	SENSOR
	NO_EVENT
)

func (elevinf *Elevatorinfo) SetEvent(){
	currentFloor := elevdriver.GetFloor()
	switch{ 
	case elevdriver.GetStopButton() && elevinf.state != EMERGENCY:
		elevinf.event = STOP
	case elevdriver.GetObs():
		elevinf.event = OBSTRUCTION
	case elevinf.DetermineDirection() != 2 && elevinf.state != ASCENDING && elevinf.state != DECENDING:
		elevinf.event = ORDER
	case currentFloor != -1:
		elevinf.event = SENSOR
		elevinf.last_floor = currentFloor
	default:
		elevinf.event = NO_EVENT
	}
	fmt.Printf("updated event\n")
}
