package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/polpettone/adventure/engine"
)

func main() {
	startPinguinBurgGame()
}

func startPinguinBurgGame() {

	game := &engine.PinguinBurfGame{}
	game.Init()

	var b []byte = make([]byte, 1)
	for {
		os.Stdin.Read(b)
		text := string(b)
		game.Update(text)

		if strings.Compare(text, "q") == 0 {
			fmt.Println("ciao")
			os.Exit(0)
		}
	}
}
