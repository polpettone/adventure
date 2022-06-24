package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var gameMap Map
var bot Bot

func main() {
	start()
}

func start() {

	bot := NewBot(GOPHER, 0, 0)
	gameMap := NewMap(10, 10)
	gameMap.Place(bot)
	fmt.Println(gameMap.Print())

	var engine Engine = SimpleEngine{
		GameMap: gameMap,
		Bot:     bot,
	}
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

type Engine interface {
	Machine(text string) string
}

type SimpleEngine struct {
	GameMap Map
	Bot     *Bot
}

func (se SimpleEngine) Machine(text string) string {

	switch text {

	case "k":
		if se.Bot.Y == se.GameMap.MaxY-1 {
			return "wall, cant move"
		}
		se.GameMap.Place(NewField(FIELD, se.Bot.X, se.Bot.Y))
		se.Bot.Y += 1
		se.GameMap.Place(se.Bot)
		fmt.Println(se.GameMap.Print())
		return "moved up"

	case "l":
		if se.Bot.X == se.GameMap.MaxX-1 {
			return "wall, cant move"
		}
		se.GameMap.Place(NewField(FIELD, se.Bot.X, se.Bot.Y))
		se.Bot.X += 1
		se.GameMap.Place(se.Bot)
		fmt.Println(se.GameMap.Print())
		return "moved right"

	case "h":
		if se.Bot.X == 0 {
			return "wall, cant move"
		}
		se.GameMap.Place(NewField(FIELD, se.Bot.X, se.Bot.Y))
		se.Bot.X -= 1
		se.GameMap.Place(se.Bot)
		fmt.Println(se.GameMap.Print())
		return "moved up"

	case "j":
		if se.Bot.Y == 0 {
			return "wall, cant move"
		}
		se.GameMap.Place(NewField(FIELD, se.Bot.X, se.Bot.Y))
		se.Bot.Y -= 1
		se.GameMap.Place(se.Bot)
		fmt.Println(se.GameMap.Print())
		return "moved up"
	}

	return "nothing happend"
}
