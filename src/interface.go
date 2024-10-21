package main

import "fmt"

func Interface() {
	NewOption := 0
	ScreenClear(0)
	for {
		msg("Select a option:")
		msg("1) Add task\n2) View due tasks\n3) Archived tasks")
		fmt.Scan(&NewOption)
		switch NewOption {
		case 1:
			status_check()
			ScreenClear(0)
			msg("Under construction")
			ScreenClear(1)
			break
		case 2:
			ScreenClear(0)
			fmt.Println("Working")
			msg("Under construction")
			ScreenClear(1)
			break
		case 3:
			ScreenClear(0)
			fmt.Println("Working")
			msg("Under construction")
			ScreenClear(1)
			break
		default:
			ScreenClear(0)
			fmt.Println("Working")
			msg("Error try again...")
			ScreenClear(1)
		}
	}
}
