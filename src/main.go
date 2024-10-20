package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

var option int8


func msg(text string) {
	fmt.Println("######################")
	fmt.Println(text)
	fmt.Println("######################\n")
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
		fmt.Println("######TODO APP########\n")
		fmt.Println("########IN GO#########\n")
		fmt.Println("Select option: \n")
		fmt.Println("1) Continue\n2) Leave\n------")
		fmt.Scan(&option)
		if option == 1 {
			app()
		} else if option == 2 {
			ScreenClear(0)
			msg("You excited out of the program")
			ScreenClear(3)
			os.Exit(3)
		} else {
			ScreenClear(0)
			msg("Error Try again...\n\n")
			ScreenClear(3)
		}
	}
}

func main() {
	menu()
}
