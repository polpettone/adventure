package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/polpettone/adventure/engine"
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

func start() {

	engine := engine.InitEngineOne()

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
