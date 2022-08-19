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
	audioCtrl *beep.Ctrl
}

func (e EngineOne) ClearScreen() {
	fmt.Print("\033[H\033[2J")
}

func (e *EngineOne) Setup() {
	setupShellSettings()
	e.audioCtrl = initSound()
}

func setupShellSettings() {
	// disable input buffering
	exec.Command("stty", "-F", "/dev/tty", "cbreak", "min", "1").Run()
	// do not display entered characters on the screen
	exec.Command("stty", "-F", "/dev/tty", "-echo").Run()
}

func initSound() *beep.Ctrl {
	f, err := os.Open("assets/plink.mp3")
	if err != nil {
		log.Fatal(err)
	}

	streamer, format, err := mp3.Decode(f)
	if err != nil {
		log.Fatal(err)
	}

	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))

	ctrl := &beep.Ctrl{Streamer: beep.Loop(-1, streamer), Paused: true}
	fast := beep.ResampleRatio(4, 5, ctrl)

	speaker.Play(fast)

	return ctrl
}

func (e *EngineOne) PlaySound() {
	speaker.Lock()
	e.audioCtrl.Paused = false
	speaker.Unlock()
}

func (e *EngineOne) StopSound() {
	speaker.Lock()
	e.audioCtrl.Paused = true
	speaker.Unlock()
}
