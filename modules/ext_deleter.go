package modules

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func ExtDeleter() {
	ClearTerminal()
	fmt.Println("[Ext Deleter] This script will delete files with the specified extension.")
	var ext string
	fmt.Print("Enter the extension of the files to delete: ")
	fmt.Scanln(&ext)

	var pathToDir string
	fmt.Print("Enter the directory path: ")
	reader := bufio.NewReader(os.Stdin)
	pathToDir, _ = reader.ReadString('\n')
	pathToDir = strings.TrimSpace(pathToDir)

	ext = strings.TrimLeft(ext, ".")
	files, err := filepath.Glob(pathToDir + "/*." + ext)
	if err != nil {
		fmt.Println("Error finding files:", err)
		return
	}

	removedFilesCount := 0
	for _, file := range files {
		err := os.Remove(file)
		if err == nil {
			fmt.Printf("removed: %s \n", file)
			removedFilesCount++
		} else {
			fmt.Printf("Problem occurred: %v\n", err)
		}
	}

	if removedFilesCount == 0 {
		fmt.Println("No applicable files were found.")
	} else {
		fmt.Printf("Total files removed: %d\n", removedFilesCount)
	}
}
