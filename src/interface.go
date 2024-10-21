package main

import "fmt"

func Interface() {
	ScreenClear(0)
	for {
		msg("Select a option:")
		msg("1) Add task\n2) View due tasks\n3) Archived tasks")
		fmt.Scan(&option)
		switch option {
		case 1:
			ScreenClear(0)
			status_check()
			text := parse("upcoming.txt")
			fmt.Println(text)
			break
		case 2:
			ScreenClear(0)
			msg("Under construction")
			ScreenClear(1)
			break
		case 3:
			ScreenClear(0)
			msg("Still working on it")
			ScreenClear(1)
			break
		default:
			ScreenClear(0)
			msg("Error try again...")
			ScreenClear(1)
		}
	}
}
