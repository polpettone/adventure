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
	player2 := models.NewPlayer(models.GOPHER, 79, 29, "w", "s", "d", "a")

	gameMap := models.NewMap(80, 30)
	gameMap.Place(models.NewField(models.BOX, 30, 5))
	gameMap.Place(models.NewField(models.BOX, 10, 10))
	gameMap.Place(models.NewField(models.BOX, 40, 15))
	gameMap.Place(models.NewField(models.BOX, 55, 20))

	gameMap.Place(player)
	gameMap.Place(player2)

	clearScreen()
	fmt.Println(gameMap.Print())

	var engine Engine = SimpleEngine{
		GameMap: gameMap,
		Player:  player,
		Player2: player2,
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
	Player2 *models.Player
}

func clearScreen() {
	fmt.Print("\033[H\033[2J")
}

func updatePlayer(key string, gameMap *models.Map, player *models.Player) string {
	switch key {

	case player.MoveUp:
		if player.Y == gameMap.MaxY-1 {
			return "wall, cant move"
		}
		clearScreen()
		gameMap.Place(models.NewField(models.FIELD, player.X, player.Y))
		player.Y += 1
		gameMap.Place(player)
		fmt.Println(gameMap.Print())
		return "moved up"

	case player.MoveRight:
		if player.X == gameMap.MaxX-1 {
			return "wall, cant move"
		}
		clearScreen()
		gameMap.Place(models.NewField(models.FIELD, player.X, player.Y))
		player.X += 1
		gameMap.Place(player)
		fmt.Println(gameMap.Print())
		return "moved right"

	case player.MoveLeft:
		if player.X == 0 {
			return "wall, cant move"
		}
		clearScreen()
		gameMap.Place(models.NewField(models.FIELD, player.X, player.Y))
		player.X -= 1
		gameMap.Place(player)
		fmt.Println(gameMap.Print())
		return "moved left"

	case player.MoveDown:
		if player.Y == 0 {
			return "wall, cant move"
		}
		clearScreen()
		gameMap.Place(models.NewField(models.FIELD, player.X, player.Y))
		player.Y -= 1
		gameMap.Place(player)
		fmt.Println(gameMap.Print())
		return "moved down"
	}
	return "nothing happend"
}

func (se SimpleEngine) Machine(key string) string {
	updatePlayer(key, &se.GameMap, se.Player)
	updatePlayer(key, &se.GameMap, se.Player2)
	return "nothing happend"
}
