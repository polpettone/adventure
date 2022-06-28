package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/polpettone/adventure/models"
)

func main() {
	setupShellSettings()
	start()
}

func setupShellSettings() {
	// disable input buffering
	exec.Command("stty", "-F", "/dev/tty", "cbreak", "min", "1").Run()
	// do not display entered characters on the screen
	exec.Command("stty", "-F", "/dev/tty", "-echo").Run()
}

func initGame() Engine {
	player := models.NewPlayer(models.GOPHER, 0, 0, "k", "j", "l", "h")
	gameMap := models.NewMap(80, 30)
	gameMap.Place(models.NewField(models.BOX, 30, 5))
	gameMap.Place(models.NewField(models.BOX, 10, 10))
	gameMap.Place(models.NewField(models.BOX, 40, 15))
	gameMap.Place(models.NewField(models.BOX, 55, 20))
	gameMap.Place(player)

	clearScreen()
	fmt.Println(gameMap.Print())

	var engine Engine = SimpleEngine{
		GameMap: gameMap,
		Player:  player,
	}
	return engine
}

func start() {

	engine := initGame()

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
	GameMap models.Map
	Player  *models.Player
}

func clearScreen() {
	fmt.Print("\033[H\033[2J")
}

func updatePlayer(key string, se SimpleEngine) string {
	switch key {

	case se.Player.MoveUp:
		if se.Player.Y == se.GameMap.MaxY-1 {
			return "wall, cant move"
		}
		clearScreen()
		se.GameMap.Place(models.NewField(models.FIELD, se.Player.X, se.Player.Y))
		se.Player.Y += 1
		se.GameMap.Place(se.Player)
		fmt.Println(se.GameMap.Print())
		return "moved up"

	case se.Player.MoveRight:
		if se.Player.X == se.GameMap.MaxX-1 {
			return "wall, cant move"
		}
		clearScreen()
		se.GameMap.Place(models.NewField(models.FIELD, se.Player.X, se.Player.Y))
		se.Player.X += 1
		se.GameMap.Place(se.Player)
		fmt.Println(se.GameMap.Print())
		return "moved right"

	case se.Player.MoveLeft:
		if se.Player.X == 0 {
			return "wall, cant move"
		}
		clearScreen()
		se.GameMap.Place(models.NewField(models.FIELD, se.Player.X, se.Player.Y))
		se.Player.X -= 1
		se.GameMap.Place(se.Player)
		fmt.Println(se.GameMap.Print())
		return "moved left"

	case se.Player.MoveDown:
		if se.Player.Y == 0 {
			return "wall, cant move"
		}
		clearScreen()
		se.GameMap.Place(models.NewField(models.FIELD, se.Player.X, se.Player.Y))
		se.Player.Y -= 1
		se.GameMap.Place(se.Player)
		fmt.Println(se.GameMap.Print())
		return "moved down"
	}
	return "nothing happend"
}

func (se SimpleEngine) Machine(key string) string {
	updatePlayer(key, se)
	return "nothing happend"
}
