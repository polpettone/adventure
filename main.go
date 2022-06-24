package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	var engine TextEngine = SimpleEngine{}

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("")
		text, _ := reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)

		response := engine.Machine(text)

		fmt.Printf("%s\n", response)

		if strings.Compare(text, "q") == 0 {
			fmt.Println("ciao")
			os.Exit(0)
		}
	}
}

type TextEngine interface {
	Machine(text string) string
}

type SimpleEngine struct{}

func (se SimpleEngine) Machine(text string) string {
	return "ich kann gar nix"
}
