package main

import (
	"os"
)

func status_check() {
	FilePath := "storage.txt"
	file, err := os.Open(FilePath)
	if err != nil {
		ScreenClear(0)
		msg("Error opening storage.txt. Maybe the file was deleted?")
		ScreenClear(600)
		return
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		ScreenClear(0)
		msg("Error getting file info")
		ScreenClear(600)
		return
	}
	if fileInfo.Size() == 0 {
		ScreenClear(0)
		msg("No tasks found.")
		ScreenClear(0)
	} else {
		ScreenClear(0)
		msg("Some tasks are present.")
		ScreenClear(600)
	}
}
