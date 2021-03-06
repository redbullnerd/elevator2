
// Gruppe12, Knut Hvamb & Christopher Benjamin Westlye, NTNU spring 2013

package elevator

func (elevinf *Elevatorinfo) MyCost (ordered_floor int) (cost int){
	// Because of package and unpackage, ordered_fllor is an int representing both wanted floor and direction
	my_cost := 0
	for i := 0; i < 4; i++ {
		for j := 0; j < 3; j++ {
			if elevinf.internal_orders[i][j] == 1 {
					my_cost++
				}
		}
	}
	
	wanted_floor := 0
	if ordered_floor == 0 {
		wanted_floor = 1
	} else if ordered_floor == 1 || ordered_floor == 3 {
		wanted_floor = 2
	} else if ordered_floor == 2 || ordered_floor == 4 {
		wanted_floor = 3
	} else if ordered_floor == 5 {
		wanted_floor = 4
	}
	
	var wanted_direction Direction = 0
	if ordered_floor == 0 || ordered_floor == 1 || ordered_floor == 3 {
		wanted_direction = 1
	} else if ordered_floor == 2 || ordered_floor == 4 || ordered_floor == 5 {
		wanted_direction = -1	
	}
	
	if elevinf.last_floor > wanted_floor && elevinf.last_direction != wanted_direction {
		my_cost++
	} else if elevinf.last_floor < wanted_floor && elevinf.last_direction != wanted_direction {
		my_cost++
	} else if elevinf.last_floor == wanted_floor {
		my_cost++
	}
	
	if elevinf.state == EMERGENCY {
		my_cost = my_cost+100
	}
	
	return cost
	
}
