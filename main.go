package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
        fmt.Println("No function specified")
        return
    }

	ConnectDB()

    switch os.Args[1] {
    case "router":
        router()
    case "migrate":
        migrate()
    default:
        fmt.Println("Unknown function: " + os.Args[1])
	}
}