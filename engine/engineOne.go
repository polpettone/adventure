package engine

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
	"github.com/faiface/beep/wav"
)

type EngineOne struct {
	sounds map[string]beep.Buffer
}

func (e EngineOne) ClearScreen() {
	fmt.Print("\033[H\033[2J")
}

func (e *EngineOne) Setup() {
	setupShellSettings()

	err := initSpeakerWithDummySound()
	if err != nil {
		log.Fatal(err)
	}

	sounds, err := initSound()
	if err != nil {
		log.Fatal(err)
	}
	e.sounds = sounds
}

func setupShellSettings() {
	// disable input buffering
	exec.Command("stty", "-F", "/dev/tty", "cbreak", "min", "1").Run()
	// do not display entered characters on the screen
	exec.Command("stty", "-F", "/dev/tty", "-echo").Run()
}

func initSpeakerWithDummySound() error {
	f, err := os.Open("assets/gunshot.mp3")
	if err != nil {
		return err
	}

	_, format, err := mp3.Decode(f)
	if err != nil {
		return err
	}
	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))

	return nil
}

func initSound() (map[string]beep.Buffer, error) {

	sounds := map[string]beep.Buffer{}

	soundFiles := []string{
		"assets/gunshot.mp3",
		"assets/mixkit-martial-arts-punch-2052.wav",
	}

	for _, s := range soundFiles {

		f, err := os.Open(s)
		if err != nil {
			return nil, err
		}

		extension := filepath.Ext(s)

		switch extension {
		case ".mp3":

			streamer, format, err := mp3.Decode(f)
			if err != nil {
				return nil, err
			}
			buffer := beep.NewBuffer(format)
			buffer.Append(streamer)
			streamer.Close()
			sounds[s] = *buffer

		case ".wav":

			streamer, format, err := wav.Decode(f)
			if err != nil {
				return nil, err
			}
			buffer := beep.NewBuffer(format)
			buffer.Append(streamer)
			streamer.Close()
			sounds[s] = *buffer
		default:
			return nil, fmt.Errorf("%s with sound encoding %s not supported", s, extension)
		}
	}

	return sounds, nil
}

func (e *EngineOne) PlaySound(soundID string) error {
	sound, exists := e.sounds[soundID]
	if exists {
		shot := sound.Streamer(0, sound.Len())
		speaker.Play(shot)
		return nil
	} else {
		return fmt.Errorf("sound with id: %s not found", soundID)
	}
}
