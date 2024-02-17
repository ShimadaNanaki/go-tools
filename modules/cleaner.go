package modules

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func ClearTerminal() {
	fmt.Print("\033[H\033[2J")
}

func FileCleaner() {
	ClearTerminal()
	fmt.Println("[File Cleaner] This script will remove files that are less than or equal to 1KB.")
	var folder string
	fmt.Print("Enter the directory path: ")
	reader := bufio.NewReader(os.Stdin)
	folder, _ = reader.ReadString('\n')
	folder = strings.TrimSpace(folder) // Remove trailing newline

	files, err := ioutil.ReadDir(folder)
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return
	}

	removedFilesCount := 0
	for _, file := range files {
		if !file.IsDir() && file.Size() <= 1024 { // 1KB = 1024 bytes
			err := os.Remove(folder + "/" + file.Name())
			if err == nil {
				fmt.Printf("removed: %s \n", file.Name())
				removedFilesCount++
			}
		}
	}

	if removedFilesCount == 0 {
		fmt.Println("No applicable files were found.")
	} else {
		fmt.Printf("Total files removed: %d\n", removedFilesCount)
	}
}
