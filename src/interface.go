package main

import "fmt"

//import "fmt"

func Interface() {
	ScreenClear(0)
	for {
		msg("Select a option:")
		msg("1) Add task\n2) View due tasks\n3) Archived tasks")
		fmt.Scan(&option)
		if option == 1 {
			status_check()
			msg("Under construction")
			ScreenClear(3)
			break
		} else if option == 2 {
			ScreenClear(0)
			msg("Under construction")
			ScreenClear(3)
			break
		} else {
			ScreenClear(0)
			msg("Error try again...")
			ScreenClear(3)
		}
	}
}
