package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

var option int8
var IsSetup int8 = 0 // false
var fileLoc string

func msg(text string) {
	fmt.Println("######################")
	fmt.Println(text)
	fmt.Println("######################")
}

func ScreenClear(period int) {
	time.Sleep(time.Duration(period) * time.Millisecond)
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func menu() {
	for {
		ScreenClear(0)
		fmt.Println("######TODO APP########")
		fmt.Println("########IN GO#########")
		fmt.Println("Select option:")
		fmt.Println("1) Continue\n2) Leave\n3) Settings\n------")
		fmt.Print(">>")
		fmt.Scan(&option)
		if option == 1 {
			for {
				if IsSetup == 0 {
					ScreenClear(0)
					msg("Please configure in the storage locations in settings first.")
					ScreenClear(600)
					break
				} else if IsSetup == 1 {
					ScreenClear(0)
					Interface(fileLoc)
					break
				} else {
					ScreenClear(0)
					msg("Error.??!!")
					ScreenClear(600)
					break
				}
			}
		} else if option == 2 {
			ScreenClear(0)
			msg("You exited the program")
			ScreenClear(600)
			os.Exit(3)

		} else if option == 3 {
			ScreenClear(0)
			fmt.Printf("Please enter a file for saving tasks.\n\n")
			msg("NOTE: The file should be a .txt file.")
			fmt.Print(">>")
			fmt.Scan(&fileLoc)

			if len(fileLoc) < 4 {
				ScreenClear(0)
				msg("File name too short")
				ScreenClear(700)
			} else if fileLoc[len(fileLoc)-4:] == ".txt" && len(fileLoc) > 5 {
				ScreenClear(0)
				outMsg := "\"" + fileLoc + "\"" + " saved as the default location for tasks."
				msg(outMsg)
				IsSetup = 1
				ScreenClear(600)

			} else {
				ScreenClear(0)
				msg("File format not supported.")
				ScreenClear(600)
			}
		} else {
			ScreenClear(0)
			msg("Error! Try again...")
			ScreenClear(600)
			if IsSetup == 0 {
				msg("Please configure in the storage locations in settings first.")
				ScreenClear(800)
			}
		}
	}
}

func main() {
	menu()
}
