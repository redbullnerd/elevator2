
// Gruppe12, Knut Hvamb & Christopher Benjamin Westlye, NTNU spring 2013

package elevator

import "elevdriver"
import "fmt"
// import "time"

var floor_button int
var direction_button int
		
func (elevinf *Elevatorinfo) ReceiveOrders (){

	floorbutton, directionbutton := elevdriver.GetButton()
	
	// Skal ideelt sett fjerne denne if'en, andre kan ta seg av eksterne ordre...
	// if elevinf.state != EMERGENCY || (elevinf.state == EMERGENCY || elevinf.event == ORDER) {
	for i := 1; i < 4; i++ { // First column of the order slice refers to UP buttons
		if i == floorbutton && directionbutton == 1 {
			fmt.Printf("Someone wants up...\n")
			elevinf.external_orders[i-1][0] = 1
		}
	}
	for i := 2; i < 5; i++ { // Second column of the order slice refers to DOWN buttons
		if i == floorbutton && directionbutton == 2 {
			fmt.Printf("Someone wants down...\n")
			elevinf.external_orders[i-1][1] = 1
		}
	}
	// }
	for i := 1; i < 5; i++ { // Third column of the order slice refers to COMMAND buttons
		if i == floorbutton && directionbutton == 0 {
			elevinf.internal_orders[i-1][2] = 1
		}
	}
	// Clearing the unused spaces
	elevinf.internal_orders[3][0] = 0
	elevinf.internal_orders[0][1] = 0
	elevinf.external_orders[3][0] = 0
	elevinf.external_orders[0][1] = 0

}

func (elevinf *Elevatorinfo) StopAtCurrentFloor()(int){

	current := elevinf.last_floor
	fmt.Printf("StopAtCurrentFloor engaging\n")
	if elevinf.state == ASCENDING {
		fmt.Printf("ASCENDING detected: Should i stop?\n")
		for i := 0; i < 3; i = i+2 {
			if current == 1 && elevinf.internal_orders[current-1][i] == 1 {
				return 1
			} else if current == 2 && elevinf.internal_orders[current-1][i] == 1 {
				return 1
			} else if current == 3 && elevinf.internal_orders[current-1][i] == 1 {
				return 1
			} else if current == 4 && elevinf.internal_orders[current-1][i] == 1 {
				return 1
			}
		}
		orders_above_current := 0
		for i := current; i < 4; i++ {
			for j := 0; j < 3; j++ {
				if elevinf.internal_orders[i][j] == 1 {
					orders_above_current++
				}
			}
		}
		if elevinf.internal_orders[current-1][1] == 1 && orders_above_current == 0 {
			return 1
		}
		if current == 4 && elevinf.internal_orders[3][1] == 1 {
			return 1	
		}
	} else if elevinf.state == DECENDING {
		fmt.Printf("DECENDING detected: Should i stop?\n")
		for i := 1; i < 3 ; i++ {
			if current == 1 && elevinf.internal_orders[current-1][i] == 1 {
				return -1
			} else if current == 2 && elevinf.internal_orders[current-1][i] == 1 {
				return -1
			} else if current == 3 && elevinf.internal_orders[current-1][i] == 1 {
				return -1
			} else if current == 4 && elevinf.internal_orders[current-1][i] == 1 {
				return -1
			}
		}
		if current == 1 && elevinf.internal_orders[0][0] == 1 {
			return -1
		}
		orders_below_current := 0
		for i := 0; i < current-1; i++ {
			for j := 0; j < 3; j++ {
				if elevinf.internal_orders[i][j] == 1 {
					orders_below_current++
				}
			}
		}
		if elevinf.internal_orders[current-1][0] == 1 && orders_below_current == 0{
			return -1
		}
	} else if elevinf.state == EMERGENCY {
	fmt.Printf("Emergency shouldistop case\n")
		for i := 0; i < 3; i++ {
			if current == 1 && elevinf.internal_orders[0][i] == 1{
				return 2
			} else if current == 2 && elevinf.internal_orders[1][i] == 1{
				return 2
			} else if current == 3 && elevinf.internal_orders[2][i] == 1{
				return 2
			} else if current == 4 && elevinf.internal_orders[3][i] == 1{
				return 2
			}
		}
	}
	
	return 0
	
}

func (elevinf *Elevatorinfo) DeleteOrders(){

	fmt.Printf("gonna delete...\n")
	if elevdriver.GetFloor() == 1{
		for i := 0; i < 3; i++ {
			elevinf.internal_orders[0][i] = 0
			elevinf.external_orders[0][0] = 0
			go ExternalSendDelete(0,0)
			
		}
	} else if elevdriver.GetFloor() == 2 && elevinf.last_direction == UP{
		for i := 0; i < 3; i = i+2 {
			elevinf.internal_orders[1][i] = 0
			elevinf.external_orders[1][0] = 0
			if elevinf.NoOrdersAbove() {
				elevinf.external_orders[1][1] = 0
				elevinf.internal_orders[1][1] = 0
				go ExternalSendDelete(1,1)
			}
			go ExternalSendDelete(1,0)
		}
	} else if elevdriver.GetFloor() == 2 && elevinf.last_direction == DOWN{
		for i := 1; i < 3; i++ {
			elevinf.internal_orders[1][i] = 0
			elevinf.external_orders[1][1] = 0
			if elevinf.NoOrdersBelow() {
				fmt.Printf("Deleting order KNUUUT\n")
				elevinf.internal_orders[1][0] = 0
				elevinf.external_orders[1][0] = 0
				go ExternalSendDelete(1,0)
			}
			go ExternalSendDelete(1,1)
		}
	} else if elevdriver.GetFloor() == 3 && elevinf.last_direction == UP{
		for i := 0; i < 3; i = i+2 {
			elevinf.internal_orders[2][i] = 0
			elevinf.external_orders[2][0] = 0
			if elevinf.NoOrdersAbove() {
				elevinf.external_orders[2][1] = 0
				elevinf.internal_orders[2][1] = 0
				go ExternalSendDelete(2,1)
			}
			go ExternalSendDelete(2,0)
		}
	} else if elevdriver.GetFloor() == 3 && elevinf.last_direction == DOWN{
		for i := 1; i < 3; i++ {
			elevinf.internal_orders[2][i] = 0
			elevinf.external_orders[2][1] = 0
			if elevinf.NoOrdersBelow() {
				elevinf.external_orders[2][0] = 0
				elevinf.internal_orders[2][0] = 0
				go ExternalSendDelete(2,0)
			}
			go ExternalSendDelete(2,1)
		}
	} else if elevdriver.GetFloor() == 4{
		for i := 0; i < 3; i++ {
			elevinf.internal_orders[3][i] = 0
			elevinf.external_orders[3][1] = 0
			go ExternalSendDelete(3,1)
		}
	}
	
}

func (elevinf *Elevatorinfo) NoOrdersAbove ()bool{

	current := elevinf.last_floor
	orders_above_current := 0
		for i := current; i < 4; i++ {
			for j := 0; j < 3; j++ {
				if elevinf.internal_orders[i][j] == 1 {
					orders_above_current++
				}
			}
		}
	if orders_above_current == 0 {
		return true
	}
	
	return false
	
}

func (elevinf *Elevatorinfo) NoOrdersBelow ()bool{

	current := elevinf.last_floor
	orders_below_current := 0
		for i := 0; i < current-1; i++ {
			for j := 0; j < 3; j++ {
				if elevinf.internal_orders[i][j] == 1 {
					orders_below_current++
				}
			}
		}
	if orders_below_current == 0 {
		return true
	}
	
	return false
	
}

