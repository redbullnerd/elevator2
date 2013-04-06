
package elevator

import "elevdriver"
import "fmt"
import "time"

// checks pressed buttons and set lights accordingly
func CheckLights(state State, event Event, order_slice [][]int)(){
	
	if state != EMERGENCY || (state == EMERGENCY && event == ORDER) {
		for i := 0; i < 3; i++ {
			if order_slice[i][0] == 1 {
				elevdriver.SetLight(i, UP)
			}
			else if order_slice[i][0] == 0 {
				elevdriver.ClearLight(i, UP)
			}
		}
		for i := 1; i < 4; i++ {
			if order_slice[i][0] == 1 {
				elevdriver.SetLight(i, DOWN)
			}
			else if order_slice[i][0] == 0 {
				elevdriver.ClearLight(i, DOWN)
			}
		}
	}
	
	for i := 0; i < 4; i++ {
		if order_slice[k][2] == 1 {
			elevdriver.SetLight(k, NONE)
		}
		else if order_slice[k][2] == 0 {
			elevdriver.ClearLight(k, NONE)
		}
	}
}

// sets floorindicator light
func FloorIndicator(){

	if GetFloor()  > 0 { 
		SetFloor(GetFloor())
	}
	
}


