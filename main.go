package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

var gameMap Map
var bot Bot

func main() {
	start()
}

func start() {

	// disable input buffering
	exec.Command("stty", "-F", "/dev/tty", "cbreak", "min", "1").Run()
	// do not display entered characters on the screen
	exec.Command("stty", "-F", "/dev/tty", "-echo").Run()

	bot := NewBot(GOPHER, 0, 0)
	gameMap := NewMap(80, 30)
	gameMap.Place(NewField(BOX, 30, 10))
	gameMap.Place(bot)
	fmt.Println(gameMap.Print())

	var engine Engine = SimpleEngine{
		GameMap: gameMap,
		Bot:     bot,
	}

	var b []byte = make([]byte, 1)
	for {
		os.Stdin.Read(b)
		text := string(b)
		engine.Machine(text)

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

func clearScreen() {
	fmt.Print("\033[H\033[2J")
}

func (se SimpleEngine) Machine(text string) string {

	switch text {

	case "k":
		if se.Bot.Y == se.GameMap.MaxY-1 {
			return "wall, cant move"
		}
		clearScreen()
		se.GameMap.Place(NewField(FIELD, se.Bot.X, se.Bot.Y))
		se.Bot.Y += 1
		se.GameMap.Place(se.Bot)
		fmt.Println(se.GameMap.Print())
		return "moved up"

	case "l":
		if se.Bot.X == se.GameMap.MaxX-1 {
			return "wall, cant move"
		}
		clearScreen()
		se.GameMap.Place(NewField(FIELD, se.Bot.X, se.Bot.Y))
		se.Bot.X += 1
		se.GameMap.Place(se.Bot)
		fmt.Println(se.GameMap.Print())
		return "moved right"

	case "h":
		if se.Bot.X == 0 {
			return "wall, cant move"
		}
		clearScreen()
		se.GameMap.Place(NewField(FIELD, se.Bot.X, se.Bot.Y))
		se.Bot.X -= 1
		se.GameMap.Place(se.Bot)
		fmt.Println(se.GameMap.Print())
		return "moved left"

	case "j":
		if se.Bot.Y == 0 {
			return "wall, cant move"
		}
		clearScreen()
		se.GameMap.Place(NewField(FIELD, se.Bot.X, se.Bot.Y))
		se.Bot.Y -= 1
		se.GameMap.Place(se.Bot)
		fmt.Println(se.GameMap.Print())
		return "moved down"
	}

	return "nothing happend"
}
