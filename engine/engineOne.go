package engine

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
)

type EngineOne struct {
	soundBuffer beep.Buffer
}

func (e EngineOne) ClearScreen() {
	fmt.Print("\033[H\033[2J")
}

func (e *EngineOne) Setup() {
	setupShellSettings()
	e.soundBuffer = initSound()
}

func setupShellSettings() {
	// disable input buffering
	exec.Command("stty", "-F", "/dev/tty", "cbreak", "min", "1").Run()
	// do not display entered characters on the screen
	exec.Command("stty", "-F", "/dev/tty", "-echo").Run()
}

func initSound() beep.Buffer {
	f, err := os.Open("assets/gunshot.mp3")
	if err != nil {
		log.Fatal(err)
	}

	streamer, format, err := mp3.Decode(f)
	if err != nil {
		log.Fatal(err)
	}

	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))

	buffer := beep.NewBuffer(format)

	buffer.Append(streamer)
	streamer.Close()
	return *buffer
}

func (e *EngineOne) PlaySound() {
	shot := e.soundBuffer.Streamer(0, e.soundBuffer.Len())
	speaker.Play(shot)
}
