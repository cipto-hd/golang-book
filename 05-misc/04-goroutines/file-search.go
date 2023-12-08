package main

import (
	"log"
	"os"
	"path/filepath"
)

func SearchFiles(dir string, lookFor string) string {
	log.Println("[SEARCHING] ", dir)
	files, err := os.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		log.Println(dir+file.Name(), file.IsDir())
		if file.Name() == lookFor {
			return "[FOUND] " + filepath.Join(dir, file.Name())
		}
	}
	return "[NOT FOUND] " + dir
}

func main() {
	result := make([]string, 0)
	go func() {
		result = append(result, SearchFiles("./test", "test2.txt"))
	}()
}
