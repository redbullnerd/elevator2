
// Gruppe12, Knut Hvamb & Christopher Benjamin Westlye, NTNU spring 2013

package elevator

import "elevdriver"
import "fmt"
import "time"

func (elevinf *Elevatorinfo) Initiate (){
	fmt.Printf("Halla\n")
	elevdriver.Init()
	fmt.Printf("Balle\n")
	StartMotor(-1)
	elevinf.last_direction = 2
	for elevdriver.GetFloor() == -1 {}
	
	elevinf.StopMotor()
	
	fmt.Printf("Elevator initiation complete!\n")
}

func (elevinf *Elevatorinfo) DetermineDirection ()(int){ // The elevators "brain", choosing which way to go depending on direction and orders
	
	current_floor := elevdriver.GetFloor()
	last_dir := elevinf.last_direction - 1
	orders_over := 0
	orders_under := 0 
	orders_at_current := 0
	
	for i := 0; i < 4; i++ {
		for j := 0; j < 3; j++ {
			if elevinf.internal_orders[i][j] == 1 && i < current_floor-1 {
				orders_under++
			} else if elevinf.internal_orders[i][j] == 1 && i > current_floor-1 {
				orders_over++
			} else if elevinf.internal_orders[i][j] == 1 && i == current_floor-1 {
				orders_at_current++
				fmt.Printf("%dorders, %dfloor, %dbutton\n",orders_at_current,i,j)
			}	
		}
	}
	
	if orders_at_current > 0 && (elevinf.internal_orders[current_floor-1][2] == 1  || elevinf.external_orders[current_floor-1][last_dir] == 1){
		fmt.Printf("Stay at floor\n")
		return -2 //Stay at floor
	} else if (orders_under > 0 && elevinf.last_direction == 2) || (orders_under > 0 && orders_over == 0) {
		fmt.Printf("blabla1\n")
		return -1 //Keep going down
	} else if (orders_over > 0 && elevinf.last_direction == 1) || (orders_over > 0 && orders_under == 0) {
		fmt.Printf("blabla1\n")
		return 1 //Keep going up
	} else {
		fmt.Printf("two is returned, do nothing\n")
		return 2 //No orders, no direction
	}
	
	return 2
} 

func StartMotor(direction int)() {
	if direction == -1 {
		elevdriver.MotorDown()
		fmt.Printf("Elevator going down...\n")
	} else if direction == 1 {
		elevdriver.MotorUp()
		fmt.Printf("Elevator going up...\n")
	} else if direction == -2 || direction == 2 {
		elevdriver.MotorStop()
	}
	
}

func (elevinf *Elevatorinfo) StopMotor(){
	if elevinf.last_direction == 1 {
		elevdriver.MotorDown()
		time.Sleep(10*time.Millisecond)
		elevdriver.MotorStop()
		fmt.Printf("Stopping...\n")
	} else if elevinf.last_direction == 2 {
		elevdriver.MotorUp()
		time.Sleep(10*time.Millisecond)
		fmt.Printf("Stopping...\n")
		elevdriver.MotorStop()
	}
}

func (elevator *Elevatorinfo) StopButtonPushed() {
	elevdriver.SetStopButton()
	fmt.Printf("Stop button has been pushed...\n")
	for i := 0; i < 4; i++ {
		for j := 0; j < 3; j++ {
			elevator.internal_orders[i][j] = 0
		}
	}
	elevdriver.MotorStop()
}







