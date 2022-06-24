package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	var engine TextEngine = SimpleEngine{}
	m := InitMap()
	m.Place(NewBot(GOPHER), 0, 0)

	reader := bufio.NewReader(os.Stdin)

	fmt.Println(m.Print())

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

func InitMap() Map {
	fields := make([][]MapElement, 10)

	for n := range fields {
		fields[n] = make([]MapElement, 10)
	}

	for x := 0; x < 10; x++ {
		for y := 0; y < 10; y++ {
			fields[x][y] = NewField(FIELD)
		}
	}
	return Map{Fields: fields}
}
