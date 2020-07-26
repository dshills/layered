package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func main() {
	home := os.Getenv("HOME")
	if home == "" {
		log.Fatal("Home not found")
	}
	logfile := filepath.Join(home, ".lyd", "lyd.log")
	f, err := os.Create(logfile)
	if err != nil {
		fmt.Printf("[ERROR] %v", err)
		os.Exit(1)
	}
	log.SetOutput(f)
	app := NewApp()
	app.ProcessKeys()
}
