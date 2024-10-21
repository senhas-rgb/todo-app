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
		fmt.Println("1) Continue\n2) Leave\n------")
		fmt.Scan(&option)
		if option == 1 {
			ScreenClear(0)
			Interface()
			ScreenClear(1)
		} else if option == 2 {
			ScreenClear(0)
			msg("You exited the program")
			ScreenClear(1)
			os.Exit(3)
		} else {
			ScreenClear(0)
			msg("Error! Try again...")
			ScreenClear(1)
		}
	}
}

func main() {
	menu()
}
