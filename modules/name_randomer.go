package modules

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func clearTerminal() {
	fmt.Print("\033[H\033[2J") // ANSIエスケープシーケンスを使用してターミナルをクリア
}

func randomString(length int) string {
	rand.Seed(time.Now().UnixNano())
	const charset = "abcdefghijklmnopqrstuvwxyz0123456789"
	result := make([]byte, length)
	for i := range result {
		result[i] = charset[rand.Intn(len(charset))]
	}
	return string(result)
}

func NameRandomer() {
	clearTerminal()
	fmt.Println("[Name Randomer] This script will change the filenames to random strings.")
	var directory string
	fmt.Print("Enter the directory path: ")
	reader := bufio.NewReader(os.Stdin)
	directory, _ = reader.ReadString('\n')
	directory = strings.TrimSpace(directory) // Remove trailing newline

	files, err := ioutil.ReadDir(directory)
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return
	}

	for _, file := range files {
		if !file.IsDir() {
			newFilename := randomString(20) + filepath.Ext(file.Name())
			err := os.Rename(filepath.Join(directory, file.Name()), filepath.Join(directory, newFilename))
			if err != nil {
				fmt.Printf("Error renaming %s: %s\n", file.Name(), err)
			} else {
				fmt.Printf("renamed: %s -> %s\n", file.Name(), newFilename)
			}
		}
	}
}
