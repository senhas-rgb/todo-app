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
			msg("Under construction")
			break
		} else if option == 2 {
			msg("Under construction")
			break
		} else {
			msg("Error try again...")
		}
	}
}
