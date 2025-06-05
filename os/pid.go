package main

import (
	"fmt"
	"os"
)

func main() {
	pid := os.Getpid()
	fmt.Printf("Current process ID: %d\n", pid)
}
