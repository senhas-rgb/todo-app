package main

import (
	"fmt"
	"strings"
)

func Interface() {
	ScreenClear(0)
	var option string
	for {
		option = ""
		fmt.Println("Select a option:")
		fmt.Println("\t1) Add task\n\t2) View due tasks\n\t3) Archived tasks\nPress enter to exit")
		fmt.Print(">>")
		fmt.Scanln(&option)
		switch option {
		case "1":
			ScreenClear(0)
			status_check()
			//text := parse("upcoming.txt")
			//fmt.Println(text)
			break
		case "2":
			ScreenClear(0)
			msg("Under construction")
			ScreenClear(300)
			fmt.Println("Reading \"Upcoming.txt\"")
			ScreenClear(100)
			text := parse("upcoming.txt")
			for i := 0; i < len(text); i++ {
				fmt.Printf("Task %d: %s\n", i+1, strings.Join(text[i], " "))
			}
			fmt.Printf("\nPress enter to exit")
			fmt.Scanln()
			ScreenClear(0)
			break
		case "3":
			ScreenClear(0)
			msg("Still working on it")
			ScreenClear(600)
			break
		case "":
			println("Exiting...")
			return
		default:
			ScreenClear(0)
			msg("Error try again...")
			ScreenClear(700)
		}
	}
}
