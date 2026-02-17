package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Do gat help")
		return
	}

	switch os.Args[1] {
	case "help", "--help", "-h":
		{
			fmt.Println("gat is a replacement of the command cat but with syntax highlighting. Use like: gat ~/.bashrc")
			return
		}
	default:
		{
			fmt.Println("Unknown command!")
		}
	}

	file := os.Args[1]

	openthefile(file)
}
