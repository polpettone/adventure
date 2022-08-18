package engine

import (
	"fmt"
	"os/exec"
)

type EngineOne struct {
}

func (e EngineOne) ClearScreen() {
	fmt.Print("\033[H\033[2J")
}

func (e EngineOne) Setup() {
	setupShellSettings()
}

func setupShellSettings() {
	// disable input buffering
	exec.Command("stty", "-F", "/dev/tty", "cbreak", "min", "1").Run()
	// do not display entered characters on the screen
	exec.Command("stty", "-F", "/dev/tty", "-echo").Run()
}
