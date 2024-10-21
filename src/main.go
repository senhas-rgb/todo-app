package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

var option int8
var IsSetup int8 = 0; // false
var Flock string

func msg(text string) {
	fmt.Println("######################")
	fmt.Println(text)
	fmt.Println("######################")
}

func ScreenClear(period int) {
	time.Sleep(time.Duration(period) * time.Second)
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
		fmt.Scan(&option)
		if option == 1 {
			for {
				if IsSetup == 0 {
					ScreenClear(0)
					msg("Please configure in the storage locations in settings first.")
					ScreenClear(1)
					break
				} else if IsSetup == 1 {
					ScreenClear(0)
					Interface()
					break
				} else {
					ScreenClear(0)
					msg("Error.??!!")
					ScreenClear(1)
					break
				}
			}
		} else if option == 2 {
			ScreenClear(0)
			msg("You exited the program")
			ScreenClear(1)
			os.Exit(3)
		} else if option == 3 {
			ScreenClear(0)
			msg("Please enter the location of the file with extension.")
			msg("NOTE: Should be a *.txt file.")
			fmt.Scan(&Flock)
			if len(Flock) < 6 {
				ScreenClear(0)
				msg("File name too small?! Try again.")
				ScreenClear(1)
			} else {
				t1 := Flock[len(Flock)-1]
				t2 := Flock[len(Flock)-3]
				x1 := Flock[len(Flock)-2]
				period := Flock[len(Flock) -4]
				if  t1 == 't' && t2 == 't' && x1 == 'x' && period == '.' && len(Flock) > 6 {
					ScreenClear(0)
					msg(Flock)
					msg("Saved as the default location for tasks.")
					IsSetup = 1
					ScreenClear(3)
				} else if len(Flock) > 6 && t1 != 't' && t2 != 't' && period == '.' && x1 != 'x' {
					ScreenClear(0)
					msg("File format not supported.")
					ScreenClear(3)
				}
			}
		} else {
			ScreenClear(0)
			msg("Error! Try again...")
			ScreenClear(1)
			if IsSetup == 0 {
				msg("Also please configure in the storage locations in settings first.")
				ScreenClear(1)
			}
		}
	}
}

func main() {
	menu()
}
