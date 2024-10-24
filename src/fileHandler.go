package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func check(e error) { // helper to streamline error checks below (even I dont know wuts happening here)
	if e != nil {
		panic(e)
	}
}

func parse(fileLoc string) [][]string {
	//filelocation := "storage.txt"

	file, err := os.Open(fileLoc)
	check(err)
	defer file.Close()
	scanner := bufio.NewScanner(file) // new scanner to get each line of file

	var lines [][]string
	for scanner.Scan() {
		line := scanner.Text()        // Reads line
		words := strings.Fields(line) // split line to words
		lines = append(lines, words)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading the file. Scanner error. : ", err)
		return nil
	}

	//	for i, words := range lines {
	//		fmt.Printf("Line %d: %v\n", i+1, words)
	//		msg(lines[0][0])
	//	}
	//println(lines)
	return lines
}

func Fwrite(wString string, fileLoc string) int {
	file, err := os.OpenFile(fileLoc, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	defer file.Close()
	_, err = file.WriteString(wString)
	check(err)
	if err == nil {
		return 0
	} else {
		return 1
	}
}
