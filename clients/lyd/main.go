package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/dshills/layered/logger"
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

	app := App{done: make(chan struct{})}
	if err := app.init(); err != nil {
		fmt.Printf("[ERROR] app.init %v", err)
		os.Exit(1)
	}
	logger.Debugf("%+v", app)
	app.ProcessKeys()
}
