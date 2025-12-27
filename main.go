package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"Minireddis/server"
	"Minireddis/storage"
)

const file = "data.bin"

func main() {
	store := storage.NewStore()

	if err := storage.LoadFromFile(store, file); err == nil {
		fmt.Println("📦 Data loaded")
	} else {
		fmt.Println("📭 Fresh start")
	}

	go server.StartTCP(":6379", store)
	go server.StartHTTP(":8080", store)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	<-c
	fmt.Println("\n💾 Saving data...")
	storage.SaveToFile(store, file)
	fmt.Println("👋 Bye")
}
