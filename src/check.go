package main

import (
	"os"
)

func status_check() {
	FilePath := "storage.txt"
	file, err := os.Open(FilePath)
	if err != nil {
		ScreenClear(0)
		msg("Error opening the storage.txt. Maybe deleted???")
		ScreenClear(1)
		return
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		ScreenClear(0)
		msg("Error getting file info???")
		ScreenClear(1)
		return
	}
	if fileInfo.Size() == 0 {
		ScreenClear(0)
		msg("No tasks found.")
		ScreenClear(1)
	} else {
		ScreenClear(0)
		msg("Some tasks are present.")
		ScreenClear(1)
	}
}
